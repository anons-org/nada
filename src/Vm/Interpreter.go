package Vm

import (
	"fmt"
)

type Interpreter struct {

	//当前执行的栈帧
	frame *Frame
	//虚拟机实例
	vms *Vms
}

func (me *Interpreter) eval() {

}

/**
设置栈帧
*/
func (me *Interpreter) setFrame(f *Frame) {
	me.frame = f
}

func (me *Interpreter) run(f *FMethod) {
	me.frame = NewFrame(f)
	me.exec()
}

/**
解析器的方法，全局查找klass
但klass是一个全局的容器
所以 出现多线程时，可能多个解释器同时查找同一个全局Klass管理器
*/
func (me *Interpreter) findKlass(e IFObject) IKlass {
	if s, ok := e.(*FString); ok {
		if s2 := vms.metaKlass.get(s.GetVal()); s2 != nil {
			return s2.(IKlass)
		} else { //尝试加载新的KLASS

		}
	}
	return nil
}

/**
根据限定符加载Klass
加载klass的时机 https://baijiahao.baidu.com/s?id=1662931890502685014&wfr=spider&for=pc
*/
func (me *Interpreter) loadKlassByQualifier(e IFObject) IKlass {
	if s, ok := e.(*FString); ok {
		if s2 := vms.metaKlass.get(s.GetVal()); s2 == nil {
			//尝试加载新的KLASS

		} else {
			return s2.(IKlass)
		}
	}
	return nil
}

func (me *Interpreter) exec() {

	for {
		op := me.frame.hasMoreOpcode()
		if op != nil {
			fmt.Printf("opcode n:%d [0x%x] %s \n", op.n, op.op, OPCODE_DESC[op.op])
			switch op.op {

			case OP.ALOAD_0:
			case OP.ALOAD_1:
				me.frame.statck.push(me.frame.localVal.get(1).(IFObject))
			case OP.INVOKESPECIAL:
			case OP.INVOKEVIRTUAL:
				//先弹出参数类型
				if opArg, ok := op.oper.(*FArray); ok {
					argT := opArg.pop()
					argStr := opArg.pop()
					mthodName := opArg.pop()
					qualifier := opArg.pop().(IFObject)
					arg := me.frame.statck.pop()
					//实例对象
					callOb := me.frame.statck.pop()
					mt := callOb.getMethod(mthodName.(*FString).GetVal())
					mt.(*FMethod).Call(NewFArray().add(argStr))
					fmt.Print(me.frame.localVal)
					fmt.Print(argT, argStr, mthodName, qualifier, callOb, arg, mt)

				}

			case OP.RETURN:

			case OP.ICONST_0: //nada:压入int到栈 jvm:压入uint8到栈

			case OP.ICONST_1: //nada:压入int到栈 jvm:压入uint8到栈

			case OP.LDC:
				me.frame.statck.push(op.oper)
			case OP.ASTORE_0, OP.ASTORE_2, OP.ASTORE_3, OP.ISTORE_0, OP.ISTORE_1,
				OP.ISTORE_2, OP.ISTORE_3: //don't need operands
			case OP.FSTORE: //为压缩空间 nada 本地变量编号全部为2字节

			case OP.LDC_W:

			case OP.INVOKESTATIC:

			case OP.ASTORE:
				ob := me.frame.statck.pop()
				me.frame.localVal.set(1, ob)
			case OP.ASTORE_1:
				ob := me.frame.statck.pop()
				me.frame.localVal.set(1, ob)
				fmt.Print(me.frame.localVal)

				//log.Print(string(debug.Stack()))
				//
				//fmt.Print(op.oper)
			case OP.LSTORE:

			case OP.ISTORE:

			case OP.BIPUSH: //nada:压入int到栈 jvm:压入uint8到栈

			case OP.SIPUSH:

			case OP.ILOAD:

			case OP.IF_CMPGE:

			case OP.NEW:

			case OP.DUP:

			case OP.INVOKEDYNAMIC:

			case OP.POP, OP.POP2: //iinc index(uint8) const(int8) to int

			case OP.IINC:

			case OP.GOTO:
			case OP.GETSTATIC:
				//类没有加载的话，可以在此处加载 所以 要验证

				if opArg, ok := op.oper.(*FArray); ok {
					//实例类型
					insType := opArg.pop()
					//实例名称
					ins := opArg.pop()
					//这个实例所在的Klass
					onwKlass := opArg.pop()
					ob := vms.metaKlass.get(onwKlass.(*FString).GetVal()).(IKlass)
					insOb := ob.getStaticField(ins.(*FString).GetVal())
					me.frame.statck.push(insOb)
					fmt.Print(me.frame.localVal, insType, onwKlass, ob, insOb)

				}

			default:

			}

		} else {
			break
		}
	}

}
