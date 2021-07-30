package Vm

/**
	栈帧
 */
type Frame struct {

	//pc指针 一定不会指向数据字节,只能在指令字节
	pc int

	//呼叫的栈帧
	sendFrame *Frame;

	//操作栈
	statck *FStack

	//localVal<int,*FMap>  每个数据类型都有自己的槽
	localVal *FMap

	//指令原始字节码
	code []byte

	codeList map[uint8]IFObject

	codes []*OpcodeStruct

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
func (me *Frame ) hasMoreOpcode() *OpcodeStruct {
	if( me.pc+1 >= len(me.codes) ){
		return nil;
	}
	me.pc++;
	return me.codes[me.pc]
}