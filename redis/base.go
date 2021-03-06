package _redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

// 初始化连接
// Redis的连接方式可以分为 普通连接 连接哨兵模式（sentinel） 连接Redis集群（cluster）
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("connection failed, err%s\n", err)
		return err
	}

	return nil
}

// 入口调用
func Start() {
	err := initClient()
	if err != nil {
		fmt.Println("initClient failure ", err)
		return
	}
	//redisExample()
	//stringDemo()
	//hashesDemo()

	// 集合操作练习
	//setsDemon()
	// 有序集合练习
	sortedSetsDemo()
}

// set/get基本示例
func redisExample() {
	err := rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "score").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("retrieve score: ", val)
}

// Strings操作
func stringDemo() {
	rdb.Set(ctx, "name", "lee", 0)

	// 返回key - value
	value := rdb.Get(ctx, "name")
	fmt.Println(value.Val())

	// GetRange 返回 start ~ end 对应key value的子串
	// 下标 0第一位 -1最后一位
	value = rdb.GetRange(ctx, "name", 1, -1)
	fmt.Println(value.Val())

	// MGet 返回所有指定key的value 不存在仍会返回一个nil
	sli := rdb.MGet(ctx, "name", "newName", "nameTemp")
	for _, v := range sli.Val() {
		fmt.Println(v)
	}
	// MSet 同MGet操作原理
	// MSetNX 给定对应的keys设置相应的values 如果其中一个key已经存在，操作都不会执行
	values := map[string]interface{}{
		"name3": "lijin3",
		"name4": "lijin4",
	}
	// 返回的是一个bool类型的值
	sli1 := rdb.MSetNX(ctx, values)
	fmt.Println("Is insert seccess ", sli1.Val())

	// 追加一个值原key-Value的结尾
	rdb.Append(ctx, "name", "Jin")
	fmt.Println(rdb.Get(ctx, "name"))

	// 统计字符串被设置为1的bit数
	// 设置BitCount Start End可以让计数在特定的位进行
	// -1 最后一位 0 第一位
	bitCount := &redis.BitCount{Start: 0, End: -1}
	length := rdb.BitCount(ctx, "name", bitCount)
	fmt.Println(length)

	// Decr 对key对应的数字做减1操作 Incr同理
	rdb.Set(ctx, "num", 10, 0)
	value1 := rdb.Decr(ctx, "num")
	fmt.Println(value1.Val())

	// DecrBy 指定减少数量 IncrBy同理
	value1 = rdb.DecrBy(ctx, "num", 2)
	fmt.Println(value1.Val())

	// 给key设置新value 并返回旧的value
	value2 := rdb.GetSet(ctx, "num", 5)
	fmt.Println("old value: ", value2.Val(), " new value: ", rdb.Get(ctx, "num").Val())

	// 指定浮点数增加
	value3 := rdb.IncrByFloat(ctx, "num", 0.1)
	fmt.Println(value3.Val())

	// 返回key的string类型value的长度 如果value非string 则返回错误
	value4 := rdb.StrLen(ctx, "name")
	fmt.Println("name length: ", value4.Val())
}

// 哈希数据操作
func hashesDemo() {
	// 设置key指定的哈希集中的指定字段的值
	rdb.HSet(ctx, "myhash", []string{"name", "lijin"})
	rdb.HSet(ctx, "myhash", []string{"name1", "lijin1"})

	// key 对应指定的哈希集 从哈希值找到对应的字段
	value1 := rdb.HGet(ctx, "myhash", "name")
	fmt.Println(value1.Val())

	// 从key指定的哈希集中移除指定的域 0 - 失败 1 - 成功
	value2 := rdb.HDel(ctx, "myhash", "name1")
	fmt.Println(value2.Result())
}

// Sets集合数据操作 （无序集合）
func setsDemon() {
	// 插入成员
	addSetMember()
	// 移出成员
	moveSetMember()
	// 查看成员
	seeSetMember("set1")
}

// 添加成员
func addSetMember() {
	resInt := rdb.SAdd(ctx, "set1", "c++")
	//result := rdb.SAdd(ctx, "set1", []string {"go", "lua"})
	num, err := resInt.Result()
	if err != nil {
		fmt.Printf("add members failed, err%s\n", err)
	}
	fmt.Printf("add %d members\n", num)
}

// 移出成员
func moveSetMember() {
	resBool := rdb.SMove(ctx, "set1", "set2", "c++")
	isSuc, err := resBool.Result()
	if err != nil {
		fmt.Printf("move member failed, err %s\n", err)
		return
	}
	fmt.Printf("move member, %v\n", isSuc)
	seeSetMember("set1")
}

// 查看集合中的成员
func seeSetMember(key string) {
	result := rdb.SMembers(ctx, key)
	strings, err := result.Result()
	if err != nil {
		fmt.Printf("get members failed, err %s\n", err)
		return
	}
	fmt.Println("set members: ", strings)
}

// 有序集合操作
func sortedSetsDemo() {
	// 添加成员
	addSortedSetMember()
	// 删除成员
	RemSortedSetMember()
	//seeSortedSetMember()
	// 查看成员
	//seeSortedSetMember1()
}

// 添加成员
func addSortedSetMember() {
	sortSet1 := &redis.Z{Member: "one", Score: 1}
	sortSet2 := &redis.Z{Member: "two", Score: 2}
	resInt := rdb.ZAdd(ctx, "sort1", sortSet1, sortSet2)
	num, err := resInt.Result()
	if err != nil {
		fmt.Printf("Sorted Sets add members, err%s\n", err)
		return
	}
	fmt.Printf("Sorted Sets add %d member\n", num)
}

// 删除成员
func RemSortedSetMember() {
	resInt := rdb.ZRem(ctx, "sort1", "two")
	num, err := resInt.Result()
	if err != nil {
		fmt.Printf("Sorted Sets remove member, err %s\n", err)
		return
	}
	fmt.Printf("Sorted Sets remove %d member\n", num)
	seeSortedSetMember1()
}

// 查看有序集合成员
func seeSortedSetMember() {
	resString := rdb.ZRange(ctx, "sort1", 0, -1)
	str, err := resString.Result()
	if err != nil {
		fmt.Printf("Sorted Sets see, err%s\n", err)
		return
	}
	fmt.Printf("Sorted Sets members %s\n", str)
}

// 查看有序结合成员带Score
func seeSortedSetMember1() {
	resSlice := rdb.ZRangeWithScores(ctx, "sort1", 0, -1)
	sli, err := resSlice.Result()
	if err != nil {
		fmt.Printf("Sorted Sets see, err%s\n", err)
		return
	}
	fmt.Printf("Sorted Sets members %v\n", sli)
}

// Sorted Sets 有序集合 - 成员都是唯一的不允许有重复
func sortedSetsDemon() {
	// 每个元素除了对应的值以外，还会多出一个分数 - 浮点数
	// 依赖分数进行排序
	// 集合是通过hash表实现的
}

// TODO 如何实现一个接口？
/*
type name interface {
	String ()
 }

type jam struct {
}

func (j *jam) String () {
	fmt.Println("this is my string")
}

func StartName() {
	var j = new(jam)
	j.String()
}
*/
