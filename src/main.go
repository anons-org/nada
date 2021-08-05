
package main

import (
	"fmt"
	"nada/src/vm"
)


func main() {


	//Utils.MiscIns().GetDescStrData("(DIZVLjava/lang/String;)I");

	//s := []string{"a", "b", "c", "d"}
	//
	//fmt.Print(s[:len(s)-1])

	vm.CommonInit()
	vm.Start()

	x := 1

	fmt.Print(x)

	for {
	}

}

//func main() {
//	str := "月色真美，风也温柔，233333333，~！@#"  //go字符串编码为utf-8
//	fmt.Println("before convert:", str)   //打印转换前的字符串
//	fmt.Println("isUtf8:", isUtf8([]byte(str)))   //判断是否是utf-8
//	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))  //使用官方库将utf-8转换为gbk
//	fmt.Println("gbk直接打印会出现乱码:", string(gbkData))   //乱码字符串
//	fmt.Println("isGBK:", isGBK(gbkData)) //判断是否是gbk
//	utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(gbkData) //将gbk再转换为utf-8
//	fmt.Println("isUtf8:", isUtf8(utf8Data) )  //判断是否是utf-8
//	fmt.Println("after convert:", string(utf8Data))   //打印转换前的字符串
//}
