package Vm

import "container/list"

type FStack struct {
	klass IKlass
	data list.List
	//测试 观察用
	debugData *FArray
	_size int;
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











