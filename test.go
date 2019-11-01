package main

import (
	"fmt"
	"time"
	// "github.com/gin-gonic/gin"
)

// func main()  {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"msg": "pong",
// 		})
// 	})
	
// 	r.Run()
// }

func main() {
testData := "2006-01-02 15:04:05"

nowStr := time.Now().Format("2006-01-02 15:04:05")
f1 := func() bool {
return testData > nowStr
}

f2 := func() bool {
nowStr := time.Now().Format("2006-01-02 15:04:05")
return testData > nowStr
}

f3 := func() bool {
t, err := time.Parse("2006-01-02 15:04:05", testData)
if err != nil {
panic(err)
}
return time.Now().After(t)
}

T("字符串比较", f1)
T("now 转换为字符串后和字符串比较", f2)
T("字符串解析为 time 后比较", f3)

}

func T(name string, f func() bool) {
s := time.Now()
for i := 0; i < 1000000; i++ {
f()
}
fmt.Printf("%v 耗时 %v\r\n", name, time.Now().Sub(s))
}