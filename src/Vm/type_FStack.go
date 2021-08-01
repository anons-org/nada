package Vm

import "container/list"

type FStack struct {
	klass IKlass
	data list.List
	//测试 观察用
	debugData *FArray
	_size int;
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
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_STACK).(*StackKlass)  )
	return me
}


func (me *FStack) pop() IFObject{
	me.debugData.pop()
	return me.data.Back().Value.(IFObject);
}

func (me *FStack) size() int{
	return me._size;
}

func (me *FStack) push(o IFObject) {
	me.debugData.add(o)
	me.data.PushBack(o);
}





func (me *FStack) test() IFObject{
	return me
}






