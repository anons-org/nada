package rtdata
//
//import (
//	"fmt"
//	"log"
//	"nada/src/rtdata"
//)
//
//type Interpretersss struct {
//
//	//当前执行的栈帧
//	frame *rtdata.Frame;
//	//虚拟机实例
//	vms *Vms;
//	//栈的调用深度
//	stackDeep int;
//
//	ins *Instruction;
//
//}
//
//func (me *Interpreter)Build() *Interpreter{
//	me.ins = NewInstruction()
//	return me
//}
//
//func (me *Interpreter)eval(){
//
//}
//
///**
//	设置栈帧
// */
//func (me *Interpreter)setFrame(f *rtdata.Frame){
//	me.frame = f;
//}
//
//
//func (me *Interpreter)run(f *FMethod){
//	me.frame = NewFrame(f)
//	me.exec();
//}
//
///**
//	解析器的方法，全局查找klass
//	但klass是一个全局的容器
//	所以 出现多线程时，可能多个解释器同时查找同一个全局Klass管理器
// */
//func (me *Interpreter)findKlass(e IFObject) IKlass {
//	if s,ok:=e.(*FString);ok{
//		if s2:=vms.metaKlass.get( s.GetVal() ) ;s2!=nil{
//			return s2.(IKlass)
//		}else{//尝试加载新的KLASS
//
//		}
//	}
//	return nil;
//}
//
///**
//	根据限定符加载Klass
//	加载klass的时机 https://baijiahao.baidu.com/s?id=1662931890502685014&wfr=spider&for=pc
// */
//func (me *Interpreter)loadKlassByQualifier(e IFObject) IKlass {
//	if s,ok:=e.(*FString);ok{
//		if s2:=vms.metaKlass.get( s.GetVal() ) ; s2==nil{
//			//尝试加载新的KLASS
//
//
//		}else{
//			return s2.(IKlass)
//		}
//	}
//	return nil;
//}
//
//
//
//func (me *Interpreter)BuildFrame(callable IFObject, argType *NSMethodArgType, args *FArray){
//	me.buildFrame(callable, argType , args)
//}
//
//
//func (me *Interpreter)buildFrame(callable IFObject, argType *NSMethodArgType, args *FArray){
//	me.stackDeep++;
//	if callable.GetKlass().getName() == KLS_T.NativeMethodKlass{
//		retOb := callable.(*FMethod).Call(args)
//		if argType.ret.GetSize() > 0{
//				me.frame.statck.push(retOb)
//		}
//	}else if callable.GetKlass().getName() == KLS_T.MethodKlass{
//		//虚方法
//		frame:=NewFrame(callable)
//		frame.sendFrame = me.frame;
//		me.frame = frame;
//	}
//
//}
//
//func (me *Interpreter)leaveFrame(){
//
//	me.frame = me.frame.sendFrame;
//	me.stackDeep--;
//}
//
//
//
//func (me *Interpreter)exec(){
//
//	for{
//		op := me.frame.HasMoreOpcode()
//
//		if op==nil{
//			goto exception
//		}
//
//
//			fmt.Printf("opcode n:%d [0x%x] %s \n", op.n,op.op,OPCODE_DESC[op.op])
//			switch op.op {
//
//			case OP.ALOAD_0:
//			case OP.ALOAD_1:
//				//me.frame.statck.push(me.frame.localVal.get(1).(IFObject))
//				me.frame.StackPush(me.frame.GetLocal(1).(IFObject))
//			case OP.ALOAD_2:
//				me.frame.statck.push(me.frame.localVal.get(2).(IFObject))
//			case OP.INVOKESPECIAL:
//			case OP.INVOKEVIRTUAL:
//				//The parameter type is displayed first
//				if opArg,ok:=op.oper.(*FArray); ok{
//					//argtype
//					argType := opArg.pop()
//					opArg.pop()
//					mthodName	:=	opArg.pop()
//					//qualifier
//					opArg.pop()
//					//Processing of eucalyptus
//					atype:=argType.(*NSMethodArgType);
//					args :=  NewFArray()
//					for i:=0;i<atype.arg.GetSize();i++{
//						args.add(me.frame.statck.pop())
//					}
//					//Instance objects
//					callOb	:= me.frame.statck.pop()
//					mt := callOb.getMethod(mthodName.(*FString).GetVal())
//					var retOb  = mt.(*FMethod).Call(args);
//					//  if has return values ,need push value to stack
//					if atype.ret.GetSize() > 0{
//						me.frame.statck.push(retOb)
//					}
//				}
//
//			case OP.RETURN:
//				fmt.Println("OP_RETURN")
//			case OP.ICONST_0://nada:压入int到栈 jvm:压入uint8到栈
//				me.frame.statck.push(NewFInt(0))
//			case OP.ICONST_1://nada:压入int到栈 jvm:压入uint8到栈
//				me.frame.statck.push(NewFInt(1))
//			case OP.LDC:
//				me.frame.statck.push(op.oper)
//			case OP.ASTORE_0 ,OP.ASTORE_3,OP.ISTORE_0,OP.ISTORE_1,
//				OP.ISTORE_2,OP.ISTORE_3: //don't need operands
//			case OP.FSTORE://为压缩空间 nada 本地变量编号全部为2字节
//
//			case OP.LDC_W:
//				me.frame.statck.push(op.oper)
//				fmt.Println(op.oper)
//
//			case OP.INVOKESTATIC:
//				/**
//					op.oper args
//					--- 限定符
//					--- 调用的方法
//					--- 描述符
//					--- 参数类型和个数
//				 */
//				//if opArg,ok:=op.oper.(*FArray); ok{
//				//	argType:=opArg.pop();
//				//	//desc
//				//	opArg.pop();
//				//	mn:=opArg.pop();
//				//	qualfier:=opArg.pop();
//				//	mt:=vms.metaKlass.get(qualfier.(*FString).GetVal()).(IKlass).getStaticMethod(mn.(*FString).GetVal()).(*FMethod)
//				//
//				//	//参数都在栈中，构造参数
//				//	atype:=argType.(*NSMethodArgType);
//				//	args :=  NewFArray()
//				//	for i:=0;i<atype.arg.GetSize();i++{
//				//		args.add(me.frame.statck.pop())
//				//	}
//				//
//				//	me.buildFrame(mt,atype,args)
//				//
//				//
//				//	//fmt.Println(op.oper,mn,qualfier,argType,desc)
//				//}
//				me.ins.opGetStatic.Execute(me.frame,op,me)
//
//
//			case OP.ASTORE:
//				ob:=me.frame.statck.pop()
//				me.frame.localVal.set(1,ob)
//			case OP.ASTORE_1:
//				ob:=me.frame.statck.pop()
//				me.frame.localVal.set(1,ob)
//				fmt.Println(me.frame.localVal)
//
//			case OP.ASTORE_2:
//				ob:=me.frame.statck.pop()
//				me.frame.localVal.set(2,ob)
//				fmt.Println(me.frame.localVal)
//
//			case OP.LSTORE:
//
//			case OP.ISTORE:
//
//			case OP.BIPUSH://nada:压入int到栈 jvm:压入uint8到栈
//
//			case OP.SIPUSH:
//
//			case OP.ILOAD:
//
//			case OP.IF_CMPGE:
//
//			case OP.NEW:
//
//			case OP.DUP:
//
//
//			case OP.INVOKEDYNAMIC:
//
//			case OP.POP,OP.POP2://iinc index(uint8) const(int8) to int
//
//			case OP.IINC:
//
//			case OP.GOTO:
//			case OP.GETSTATIC:
//				//类没有加载的话，可以在此处加载 所以 要验证
//				if opArg,ok:=op.oper.(*FArray); ok{
//					//实例类型
//					insType		:= 	opArg.pop()
//					//实例名称
//					ins			:= 	opArg.pop()
//					//这个实例所在的Klass
//					onwKlass	:=	opArg.pop()
//					ob 		:= vms.metaKlass.get(onwKlass.(*FString).GetVal()).(IKlass)
//					insOb 	:= ob.getStaticField(  ins.(*FString).GetVal()  )
//					me.frame.statck.push(insOb)
//					log.Print(me.frame.localVal,insType,onwKlass,ob,insOb)
//
//
//
//				}
//
//
//
//
//
//			default:
//
//
//			}
//
//
//
//	}
//
//	exception:{
//		me.leaveFrame();
//		//如果不是最底部的栈帧 需要继续执行
//		if me.stackDeep > 0{
//			me.exec();
//		}
//	}
//
//
//
//}