package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(printNumWith2(33.6787))
	byteArr := []byte{8, 0, 0, 0}
	printBytes(byteArr)
}

// 输出两位小数
func printNumWith2(num float64) string {
	// TODO 精度回头看
	var str = fmt.Sprintf("%.2f", num)
	return str
}

func printBytes(data []byte) string {
	// 序列化比较数据是否正确可能会用到
	//sign := md5.Sum(data)
	//signStr := fmt.Sprintf("%x", sign) //整型号以16进制方式显示

	hexStringData := hex.EncodeToString(data)
	fmt.Println(hexStringData)                    //byte 转 16进制 的结果
	hexData, _ := hex.DecodeString(hexStringData) //返回十六进制字符串所代表的字节数。
	fmt.Println(hexData)
	return ""
}
