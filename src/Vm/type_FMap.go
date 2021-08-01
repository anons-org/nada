package Vm




type FMap struct {
	klass IKlass
	data *Map
}

func (me *FMap) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FMap) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FMap) getField(k string) IFObject {
	panic("implement me")
}

func (me *FMap) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *FMap) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *FMap) setMethod(k string, v IFObject) {
	panic("implement me")
}

func (me *FMap) Build() *FMap{
	me.data = NewMap()
	return me
}

//VM内部使用
func (me *FMap) get(key interface{}) interface{}{

	return me.data.Get(key)
}

//vm内部使用
func (me *FMap) set(key interface{}, val interface{}) {
	me.data.Set(key,val)
}




func (me *FMap) GetKlass() IKlass {
	return me.klass
}



func (me *FMap) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me *FMap) test() IFObject{
	return me
}




