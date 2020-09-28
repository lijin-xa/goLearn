package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// 使用MongoDB Go Driver
// MongoDB所采用的数据格式被称为BSON - 基于JSON的二进制序列化格式
// Go 使用BSON对象
// Go Driver 有两个系列表示BSON D系列 Raw系列
// D类型有4种类型
// D - BSON文档
// M - Map
// A - 数组
// E - 在D里面的一个单一的子项
// 用于存储文档并进行远程过程调用
var Client *mongo.Client

// context - 传递和请求具有生命周期的变量的标准方法
// 程序单元的一个运行状态，现场，快照
var collection *mongo.Collection
var ctx context.Context

func Start() {
	fmt.Println("Start connect")

	// 使用Go Driver连接到MongoDB
	// 配置driver的设定 write concern 错误捕获机制
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB 必须传递两个参数 context 和 options.ClientOptions对象
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("connect failed, err%s\n", err)
		return
		//log.Fatal(err)
	}

	// 测试连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("client Ping failed, err%s\n", err)
		return
		//log.Fatal(err)
	}
	Client = client
	fmt.Println("Connected to MongoDB")

	OperateDemo()
}

// 基础操作练习
func OperateDemo() {
	// 返回一个操作集合的句柄
	collection = Client.Database("test").Collection("temp")
	fmt.Println("connect collection name:", collection.Name())

	//insertDemo()
	//findDemo()
	//deleteDemo()
	updateDemo()
}

type user struct {
	Id   int    `bson:"_id"`
	Name string `bson:"name"`
}

// 查询集合文档
func findDemo() {
	// 需要过滤文档
	filter := bson.M{"_id": 1}
	result := collection.FindOne(ctx, filter)

	msg := &user{}
	// 结果的解码
	err := result.Decode(msg)
	if err != nil {
		fmt.Printf("find one failed, err %s\n", err)
	}
	fmt.Printf("_id: %d, name: %s\n", msg.Id, msg.Name)

	// TODO 多条数据查找 暂时还存在问题
	filter = bson.M{"[key]": "[value]"}
	collection.Find(ctx, filter)
}

// 插入文档练习
func insertDemo() {
	// map类型
	b := bson.M{
		"_id":  1,
		"name": "lijin",
	}
	// 插入bson数据不能有_id字段
	_, err := collection.InsertOne(ctx, b)
	if err != nil {
		fmt.Printf("insert data failed, err%s\n", err)
		//return
	}
	// fmt.Printf("insert id %v\n", result.InsertedID)

	// map切片
	bsonList := []interface{}{
		bson.M{
			"_id":  2,
			"name": "go",
		},
		bson.M{
			"_id":  3,
			"name": "lua",
		},
	}
	// 插入多条数据
	res, err := collection.InsertMany(ctx, bsonList)
	if err != nil {
		fmt.Printf("InsertMany failed, err%s\n", err)
		return
	}
	fmt.Printf("InsertMany ids %v\n", res.InsertedIDs)
}

// 删除练习
func deleteDemo() {
	// 删除一条数据
	filter := bson.M{"_id": 3}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("delete failed, err%s\n", err)
		return
	}
	fmt.Println("result.DeletedCount", result.DeletedCount)
}

// 修改文档练习
func updateDemo() {
	filter := bson.M{"_id": "2"}
	update := bson.M{"name": "go"}

	// TODO update bson 暂时还有些问题
	// mongo.UpdateResult
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Printf("update failed, err %s\n", err)
	}
	fmt.Println("ModifiedCount: ", result.ModifiedCount)
}

// 以下老代码
type Trainer struct {
	Name string
	Age  int
	City string
}

// CURT操作（create update retrieve delete）增删改查操作
func OperateTest(client *mongo.Client) {
	// 获取test数据库的 trainers集合的handle句柄
	collection := client.Database("test").Collection("trainers")

	// 插入文档
	jam := Trainer{"Jam", 18, "xian"}
	insertResult, err := collection.InsertOne(context.TODO(), jam)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert a single document Id:", insertResult.InsertedID)

	// 更新文档
	filter := bson.D{{"name", "Jam"}} // 匹配过滤器
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateResult, err1 := collection.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Matched %v documents and update %v documents \n", updateResult.MatchedCount,
		updateResult.ModifiedCount)

	// 查找文档
	var checkResult Trainer

	// 错误记录argument to Decode must be a pointer or a map, but got { 0 }
	// Decode参数必须是一个指针或者map
	err = collection.FindOne(context.TODO(), filter).Decode(&checkResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", checkResult)

	// 删除文档
	deleteResult, err2 := collection.DeleteOne(context.TODO(), filter)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Printf("Delete %v document in the trainers collection\n", deleteResult.DeletedCount)
}
