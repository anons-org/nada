package Vm

type FArray struct {
	klass IKlass
	val []IFObject
}


func (me *FArray) getVal() []IFObject{
	return me.val
}

func (me *FArray) setVal(v []IFObject) *FArray {
	me.val = v;
	return me
}

func (me *FArray) GetKlass() IKlass {
	return me.klass
}





/**
绑定klass
*/
func (me *FArray) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me* FArray) Build() *FArray {
	//设置class
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_ARRAY).(*ArrayKlass)  )
	return me

}

//添加数组元素
func (me *FArray) add(v IFObject) *FArray {
	me.val = append(me.val, v);
	return me
}

func (me *FArray) get(idx IFObject) IFObject {
	fint:=idx.(*FInt);
	return me.val[fint.val]
}










