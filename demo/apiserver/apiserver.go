package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SR_File_Max_Bytes = 1024 * 1024 * 2
)
type user struct {
	ID   int
	Name string	`json:"名字"`
	Age  int
}

func main() {
	//关闭gin debug
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(JSONMiddleware())
	//提交生产api
	router.POST("/test", productionOrderHandler)
	//测试api是否正常访问
	router.GET("/ping", func(c *gin.Context) {
		//c.JSON(200, user{ID: 123, Name: "张三", Age: 20})

		data := Data{
			"pong",
			nil,
		}
		jsonData, _ := json.Marshal(data)
		fmt.Println(string(jsonData))
		c.JSON(http.StatusOK, data)
	})

	router.Run(":8081")
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func productionOrderHandler(c *gin.Context) {
	// eg: 如下代码，举例读取 post 中的 csv 文件
	rFile, err := c.FormFile("file")
	if err != nil {
		c.String(400, "文件格式错误")
		return
	}

	if rFile.Size > SR_File_Max_Bytes {
		c.String(400, "文件大小超过2M")
		return
	}

	file, err := rFile.Open()
	if err != nil {
		c.String(400, "文件格式错误")
		return
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	lineNumber := 0
	results := []BuglyCrashInfo{}
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		//line 就是每一行的内容
		fmt.Println(line)
		//line[0] 就是第几列
		fmt.Println(line[0])
		if lineNumber != 0 {
			result := BuglyCrashInfo{
				line[0],
				line[1],
				line[2],
				line[3],
				line[4],
				line[5],
				line[6],
				line[7],
				line[8],
				line[9],
				line[10],
				line[11],
				line[12],
				line[13],
				line[14],
				line[15],
				line[16],
			}
			results = append(results, result)
		}
		lineNumber += 1
	}
	message := "未知错误"
	data := Data{
		 message,
		 results,
	}
	btys, err := json.Marshal(data)
	if err != nil {
		c.String(200, "there is a wrong")
	}
	fmt.Println(string(btys))
	c.JSON(http.StatusOK, gin.H{
		"response": data,
	})
}

type Data struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type BuglyCrashInfo struct {
	CrashHash string `json:"crash_hash"`
	IssueId string	`json:"issue_id"`
	CrashId string	`json:"crash_id"`
	User_ID string	`json:"user_id"`
	Device_ID string	`json:"device_id"`
	Upload_time string	`json:"upload_time"`
	Crash_time string `json:"crash_time"`
	Bundle_id string `json:"bundle_id"`
	App_version string `json:"app_version"`
	Device_model string `json:"device_model"`
	Sys_version string `json:"sys_version"`
	ROM_des string `json:"rom_des"`
	CPU_struct string `json:"cpu_struct"`
	Is_jump string `json:"is_jump"`
	Memeory_size string `json:"memeory_size"`
	Store_size string `json:"store_size"`
	Sd_size string `json:"sd_size"`
}

