package vm

import (
	"fmt"
)

type Interpreter struct {

	//当前执行的栈帧
	frame *Frame;
	//虚拟机实例
	vms *Vms;
	//栈的调用深度
	stackDeep int;

	ins *Instruction;

}

func (me *Interpreter)Build() *Interpreter{
	me.ins = NewInstruction()
	return me
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

/**
	解析器的方法，全局查找klass
	但klass是一个全局的容器
	所以 出现多线程时，可能多个解释器同时查找同一个全局Klass管理器
 */
func (me *Interpreter)findKlass(e IFObject) IKlass {
	if s,ok:=e.(*FString);ok{
		if s2:=vms.metaKlass.get( s.GetVal() ) ;s2!=nil{
			return s2.(IKlass)
		}else{//尝试加载新的KLASS

		}
	}
	return nil;
}

/**
	根据限定符加载Klass
	加载klass的时机 https://baijiahao.baidu.com/s?id=1662931890502685014&wfr=spider&for=pc
 */
func (me *Interpreter)loadKlassByQualifier(e IFObject) IKlass {
	if s,ok:=e.(*FString);ok{
		if s2:=vms.metaKlass.get( s.GetVal() ) ; s2==nil{
			//尝试加载新的KLASS


		}else{
			return s2.(IKlass)
		}
	}
	return nil;
}



func (me *Interpreter)BuildFrame(callable IFObject, argType *NSMethodArgType, args *FArray){
	me.buildFrame(callable, argType , args)
}


func (me *Interpreter)buildFrame(callable IFObject, argType *NSMethodArgType, args *FArray){
	me.stackDeep++;
	if callable.GetKlass().getName() == KLS_T.NativeMethodKlass{
		retOb := callable.(*FMethod).Call(args)
		if argType.ret.GetSize() > 0{
				me.frame.StackPush(retOb)
		}
	}else if callable.GetKlass().getName() == KLS_T.MethodKlass{
		//虚方法
		frame:=NewFrame(callable)
		frame.SetSendFrame(me.frame)
		me.frame = frame;
	}

}

func (me *Interpreter)leaveFrame(){
	me.frame = me.frame.GetSendFrame()
	me.stackDeep--;
}



func (me *Interpreter)exec(){

	for{
		op := me.frame.HasMoreOpcode()

		if op==nil{
			goto exception
		}


			fmt.Printf("opcode n:%d [0x%x] %s \n", op.n,op.op,OPCODE_DESC[op.op])
			switch op.op {
			case OP.ALOAD:
				soltNum:=int(op.oper.(*FUint16).getVal())
				v:=me.frame.GetLocal(soltNum).(IFObject)
				me.frame.StackPush(v)
			case OP.ALOAD_0:
			case OP.ALOAD_1:
				//me.frame.statck.push(me.frame.localVal.get(1).(IFObject))
				me.frame.StackPush(me.frame.GetLocal(1).(IFObject))
			case OP.ALOAD_2:
				me.frame.StackPush(me.frame.GetLocal(2).(IFObject))
			case OP.INVOKESPECIAL:
			case OP.INVOKEVIRTUAL:
				//The parameter type is displayed first
				if opArg,ok:=op.oper.(*FArray); ok{
					//argtype
					argType := opArg.pop()
					opArg.pop()
					mthodName	:=	opArg.pop()
					//qualifier
					opArg.pop()
					//Processing of eucalyptus
					atype:=argType.(*NSMethodArgType);
					args :=  NewFArray()
					for i:=0;i<atype.arg.GetSize();i++{
						args.add(me.frame.StackPop())
					}
					//Instance objects
					callOb	:= me.frame.StackPop()
					mt := callOb.getMethod(mthodName.(*FString).GetVal())
					var retOb  = mt.(*FMethod).Call(args);
					//  if has return values ,need push value to stack
					if atype.ret.GetSize() > 0 {
						me.frame.StackPush(retOb)
					}
				}

			case OP.RETURN:
				fmt.Println("OP_RETURN")
			case OP.ICONST_0://nada:压入int到栈 jvm:压入uint8到栈
				me.frame.StackPush(NewFInt(0))
			case OP.ICONST_1://nada:压入int到栈 jvm:压入uint8到栈
				me.frame.StackPush(NewFInt(1))
			case OP.ICONST_2://nada:压入int到栈 jvm:压入uint8到栈
				me.frame.StackPush(NewFInt(2))
			case OP.ICONST_3://nada:压入int到栈 jvm:压入uint8到栈
				me.frame.StackPush(NewFInt(3))
			case OP.ICONST_4://nada:压入int到栈 jvm:压入uint8到栈
				me.frame.StackPush(NewFInt(4))
			case OP.ICONST_5://nada:压入int到栈 jvm:压入uint8到栈
				me.frame.StackPush(NewFInt(5))
			case OP.LDC:
				me.frame.StackPush(op.oper)
			case OP.ASTORE_0 ,OP.ASTORE_3: //don't need operands
			case OP.FSTORE://为压缩空间 nada 本地变量编号全部为2字节

			case OP.LDC_W:
				me.frame.StackPush(op.oper)
				fmt.Println(op.oper)
			case OP.INVOKESTATIC:
				me.ins.opGetStatic.Execute(me.frame,op,me)
			case OP.ASTORE:

				soltNum:=int(op.oper.(*FUint16).getVal())
				ob:=me.frame.StackPop()
				me.frame.SetLocal(soltNum,ob)
			case OP.ASTORE_1:
				ob:= me.frame.StackPop()
				me.frame.SetLocal(1,ob)
			case OP.ASTORE_2:
				ob:=me.frame.StackPop()
				me.frame.SetLocal(2,ob)
			case OP.AASTORE:
				//ob:=me.frame.StackPop()
				//me.frame.SetLocal(2,ob)
			case OP.LSTORE:
				soltNum:=int(op.oper.(*FUint16).getVal())
				ob:=me.frame.StackPop()
				me.frame.SetLocal(soltNum,ob)
			case OP.ISTORE:
				ob:=me.frame.StackPop()
				soltNum:=int(op.oper.(*FUint16).getVal())
				me.frame.SetLocal(soltNum,ob)

			case OP.ISTORE_0:
				ob:=me.frame.StackPop()
				me.frame.SetLocal(0,ob)
			case OP.ISTORE_1:
				ob:=me.frame.StackPop()
				me.frame.SetLocal(1,ob)
			case OP.ISTORE_2:
				ob:=me.frame.StackPop()
				me.frame.SetLocal(2,ob)
			case OP.ISTORE_3:
				ob:=me.frame.StackPop()
				me.frame.SetLocal(3,ob)
			case OP.BIPUSH://1 byte to int
				slotNum:=op.oper.(*FInt).getVal()
				me.frame.StackPush(NewFInt(slotNum))
			case OP.SIPUSH:

			case OP.ILOAD:

			case OP.IF_CMPGE:

			case OP.NEW:

			case OP.DUP:


			case OP.INVOKEDYNAMIC:

			case OP.POP,OP.POP2://iinc index(uint8) const(int8) to int

			case OP.IINC:

			case OP.GOTO:
			case OP.GETSTATIC:
				//类没有加载的话，可以在此处加载 所以 要验证
				if opArg,ok:=op.oper.(*FArray); ok{
					//实例类型
					insType		:= 	opArg.pop()
					//实例名称
					ins			:= 	opArg.pop()
					//这个实例所在的Klass
					onwKlass	:=	opArg.pop()
					ob 		:= vms.metaKlass.get(onwKlass.(*FString).GetVal()).(IKlass)
					insOb 	:= ob.getStaticField(  ins.(*FString).GetVal()  )

					me.frame.StackPush(insOb)
					fmt.Println(insType)

				}
			default:


			}



	}

	exception:{
		me.leaveFrame();
		//如果不是最底部的栈帧 需要继续执行
		if me.stackDeep > 0{
			me.exec();
		}
	}



}