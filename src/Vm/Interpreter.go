package Vm

import (
	"fmt"
)

type Interpreter struct {

	//当前执行的栈帧
	frame *Frame;
	//虚拟机实例
	vms *Vms;
}



func (me *Interpreter)eval(){

}

/**
	设置栈帧
 */
func (me *Interpreter)setFrame(f *Frame){
	me.frame = f;
}


func (me *Interpreter)run(f *FMethod){
	me.frame = NewFrame(f)
	me.exec();
}




func (me *Interpreter)exec(){



	for{

		op:=me.frame.hasMoreOpcode()

		if op!=nil{



			fmt.Printf("opcode n:%d [0x%x] %s \n", op.n,op.op,OPCODE_DESC[op.op])

			switch op.op {

			case OP.ALOAD_0:

			case OP.INVOKESPECIAL:

			case OP.RETURN:

			case OP.ICONST_0://nada:压入int到栈 jvm:压入uint8到栈

			case OP.ICONST_1://nada:压入int到栈 jvm:压入uint8到栈

			case OP.LDC:

			case OP.ASTORE_0, OP.ASTORE_1, OP.ASTORE_2, OP.ASTORE_3,OP.ISTORE_0,OP.ISTORE_1,
				OP.ISTORE_2,OP.ISTORE_3: //don't need operands

			case OP.FSTORE://为压缩空间 nada 本地变量编号全部为2字节

			case OP.LDC_W:

			case OP.INVOKESTATIC:

			case OP.ASTORE:

			case OP.LSTORE:

			case OP.ISTORE:

			case OP.BIPUSH://nada:压入int到栈 jvm:压入uint8到栈

			case OP.SIPUSH:

			case OP.ILOAD:

			case OP.IF_CMPGE:

			case OP.NEW:

			case OP.DUP:


			case OP.INVOKEDYNAMIC:

			case OP.POP,OP.POP2://iinc index(uint8) const(int8) to int

			case OP.IINC:

			case OP.GOTO:

			default:


			}


		}else{
			break
		}
	}

}