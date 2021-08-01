package Vm

type FObject struct {
	val   string
	len   int
	klass IKlass

	fieldDict *FMap
	methodDict *FMap
}

func (me * FObject) getFieldDict() IFObject {
	return me.fieldDict
}

func (me * FObject) getMethodDict() IFObject {
	return me.methodDict
}

func (me * FObject) getField(k string) IFObject {
	return me.fieldDict.get(k).(IFObject)
}

func (me * FObject) getMethod(k string) IFObject {
	panic("implement me")
}

func (me * FObject) setField(k string, v IFObject) {
	me.GetKlass().setObField(me,NewFString(k),v)
}

func (me * FObject) setMethod(k string, v IFObject) {
	me.GetKlass().setObMethod(me,NewFString(k),v)
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
	me.methodDict = NewFMap()
	me.fieldDict = NewFMap()
	return me

}


func (me *FObject) test() IFObject{
	return me
}






