package vm

type OpGetStatic struct {
	
}

func (me *OpGetStatic) Execute(frame *Frame, op *OpcodeStruct,interp *Interpreter)  {


	/**
		op.oper args
		--- 限定符
		--- 调用的方法
		--- 描述符
		--- 参数类型和个数
	*/

	if opArg,ok:=op.GetOper().(*FArray); ok{
		argType		:=	opArg.Pop();
		//desc
		opArg.Pop();
		mn			:=	opArg.Pop();
		qualfier	:=	opArg.Pop();
		mt			:=	GetMetaKlass( qualfier.(*FString).GetVal() ).(IKlass).GetStaticMethod(mn.(*FString).GetVal()).(*FMethod)
		//参数都在栈中，构造参数
		atype:=argType.(*NSMethodArgType);
		args :=  NewFArray()
		for i:=0;i<atype.GetArg().GetSize();i++{
			args.Add( frame.StackPop() )
		}
		interp.BuildFrame(mt, atype, args)

	}
}
