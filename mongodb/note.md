mongodb  
MongoDb的集合collection相当于Mysql的一个表  
文档 相当于mysql里面的一行数据  

// use 创建数据库
use test
// show dbs 查看数据库

// 统计数据库信息
db.stats() 统计数据库信息

// 查询当前数据下的所有数据集合  
db.getCollectionNames()

// 创建集合  
db.createCollection(name, options)
options
capped - 启用封闭的集合  
size - 指定上限集合的最大大小  
max - 指定上限集合中允许的最大文档数  

文档是mongo中存储的基本单元，是一组有序的键值对集合  
键名 - 区别大小写 不能重复
键值 - 区别类型 
{“uid”:5}

// 集合中插入文档  
db.collection.insert(  
	{name:"lijin"},   
{	  
	// 可选字段   
	writeConcern:<document> // 出错捕获机制  
	ordered:<bool> // 默认为true 有序插入  
	}  
)  
db.collection.insertOne() // 插入一条  
db.collection.insertMany() // 插入多条  

// 更新或修改集合中的文档  
db.collection.update(  
	{name:xx},  
	$set{name:xxx}  
)  
 
// 如果存在_id相同的，则直接替换  
// 如果不存在，则会新插入  
db.collection.save({_id:xx, name:"xx"})  

// 删除数据  
db.collection.remove(  
	{name:"xx"},  // 查询条件  
	{  
	justOne:<bool> // 默认false删除条件的所有文件 true-只删除一条  
	}  
)  

// 删除一行数据  
db.collection.deleteOne({_id:1})  
// 删除所有匹配数据  
db.collection.deleteMany({_name:"xx"})  

// 查询集合  
db.collection.find()  
db.collection.find().pretty() // 格式化显示数据  

// 查询并按_id升序返回结果 1升序 -1降序  
db.collection.find().sort({_id:1})  
