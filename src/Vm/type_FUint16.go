package Vm

type FUint16 struct {

	klass IKlass
	val uint16
}

func (me *FUint16) getField(k string) IFObject {
	panic("implement me")
}

func (me *FUint16) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *FUint16) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *FUint16) setMethod(k string, v IFObject) {
	panic("implement me")
}

func (me *FUint16) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FUint16) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FUint16) getVal() uint16 {
	return me.val
}

func (me *FUint16) setVal(v uint16) *FUint16 {
	me.val = v;
	return me
}

func (me *FUint16) GetKlass() IKlass {
	return me.klass
}


/**
	绑定klass
 */
func (me *FUint16) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me* FUint16) Build() *FUint16 {
	//设置class
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_UINT16).(*Uint16Klass)  )
	return me

}




func (me *FUint16) test() IFObject{
	return me
}






