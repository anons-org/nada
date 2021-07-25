package Vm

import "container/list"

/**
	栈帧
 */
type Frame struct {

	//pc指针
	pc int

	//呼叫的栈帧
	sendFrame *Frame;

	//操作栈
	statck list.List;

	//localVal<int,*FMap>  每个数据类型都有自己的槽
	localVal *FMap

	//指令
	code []byte

	//可能是Method 也可能是func
	fn IFObject

}

func (ne *Frame ) name()  {

}

/**
	是否还有指令可执行
	操作数最好在生成Klass的时候就转译为VM内置对象
	这样多少能提升一点性能
 */
func (ne *Frame ) hasMoreOpcode() bool {
	return true
}