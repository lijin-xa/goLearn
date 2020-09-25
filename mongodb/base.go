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
// 用于存储文档并进行远程过程调用

func Start() {
	fmt.Println("Start connect")

	// 使用Go Driver连接到MongoDB

	// 连接选项 URI不需要IP地址 直接使用端口号即可
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	OperateTest(client)
}

type Trainer struct {
	Name string
	Age int
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
	filter := bson.D{{"name", "Jam"}}  // 匹配过滤器
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
	err = collection.FindOne(context.TODO(), filter). Decode(&checkResult)
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
