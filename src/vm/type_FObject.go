package vm

type FObject struct {
	name  string
	klass IKlass

	fieldDict *FMap
	methodDict *FMap
}

func (me *FObject) getFieldDict() IFObject {
	return me.fieldDict
}

func (me *FObject) getMethodDict() IFObject {
	return me.methodDict
}

func (me *FObject) getField(k string) IFObject {
	return me.fieldDict.get(k).(IFObject)
}

func (me *FObject) getMethod(k string) IFObject {
	return me.GetKlass().getObMethod(me, NewFString(k)).(IFObject)

}

func (me *FObject) setField(k string, v IFObject) {
	me.GetKlass().setObField(me, NewFString(k),v)
}

func (me *FObject) setMethod(k string, v IFObject) {
	me.GetKlass().setObMethod(me, NewFString(k),v)
}

func (me*FObject) GetKlass() IKlass {
	return me.klass
}




/**
绑定klass
*/
func (me*FObject) SetKlass(klass IKlass) {
	me.klass = klass
}


func (me*FObject) Build() *FObject {
	//不需要设置 本质就是个空壳对象，外部给予CLASS 另外 ObjectKlass 和FObject没有直接关系
	me.methodDict = NewFMap()
	me.fieldDict = NewFMap()
	return me

}


func (me *FObject) test() IFObject {
	return me
}






