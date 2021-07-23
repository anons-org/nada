package Vm

var TRUE *FString
var FALSE *FString
var NONE *FString
//ACC_PUBLIC	0x0001	方法是否为public
//ACC_PRIVATE	0x0002	方法是否为private
//ACC_PROTECTED	0x0004	方法是否为protected
//ACC_STATIC	0x0008	方法是否为static
//ACC_FINAL	0x0010	方法是否为final
//ACC_SYHCHRONRIZED	0x0020	方法是否为synchronized
//ACC_BRIDGE	0x0040	方法是否是有编译器产生的方法
//ACC_VARARGS	0x0080	方法是否接受参数
//ACC_NATIVE	0x0100	方法是否为native
//ACC_ABSTRACT	0x0400	方法是否为abstract
//ACC_STRICTFP 0x0800	方法是否为	strictfp
//ACC_SYNTHETIC	0x1000	方法是否是有编译器自动产生的

//访问标识
var ACC_PUBLIC 	= 0x0001
var ACC_PRIVATE = 0x0002
var ACC_STATIC 	= 0x0008
var ACC_SYNTHETIC = 0x1000

//虚拟方法和原生方法（go函数）
var METHOD_NATIVE = 1
var  METHOD_VIRTUAL=2



var KlassBean *FMap

func initKlass()  {

	KlassBean = NewFMap()
	KlassBean.set(BUILTN.NADA_TYPE,   	new(TypeKlass).Init())
	KlassBean.set(BUILTN.NADA_NATIVE, 	new(NativeMethodKlass).Init())
	KlassBean.set(BUILTN.NADA_VIRTUAL,	new(VirtualMethodKlass).Init())
	KlassBean.set(BUILTN.NADA_STRING,   new(StringKlass).Init())


}



func createKlass(name string) *Klass {

	kls:=new(Klass)
	kls.name = NewFString(name)
	typeObj := new(TypeObject)
	typeObj.SetOnwKlass(kls)
	kls.SetTypeObject( typeObj )
	kls.dict = make(map[string]IFObject)

	kls.staticDict = make(map[string]IFObject)
	kls.insDict    = make(map[string]IFObject)

	return kls
}




func CommonInit()  {
	initKlass()
	TRUE = NewFString("true")
	FALSE = NewFString("false")
	NONE = NewFString("none")
}




