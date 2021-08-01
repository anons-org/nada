package Vm

type FObject struct {
	val   string
	len   int
	klass IKlass

	fieldDict *FMap
	methodDict *FMap
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

//设置属性
func (me* FObject) setField(k IFObject,v IFObject) *FObject {

	return me

}

//增加方法
func (me* FObject) setMethod(k IFObject,v IFObject) *FObject {
	return me
}


func (me* FObject) getFieldDict() IFObject {
	return me.fieldDict
}

func (me* FObject) getMethodDict() IFObject {
	return me.methodDict
}






