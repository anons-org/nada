package Vm


type FMap struct {

	klass IKlass
	data *Map
}

func (me *FMap) GetKlass() IKlass {
	return me.klass
}

func (me *FMap) ValToString(){

}

func (me *FMap) SetKlass(klass IKlass) {
	me.klass = klass;
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











