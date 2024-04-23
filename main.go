package main

import (
	"fmt"
	"os"
	"util-go/models"
	"util-go/utils"
)

func init() {
	fmt.Println("init()")

}
func main() {
	fmt.Println("main()")

	// 读取 files/peoples.xml 文件
	peoples_bytes, err := os.ReadFile("files/peoples.xml")
	if err != nil {
		panic(err)
	}

	// 解析 xml 文件
	peoples := models.Peoples{}
	err = utils.ParseXML(peoples_bytes, &peoples)
	if err != nil {
		panic(err)
	}

	// 输出结果
	fmt.Println(peoples)

}
