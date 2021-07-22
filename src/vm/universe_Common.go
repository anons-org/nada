package Vm

var TRUE *FString
var FALSE *FString
var NONE *FString

//访问标识
var ACC_PUBLIC 	= 0x01
var ACC_PRIVATE = 0x02
var ACC_STATIC 	= 0x08

//虚拟方法和原生方法（go函数）
var METHOD_NATIVE = 1
var  METHOD_VIRTUAL=2



var KlassBean map[string]IKlass

func initKlass()  {
	KlassBean = make( map[string]IKlass)
	KlassBean["nada.lib.NativeMethodKlass"] 	= new(NativeMethodKlass).Init()
	KlassBean["nada.lib.VirtualMethodKlass"] 	= new(VirtualMethodKlass).Init()
	KlassBean["nada.lib.StringKlass"] 			= new(StringKlass).Init()
}



func createKlass(name string) *Klass {

	kls:=new(Klass)
	kls.name = NewFString(name)
	typeObj := new(TypeObject)
	typeObj.SetOnwKlass(kls)
	kls.SetTypeObject( typeObj )
	kls.dict = make(map[string]IFObject)

	kls.staticDict = make(map[string]IFObject)
	kls.insDict = make(map[string]IFObject)

	return kls
}




func CommonInit()  {
	initKlass()
	TRUE = NewFString("true")
	FALSE = NewFString("false")
	NONE = NewFString("none")
}




