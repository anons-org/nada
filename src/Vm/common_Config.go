package Vm

import "strings"


var OPCODE_DESC = map[uint8]string{
	OP.ICONST_0: "ICONST_0",
	OP.ICONST_1: "ICONST_1",
	OP.ICONST_2: "ICONST_2",
	OP.ICONST_3: "ICONST_3",
	OP.ICONST_4: "ICONST_4",
	OP.ICONST_5: "ICONST_5",
	OP.BIPUSH: "BIPUSH",
	OP.SIPUSH: "SIPUSH",
	OP.LDC:"LDC",
	OP.LDC_W:"LDC_W",
	OP.ILOAD:"ILOAD",


	OP.IF_CMPEQ:"IF_CMPEQ",
	OP.IF_CMPNE:"IF_CMPNE",
	OP.IF_CMPLT:"IF_CMPLT",
	OP.IF_CMPGE:"IF_CMPGE",
	OP.IF_CMPGT:"IF_CMPGT",
	OP.IF_CMPLE:"IF_CMPLE",

	//0x2a aload_0    将第一个引用类型本地变量推送至栈顶 在非静态方法中， aload_0 表示对this的操作，在static 方法中，aload_0表示对方法的第一参数的操作。
	OP.ALOAD_0: "ALOAD_0",
	OP.ALOAD_1: "ALOAD_1",
	OP.ISTORE: "ISTORE",
	OP.LSTORE: "LSTORE",
	OP.FSTORE: "FSTORE",
	OP.ASTORE: "ASTORE",
	OP.ISTORE_0: "ISTORE_0",
	OP.ISTORE_1: "ISTORE_1",
	OP.ISTORE_2: "ISTORE_2",
	OP.ISTORE_3: "ISTORE_3",
	OP.ASTORE_0: "ASTORE_0",
	OP.ASTORE_1: "ASTORE_1",
	OP.ASTORE_2: "ASTORE_2",
	OP.ASTORE_3: "ASTORE_3",
	OP.POP:"POP",
	OP.POP2:"POP2",
	OP.DUP:"DUP",
	OP.IINC:"IINC",
	OP.GOTO:"GOTO",
	OP.RETURN:"RETURN",
	OP.GETSTATIC:"GETSTATIC",
	OP.INVOKEDYNAMIC:"INVOKEDYNAMIC",
	OP.INVOKEVIRTUAL:"INVOKEVIRTUAL",
	OP.INVOKESPECIAL:"INVOKESPECIAL",
	OP.INVOKESTATIC:"INVOKESTATIC",
	OP.NEW:"NEW",
}



var Z_TYPE  = map[string]string{
	"Z":"Z",
	"B":"B",
	"C":"C",
	"S":"S",
	"I":"I",
	"J":"J",
	"F":"F",
	"D":"D",
	"V":"V",
}


//opcode index
var OP = struct {
	ICONST_0,
	ICONST_1,
	ICONST_2,
	ICONST_3,
	ICONST_4,
	ICONST_5,
	BIPUSH,
	SIPUSH,
	LDC,
	LDC_W,
	ILOAD,

	IF_CMPEQ,
	IF_CMPNE,
	IF_CMPLT,
	IF_CMPGE,
	IF_CMPGT,
	IF_CMPLE,

	ALOAD_0,
	ALOAD_1,
	ISTORE,
	LSTORE,
	FSTORE,
	ASTORE,
	ISTORE_0,
	ISTORE_1,
	ISTORE_2,
	ISTORE_3,
	ASTORE_0,
	ASTORE_1,
	ASTORE_2,
	ASTORE_3,
	POP,
	POP2,
	DUP,
	IINC,
	GOTO,
	RETURN,
	GETSTATIC,
	INVOKEDYNAMIC,
	INVOKEVIRTUAL,
	INVOKESPECIAL,
	INVOKESTATIC,
	NEW uint8
}{
	ICONST_0:0x03,
	ICONST_1:0x04,
	ICONST_2:0x05,
	ICONST_3:0x06,
	ICONST_4:0x07,
	ICONST_5:0x08,
	BIPUSH:0x10,
	SIPUSH:0x11,
	LDC:0x12,
	LDC_W:0x14,
	ILOAD:0x15,

	IF_CMPEQ:0x9f,
	IF_CMPNE:0xa0,
	IF_CMPLT:0xa1,
	IF_CMPGE:0xa2,
	IF_CMPGT:0xa3,
	IF_CMPLE:0xa4,

	ALOAD_0:0x2A,
	ALOAD_1:0x2b,
	ISTORE :0x36,
	LSTORE :0x37,
	FSTORE:0x38,
	ASTORE:0x3a,
	ISTORE_0:0x3b,
	ISTORE_1:0x3c,
	ISTORE_2:0x3d,
	ISTORE_3:0x3e,
	ASTORE_0:0x4b,
	ASTORE_1:0x4c,
	ASTORE_2:0x4d,
	ASTORE_3:0x4e,
	POP:0x57,
	POP2:0x58,
	DUP:0x59,
	IINC:0x84,
	GOTO:0xa7,
	RETURN:0xb1,
	GETSTATIC:0xb2,
	INVOKEDYNAMIC:0xba,
	INVOKEVIRTUAL:0xb6,
	INVOKESPECIAL:0xb7,
	INVOKESTATIC:0xb8,
	NEW:0xbb,
}

//if_icmpeq = 159 (0x9f)
//
//if_icmpne = 160 (0xa0)
//
//if_icmplt = 161 (0xa1)
//
//if_icmpge = 162 (0xa2)
//
//if_icmpgt = 163 (0xa3)
//
//if_icmple = 164 (0xa4)
//object type---------------------
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


type MethodType int

const (
	METHOD_TYPE_NATIVE      MethodType = 1
	METHOD_TYPE_VIRTUAL     MethodType = 2
	METHOD_TYPE_FUNC      	MethodType = 3

)



var KLS_T = struct {
	StringKlass,
	NativeMethodKlass string
}{
	StringKlass:"stringKlass",
	NativeMethodKlass:"NativeMethodKlass",
}



var BUILTN = struct {
	JAVA_LANG_STRING   string
	NADA_STRING string
	NADA_VIRTUAL string
	NADA_NATIVE string
	NADA_TYPE string
	NADA_FLOAT string
	NADA_UINT16 string
	NADA_INT string
	NADA_BYTE string
	NADA_ARRAY string
	NADA_STACK string
	NADA_OBJECT string

}{

	NADA_NATIVE: "nada.core.Native",
	NADA_VIRTUAL: "nada.core.Method",
	NADA_STRING:"nada.core.String",
	NADA_FLOAT:"nada.core.Float",
	NADA_UINT16:"nada.core.Uint16",
	NADA_INT:"nada.core.Int",
	NADA_BYTE:"nada.core.Byte",
	NADA_TYPE:"nada.core.Type",
	NADA_ARRAY:"nada.core.Array",
	NADA_STACK:"nada.core.Stack",
	NADA_OBJECT: "nada.core.Object",
	JAVA_LANG_STRING:"nada.core.String",
}





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
////
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








