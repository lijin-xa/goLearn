package _package
// time时间包

import (
	"fmt"
	"time"
)

type Time struct {
	wall uint64          // 距离公元1年1月1日00:00:00UTC的秒数
	ext  int64           // 表示纳秒
	//loc  *Location       // 代表时区
}

// UTC (Universal Time Coordinated，世界协调时间) 是以 GMT (Greewich Mean Time，格里尼治时间)这个时区为主
// UTC + 时区差 = 本地时间
// UTC + 8小时 = 北京时间
// time.UTC : UTC时间
// time.Local : 本地时间

func StartTime() {
	now := time.Now()  // 当前时间对象

	// 2020-09-22 11:45:44
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	timestamp := now.Unix()           // 时间戳 - 秒
	timestampExt := now.UnixNano()    // 时间戳 - 纳秒
	fmt.Println("timestamp: ", timestamp)
	fmt.Println("timestampExt: ", timestampExt)

	// 将时间戳转化为时间格式  相当于获得 time.Now()
	timeObj := time.Unix(timestamp, 0)
	fmt.Println("now time: ", timeObj)

	weekday := now.Weekday()  // 星期几
	fmt.Println("today is: ", weekday)

	operateTest()
}

// 时间操作函数测试
func operateTest() {

	// 两段时间间隔值
	now := time.Now()
	// 当前时间加一个小时后的时间
	later := now.Add(time.Hour)
	fmt.Println(later)

	// 定时器 time.Second 定时一秒执行任务
	ticker := time.Tick(time.Second)
	counter := 1
	for i := range ticker {
		fmt.Println("ticker ", i)
		if counter > 5 {
			break
		}
		counter += 1
	}

	// go语言的格式化 有些不同
	// 使用的是Go语言的诞生时间 2006年1月2日15点04分05秒 Mon - 星期 Jan - 月份
	var format = "2006-01-02 15:04:05 Mon Jan"
	fmt.Println(now.Format(format))

	// PM - 12小时制
	var format1 = "2006-01-02 15:04:05 PM Mon Jan"
	fmt.Println(now.Format(format1))

	var format2 = "15:04:05"
	fmt.Println(now.Format(format2))

	// 解析字符串
	// Parse() - 解析一个格式化的时间字符串并返回它代表的时间
	// 签名 func(layout, value string) (Time, error)
	// 签名 ParseInLocation(layout, value string, loc)

	var layout = "2006-01-02 15:04:05"
	var timeStr = "2020-09-22 14:51:55"

	// Parse 默认时区是UTC
	timeObj, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj)

	// ParseInLocation 可以设置loc时区
	// time.UTC   time.Local 得到CST时间
	timeObj, _ = time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj)
}