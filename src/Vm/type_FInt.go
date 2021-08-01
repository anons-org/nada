package Vm

type FInt struct {
	klass IKlass
	val int
}

func (me *FInt) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FInt) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FInt) getVal() int {
	return me.val
}

func (me *FInt) setVal(v int) *FInt {
	me.val = v;
	return me
}

func (me *FInt) GetKlass() IKlass {
	return me.klass
}






/**
绑定klass
*/
func (me *FInt) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me* FInt) Build() *FInt {
	//设置class
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_INT).(*IntKlass)  )
	return me

}










