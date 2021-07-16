package Vm

import (
	"container/list"
	"strings"
)

type FrameObject struct {
	//常量池
	constanPool []IFObject
	stack list.List
}


func (me *FrameObject)eval(){

}

func (me *FrameObject)Init(){
	me.stack.Init()
}




//把字符串指令解析成可执行指令
func (me *FrameObject)PsarserOpcodeOfString(intr string) []IFObject {


	parts := strings.Split(intr," ")

	opi := me.GetOPIdxByValStr( parts[0] )

	//根据指令生成不同类型操作数
	switch opi{

	case OP.ICONST_0:


	case OP.ICONST_1:



	default:
		

	}

	return nil
}


/**
根据指令字符串获取指令索引
*/
func (me *FrameObject)GetOPIdxByValStr(val string) int16{
	for _, value := range OPD {
		if value.ValStr == val{
			return value.Intr
		}
	}
	return 0
}
