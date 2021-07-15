package Vm

var TRUE *FString
var FALSE *FString
var NONE *FString

var KlassBean map[string]IKlass

func initKlass()  {
	KlassBean = make( map[string]IKlass)
	KlassBean["nada.lib.NativeMethodKlass"] 	= new(NativeMethodKlass).Init()
	KlassBean["nada.lib.VirtualMethodKlass"] 	= new(VirtualMethodKlass).Init()
	KlassBean["nada.lib.StringKlass"] 			= new(StringKlass).Init()
}

func CommonInit()  {
	initKlass()
	TRUE = NewFString().SetVal("true")
	FALSE = NewFString().SetVal("false")
	NONE = NewFString().SetVal("none")
}




