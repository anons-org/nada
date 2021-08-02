package Vm

type FDouble struct {
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

func (me *FDouble) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FDouble) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FDouble) getField(k string) IFObject {
	panic("implement me")
}

func (me *FDouble) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *FDouble) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *FDouble) setMethod(k string, v IFObject) {
	panic("implement me")
}

func (me *FDouble) getVal() float64 {
	return me.val
}

func (me *FDouble) setVal(v float64) *FDouble {
	me.val = v;
	return me
}

func (me *FDouble) GetKlass() IKlass {
	return me.klass
}





/**
	绑定klass
 */
func (me *FDouble) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me* FDouble) Build() *FDouble {
	//设置class
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_DOUBLE).(*DoubleKlass)  )
	return me

}


func (me *FDouble) test() IFObject{
	return me
}
