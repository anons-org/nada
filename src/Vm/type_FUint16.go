package Vm

type FUint16 struct {

	klass IKlass
	val uint16
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









