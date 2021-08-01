package Vm



type FStack struct {
	klass IKlass
	data *FArray
	//测试 观察用
	debugData *FArray

}

func (me *FStack) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FStack) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FStack) getField(k string) IFObject {
	panic("implement me")
}

func (me *FStack) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *FStack) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *FStack) setMethod(k string, v IFObject) {
	panic("implement me")
}

func (me *FStack) GetKlass() IKlass {
	return me.klass
}


func (me *FStack) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me *FStack) Build() *FStack{
	me.debugData = NewFArray();
	me.data = NewFArray()
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_STACK).(*StackKlass)  )
	return me
}


func (me *FStack) pop() IFObject{
	me.debugData.pop()
	return me.data.pop().(IFObject);
}

func (me *FStack) size() int{
	return me.data.size;
}

func (me *FStack) push(o IFObject) {
	me.debugData.add(o)
	me.data.add(o);
}





func (me *FStack) test() IFObject{
	return me
}






