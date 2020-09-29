Redis 笔记 (txt格式)  
Redis是一个开源的，内存中的数据结构存储系统，它可以用作数据库、缓存和中间件

redis学习 - 目前是官网学习  
1 - 支持的数据结构类型  
strings - 字符串  
SADD
SET key value 设置键值对  
GETSET key value 自动将key对应到value并且返回原来key对应的value  
INCR key 对存储在指定key的数值执行原子的加1操作  

DEL key 删除键  

EXPIRE    
expire 设置key的过期时间，超过时间后会自动删除该key  
expire key 10 10秒过期  
Redis 2.6开始 过期时间误差缩小到0-1毫秒 使用unix时间戳（毫秒）存储  
需要注意 如果两台时间不同步的主机同步RDB文件，过期时间就会变得不可控  

PERSIST  
persist 移除给定key的生存时间  

INFO 返回关于redis服务器的各种信息和统计数值  
通过给定的可选参数section  
比如:  
server Redis服务器相关的的信息  
clients 已连接的客户端的信息  
memory 则是内存使用情况  

KETS 查找所有符合给定模式pattern（正则表达式）的key  

TYPE 返回key所存储的value的数据类型 key不存在返回none  

Redis并不是简单的key-value存储，实际上他是一个数据结构服务器  
14个redis命令组 200多个redis命令  
二进制安全的字符串  
Lists - 列表  
按插入顺序排序的字符串元素的集合 链表  

Hashes - 散列  
由field和关联的value组成的map filed和value都是字符串  

Redis key值是二进制安全的，这意味着可以用任何二进制序列作为key值  
Strings - 字符串    

Sets - 集合  
不重复且无序的字符串元素的集合  
SAdd key member [member...] 添加成员到指定的集合key中  

SMembers key 返回key集合中的所有元素  

SCard key 返回key集合的元素个数  

SDiff key1 [key2...] 返回key1 和 key2... 差集的元素    
SDiffStore destination key1 [key2...] 跟上条命令基本一致 不过会将差集放在一个新的集合中  
如果 destination已经存在 会被覆盖重写  

SInter key1 [key2...] 返回key1 和 key2... 交集的元素  
SInterStore key1 [key2...] 跟上条命令基本一致 会将交集放在一个新的集合中    

SIsMember key member member是否是集合key的成员中 1 - 存在 0 - 不存在    

SMove source destination member 将source集合中的member移出到destination中    

SPop key [count] 从key的集合中移出并返回一个或多个元素 count可选项在3.2版本中可以使用  

SRandMember key [count] 随机返回key集合中的一个元素 如果存在count参数 则返回一个数组元素  

SRem key member [member...] 从key集合删除指定元素 返回删除元素的个数  

SUnion key1 [key2...] 返回给定集合的多个集合的并集中的所有成员  
SUnionStore destination key1 [key2...] 功能基本一致 不过将会返回的并集保存在destination中  

Sorted sets - 有序集合  
类似于Sets， 每一个字符串元素都会关联到 一个叫score浮动数值    
并且通过score排序，它可以检索到一系列元素    

ZADD key score member ...  向有序集合的key中添加成员  
ZCARD key 有序集合key中的成员数  
ZCOUNT key min max 根据分数返回对应的成员列表数量   
(min/max 可以不包含对应分数  -inf +inf 从最低分 到 最高分  
ZRANGE key start stop [withscores] 返回指定范围的元素  
加上withscores 会以此将值和score都显示  
不带withscores 只会显示成员值  

ZINCRBY key increment member 指定集合key的成员member的score值加上增量increment  
ZINCRBY sortset 2 "one"  

ZInterStore 计算一个或多个有序集的交集 - 交集中成员score是给定集的和  
ZInterStore destination numkeys key1 key2 ... [Weights weight]  
Weight - 乘数因子 给定集在合并前 先乘以这个因子 再合并交集值  

ZUnionStore  计算给定有序集的并集  
ZUnionStore "新的并集key" numkeys key1 key2 ...   

ZLEXCOUNT key min max 计算指定集合中指定成员间的成员个数（包含指定位置）  
min max 可以使用 - + 来表示score最小值 和score最大值  
ZLEXCOUNT key [min [max  

Lex（表示依照字典顺序 - 有序集key中的成员score必须相同 否则可能会输出错误数据）  
正序  
ZRangeByLex key min max 返回指定区成员，并按字典正序排序 （排序的成员score必须相同，不然可能获取的结果并不准确）  
ZRangeByScore key min max [withscores] [limit offset count]  
返回指定范围内的元素 并按score从小到大排序   
ZRangByScore key -inf +inf 显示从最低分到最高分所有数据  

排名  
ZRank key member 返回成员在有序集key中排名 - 按score从小到大的顺序排列   
排序的顺序 从0起 即score最小的成员为0 score最大的成员为 成员个数-1   

删除  
ZRem (remove) key member [member...] 删除有序集key中的成员  
ZRemRangeByLex key min max 删除名称按字典由低到高排序的所有成员（删除的有序集score必须相同）  
ZRemRangeByRank key start stop 删除有序集合中 指定排名区间内的成员  
ZRemRangeByScore key min max 删除后有序集中 score介于 min 和 max 之间的数  
使用Score 由于分数是浮点数 float 在指定区域时最低分 ~ 最高分 -inf +inf 而非 - +   

降序  
ZRevRange (rev = reverse) key start stop [withscores] 按分数降序输出成员  
ZRevRangeByLex key max min 返回指定范围元素 按字典顺序排列  
ZRevRangeByScore key max min 返回介于score之间的元素 分数有高到底排序  
ZRevRank key member 返回成员分数降序的排名  

ZScore key member 返回成员的分数  


