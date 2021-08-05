package vm

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

func (ne *Frame) name()  {

}

/**
	是否还有指令可执行
	操作数最好在生成Klass的时候就转译为VM内置对象
	这样多少能提升一点性能
 */
func (me *Frame) hasMoreOpcode() *OpcodeStruct {



	if( me.pc  > len(me.codes)-1 ){
		return nil;
	}

	code:=me.codes[me.pc]
	me.pc++;

	return code
}
//-------------------------------------------public------------------------------------
func (me *Frame) HasMoreOpcode() *OpcodeStruct {
	return me.hasMoreOpcode()
}

func (me *Frame) GetLocal(k interface{}) interface{} {
	return me.localVal.Get(k)
}

func (me *Frame) SetLocal(k interface{},v interface{}) {
	 me.localVal.Set(k,v)
}


func (me *Frame) StackPop() IFObject {
	return me.statck.Pop()
}

func (me *Frame) StackPush(v IFObject)  {
	me.statck.Push(v)
}


func (me *Frame) SetSendFrame(sendFrame *Frame)  {
	me.sendFrame = sendFrame
}

func (me *Frame) GetSendFrame() *Frame {
	return me.sendFrame
}


