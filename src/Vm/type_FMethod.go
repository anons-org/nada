package Vm

type FMethod struct {
	//方法名
	mName string
	//方法字节码
	code  []byte
	//已经处理好的立即数
	codeList map[uint8]IFObject
	klass IKlass
	//原生方法保存
	call func(args []IFObject) IFObject

	//参数和返回值的类型
	argsType string
	//限定符 说白了就是类的路径 java/lang/xxoo
	spec map[string]string

	//方法访问权限
	//静态还是动态方法

	//方法参数个数
	//方法返回值
	//方法返回值个数
	//是否是匿名方法

}

func (me *FMethod) GetKlass() IKlass {
	return me.klass
}


/**
	设置方法名
 */
func (me* FMethod) SetMethod(name string) {

}


func (me* FMethod) setSpec(s string) {
	me.spec["java"] = s;
	me.spec["js"] = s;
	me.spec["php"] = s;
}

/**
	绑定klass
 */
func (me *FMethod) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me *FMethod) Call(caller string, args []IFObject) IFObject {
	return me.call(args)
}

/**
	设置原生方法的call
 */
func (me *FMethod) SetCall(call func(args []IFObject) IFObject) *FMethod {
	me.call = call
	return me
}




/**
	初始化方法
 */
func (me* FMethod) Build(t MethodType) *FMethod {
	//设置class
	if t  ==  METHOD_TYPE_VIRTUAL{
		//虚拟方法
		me.SetKlass(   vms.metaKlass.get(BUILTN.NADA_VIRTUAL).(*MethodKlass) )
	}else{
		//原生方法
		me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_NATIVE).(*NativeMethodKlass) )
	}
	me.codeList = make(map[uint8]IFObject)
	me.spec = make(map[string]string)
	me.spec["java"] = "";
	return me
}




func (me *FMethod) setCode(code []byte)  *FMethod {
	me.code = code
	return me
}









