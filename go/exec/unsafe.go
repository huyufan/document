package main

import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

func main() {
	fmt.Println(os.Args)
	// 解析标准格式
	t1, err := time.Parse("2006-01-02 15:04:05", "2023-10-25 14:30:00")
	if err != nil {
		fmt.Println("解析错误:", err)
		return
	}
	fmt.Println("解析结果:", t1)

	// 解析日期
	t2, _ := time.Parse("2006-01-02", "2023-10-25")
	fmt.Println("仅日期:", t2)

	// 解析时间
	t3, _ := time.Parse("15:04:05", "14:30:00")
	fmt.Println("仅时间:", t3)
	utcTime, _ := time.Parse("2006-01-02 15:04:05", "2023-10-25 14:30:00")
	fmt.Println("UTC时间:", utcTime)

	// 转换为本地时区
	localTime := utcTime.Local()
	fmt.Println("本地时间:", localTime)

	var f float64 = 3.1415
	fpint := &f

	iInt := (*int64)(unsafe.Pointer(fpint))

	fmt.Println(*iInt)

	fmt.Printf("float64: %f\n", f)
	fmt.Printf("底层二进制: 0x%x\n", *iInt)
	*iInt = 0x400921FB54442D18 // 这是 π 的 IEEE 754 表示
	fmt.Printf("修改后的 float64: %f\n", f)
}
