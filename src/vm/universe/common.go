package Universe

import (
	"vm/klass"
	"vm/type"
)

var TRUE *ftype.FString
var FALSE *ftype.FString
var NONE *ftype.FString

var KlassBean map[string]klass.IKlass

func initKlass()  {
	KlassBean = make( map[string]klass.IKlass)
	KlassBean["nada.lib.NativeMethodKlass"] 	= new(klass.NativeMethodKlass).Init()
	KlassBean["nada.lib.VirtualMethodKlass"] 	= new(klass.VirtualMethodKlass).Init()
	KlassBean["nada.lib.StringKlass"] 			= new(klass.StringKlass).Init()
}

func CommonInit()  {
	initKlass()
	TRUE 	= ftype.NewFString().SetVal("true")
	FALSE 	= ftype.NewFString().SetVal("false")
	NONE 	= ftype.NewFString().SetVal("none")
}




