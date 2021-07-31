package Vm

type FObject struct {
	val   string
	len   int
	klass IKlass
}



func (me* FObject) GetKlass() IKlass {
	return me.klass
}


func (me *FObject) GetLen() int{
	return me.len;
}

func (me *FObject) GetVal() string{
	return me.val;
}

func (me *FObject) SetVal(val string) *FObject {
	me.val = val;
	return me;
}


/**
绑定klass
*/
func (me* FObject) SetKlass(klass IKlass) {
	me.klass = klass
}


func (me* FObject) Build() *FObject {
	//不需要设置 本质就是个空壳对象，外部给予CLASS 另外 ObjectKlass 和FObject没有直接关系
	return me

}




