package Vm

import (
	"container/list"
	"strings"
	VmCore "vm/object"
	"vm/type"
)

type FrameObject struct {
	//常量池
	constanPool []ftype.IFObject
	stack list.List
}


func (me *FrameObject)eval(){

}

func (me *FrameObject)Init(){
	me.stack.Init()
}




//把字符串指令解析成可执行指令
func (me *FrameObject)PsarserOpcodeOfString(intr string) []ftype.IFObject {


	parts := strings.Split(intr," ")

	opi := me.GetOPIdxByValStr( parts[0] )

	//根据指令生成不同类型操作数
	switch opi{

	case VmCore.OP.ICONST_0:


	case VmCore.OP.ICONST_1:



	default:
		

	}

	return nil
}


/**
根据指令字符串获取指令索引
*/
func (me *FrameObject)GetOPIdxByValStr(val string) int16{
	for _, value := range VmCore.OPD {
		if value.ValStr == val{
			return value.Intr
		}
	}
	return 0
}

