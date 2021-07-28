package Vm

type FFloat struct {
	//方法名
	mName string
	//方法字节码
	code  []byte
	//已经处理好的立即数
	codeList map[uint8]IFObject
	klass IKlass
	//原生方法保存
	call func(args []IFObject) IFObject

	val float64

}


func (me *FFloat) getVal() float64 {
	return me.val
}

func (me *FFloat) setVal(v float64) *FFloat {
	me.val = v;
	return me
}

func (me *FFloat) GetKlass() IKlass {
	return me.klass
}

func (me *FFloat) ValToString(){

}




/**
	绑定klass
 */
func (me *FFloat) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me* FFloat) Build() *FFloat {
	//设置class
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_FLOAT).(*FloatKlass)  )
	return me

}
