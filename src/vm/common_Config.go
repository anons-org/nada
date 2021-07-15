package Vm

import (
	"strings"
)

//指令
var OPD = []struct {
	Intr int16
	ValStr string
}{
	{0x03, "ICONST_0"},
	{0x04, "ICONST_1"},
}


//指令索引
var OP = struct {
	ICONST_0 int16
	ICONST_1   int16

}{
	ICONST_0:0x03,
	ICONST_1:0x04,

}





//对象类型---------------------
var G_OT_DEF = []struct {
	Id   int
	Name string
}{
	{0, "int"},
	{1, "string"},
	{2, "bool"},
	{3, "function"},
	{4, "method"},
}

var G_OT = struct {
	INT   int
	STRING int
}{
	INT:1,STRING: 2,
}




//
//func PsarserOpcodeOfString(intr string) []IFObject{
//
//
//	parts := strings.Split(intr," ")
//
//	opi := GetOPIdxByValStr( parts[0] )
//
//	//根据指令生成不同类型操作数
//	switch opi{
//
//		case OP.ICONST_0:
//
//
//		case OP.ICONST_1:
//
//
//
//		default:
//
//
//	}
//
//}
//
//
//
//func buildOPICONST_0Data(op int16, v string){
//
//
//
//}


func ConverStrOp(str string) []string{

	opcodeArr := strings.Split(str,"\n")

	var vals []string
	for _, v := range opcodeArr {
		v = strings.Replace(v, "\t", "", -1)
		//fmt.Print(v)
		vals = append(vals, v)
	}
	return vals
}








