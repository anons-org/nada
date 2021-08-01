package Vm

import "fmt"

type Vms struct {

	//所有的Klass元类
	metaKlass *FMap

	//项目的class目录 默认当前的lib目录
	classPath string

	//所有的解释器

}

func (me *Vms) start() {
	//初始化参数 klass等
	var fileName = "nada/untitled/target/classes/Test.class"
	classN := NewClassByteFile()
	classN.Load(fileName).Parser()
	//执行的时候，一定是有运行目录的 Test.class 在执行的根目录
	main := me.metaKlass.get("Test").(IKlass)
	f := main.getStaticMethod("main").(*FMethod)

	go func() {
		interp := new(Interpreter)
		interp.run(f)
	}()
	fmt.Print("strap!\n")

}

func testCall(arg *FArray) IFObject {
	fmt.Print("testCall native")
	return NONE
}

func (me *Vms) Build() *Vms {
	maps := make(map[string]IKlass)
	//创建内置类型
	maps["BUILTN.NADA_TYPE"] = new(TypeKlass).Init()
	me.metaKlass = NewFMap()
	me.metaKlass.set(BUILTN.NADA_TYPE, maps["BUILTN.NADA_TYPE"])
	me.metaKlass.set(BUILTN.NADA_NATIVE, new(NativeMethodKlass).Init())
	me.metaKlass.set(BUILTN.NADA_VIRTUAL, new(MethodKlass).Init())
	me.metaKlass.set(BUILTN.NADA_STRING, new(StringKlass).Init())
	me.metaKlass.set(BUILTN.NADA_FLOAT, new(FloatKlass).Init())
	me.metaKlass.set(BUILTN.NADA_UINT16, new(Uint16Klass).Init())
	me.metaKlass.set(BUILTN.NADA_INT, new(IntKlass).Init())
	me.metaKlass.set(BUILTN.NADA_BYTE, new(ByteKlass).Init())
	me.metaKlass.set(BUILTN.NADA_ARRAY, new(ArrayKlass).Init())
	me.metaKlass.set(BUILTN.NADA_STACK, new(StackKlass).Init())
	me.metaKlass.set(BUILTN.NADA_OBJECT, new(ObjectKlass).Init())


	/**
			生成内置基础类 过程 java.lang.System

			java/io/PrintStream




	 */

	printStream := createKlass("PrintStream", "java/io/PrintStream")
	printStream.setStaticMethod("print", NewNativeFMethod("print", testCall))

	//注册到虚拟机
	addKlassToVm("java/io/PrintStream", printStream)

	systemKlass := createKlass("System", "java/lang/System")

	//out 是实例....需要生成实例。。。
	//systemKlass.addStaticField("out",))
	ob := createKlassIns(printStream.GetTypeObject(),NewFArray())
	systemKlass.setStaticField("out",ob)



	ob.GetKlass().getStaticMethod("printf").GetKlass()

	fmt.Print(ob)
	//添加System
	addKlassToVm("java/lang/System", systemKlass)

	//fmt.Print(systemKlass.getStaticDict("out").(*FMethod).Call(  NewFArray()  ))

	//addKlassToVm( "java/io/PrintStream", createKlass("PrintStream","java/io/PrintStream") )

	return me
}