//
//
//
//
//public static final int OP_ICONST_0 = 0x03;
//public static final int OP_ICONST_1 = 0x04;
//public static final int OP_ICONST_2 = 0x05;
//public static final int OP_ICONST_3 = 0x06;
//public static final int OP_ICONST_4 = 0x07;
//public static final int OP_ICONST_5 = 0x08;
//
//// 将一个byte型常量值推送至栈顶
//public static final int OP_BIPUSH = 0x10;
//
//// 0x11 sipush 将一个短整型常量值(-32768~32767)推送至栈顶
//public static final int OP_SIPUSH = 0x11;
//
//// 将int、float或String型常量值从常量池中推送至栈顶
//public static final int OP_LDC = 0x12;
//public static final int OP_LDC2W = 0x14;
//
//// 0x15 iload 将指定的int型本地变量推送至栈顶
//public static final int OP_ILOAD = 0x15;
//
//// 0x16 lload 将指定的long型本地变量推送至栈顶
//
//public static final int OP_LLOAD = 0x16;
//
//// 0x1a iload_0 将第一个int型本地变量推送至栈顶
//// 0x1b iload_1 将第二个int型本地变量推送至栈顶
//// 0x1c iload_2 将第三个int型本地变量推送至栈顶
//public static final int OP_ILOAD_0 = 0x1a;
//public static final int OP_ILOAD_1 = 0x1b;
//public static final int OP_ILOAD_2 = 0x1c;
//
//// 0x1d iload_3 将第四个int型本地变量推送至栈顶
//public static final int OP_ILOAD_3 = 0x1d;
//
//// 0x19 aload 将指定的引用类型本地变量推送至栈顶
//public static final int OP_ALOAD = 0x19;
//
//// 0x21 lload_3 将第四个long型本地变量推送至栈顶
//public static final int OP_LLOAD_3 = 0x21;
//
////	0x2a aload_0    将第一个引用类型本地变量推送至栈顶
//
//public static final int OP_ALOAD_0 = 0x2a;
//
////	0x9f if_icmpeq    比较栈顶两int型数值大小，当结果等于0时跳转
////	0xa0 if_icmpne    比较栈顶两int型数值大小，当结果不等于0时跳转
////	0xa1 if_icmplt    比较栈顶两int型数值大小，当结果小于0时跳转
////	0xa2 if_icmpge    比较栈顶两int型数值大小，当结果大于等于0时跳转
////	0xa3 if_icmpgt    比较栈顶两int型数值大小，当结果大于0时跳转
////	0xa4 if_icmple    比较栈顶两int型数值大小，当结果小于等于0时跳转
////	0xa5 if_acmpeq    比较栈顶两引用型数值，当结果相等时跳转
////	0xa6 if_acmpne    比较栈顶两引用型数值，当结果不相等时跳转
//public static final int OP_IFCMPGE = 0xa2;
//
//
//
////	0xa4 if_icmple    比较栈顶两int型数值大小，当结果小于等于0时跳转
//public static final int OP_IFCMPLE = 0xa4;
//
//
//// 0xa7 goto 无条件跳转
//
//public static final int OP_GOTO = 0xa7;
//
//// 0xb1 return 从当前方法返回void
//
//public static final int OP_RETURN = 0xb1;
////0xb3
//public static final int OP_PUT_STATIC = 0xb3;
//
////0xba invokedynamic
//public static final int OP_CALL_DYN = 0xba;
//
//// 0x36 istore 将栈顶int型数值存入指定本地变量
//public static final int OP_ISTORE = 0x36;
//
//// 0x37 lstore 将栈顶long型数值存入指定本地变量
//public static final int OP_LSTORE = 0x37;
//
//// 0x3a astore 将栈顶引用型数值存入指定本地变量
//public static final int OP_ASTORE = 0x3a;
//public static final int OP_ISTORE_1 = 0x3c;
//public static final int OP_ISTORE_2 = 0x3d;
//
//// 0x3e istore_3 将栈顶int型数值存入第四个本地变量
//
//public static final int OP_ISTORE_3 = 0x3e;
//
//
//
////0x40	lstore_1	将栈顶long型数值存入第二个本地变量
//public static final int OP_LSTORE_1 = 0x40;
//// 0x42 lstore_3 将栈顶long型数值存入第四个本地变量
//
//public static final int OP_LSTORE_3 = 0x42;
//
//// 0x57 pop 将栈顶数值弹出 (数值不能是long或double类型的)
//public static final int OP_POP = 0x57;
//
////0x58	pop2	将栈顶的一个（long或double类型的)或两个数值弹出（其它）
//
//public static final int OP_POP2 = 0x58;
//
//
//// 0x59 dup 复制栈顶数值并将复制值压入栈顶
//
//public static final int OP_DUP = 0x59;
//
//// 0x60 iadd 将栈顶两int型数值相加并将结果压入栈顶
//public static final int OP_IADD = 0x60;
//
//// 0x65 lsub 将栈顶两long型数值相减并将结果压入栈顶
//public static final int OP_LSUB = 0x65;
//
//// 0x6d ldiv 将栈顶两long型数值相除并将结果压入栈顶
//
//public static final int OP_LDIV = 0x6d;
//
//// 0x84 iinc 将指定int型变量增加指定值（i++, i--, i+=2）
//public static final int OP_IINC = 0x84;
//
//// 0x9e ifle 当栈顶int型数值小于等于0时跳转
//public static final int OP_IFLE = 0x9e;
//
//// 从当前方法返回int
//public static final int OP_IRETURN = 0xac;
//
//// 0xb2 getstatic 获取指定类的静态域，并将其值压入栈顶
//
//public static final int OP_GET_STATIC = 0xb2;
//
///*
// * =============================================================================
// * =======
// *
// * 0xb6 invokevirtual 调用实例方法
// *
// * 方法的动态分派 方法的动态分派涉及到一个重要概念： 方法接收者。 invokevirtual字节码指令的多态查找流程（执行期）
// * 1、找到操作数栈顶的第一个元素，它所指向对象的实际类型 2、在实际对象中找对应的方法（test()方法），检查访问类型是否可以访问，找到了就调用。
// * 如果没找到，继续往上找。
// *
// *
// * 如果解析出来的方法有参数的话，在操作数栈上，所有的参数的值必须按照方法描述中规定的 数量，类型和顺序自上而下跟在objectref后面。自上而下的意思
// * 参数都在栈顶。
// *
// * object arg1 arg2 栈顶-> arg...
// *
// *
// *
// *
// * 比较方法重载（overload）与方法重写（overwrite),我们可以得到这样的结论 方法重载是静态的，是编译器行为；
// * 方法重写是动态的，是运行期行为。
// *
// * 网友翻译的官方原文 https://blog.csdn.net/lvvista/category_2440811.html
// *
// * =============================================================================
// * =======
// */
//
//public static final int OP_CALL_VIRTUAL = 0xb6;
//
//// 0xb7 invokespecial 调用超类构造方法，实例初始化方法，私有方法
//public static final int OP_CALL_SUPER = 0xb7;
//
//// 调用静态方法
//public static final int OP_CALL_STATIC = 0xb8;
//
//// 调用静态方法
//public static final int OP_NEW = 0xbb;








