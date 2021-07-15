package ftype

import (
	Universe "vm/universe"
)

type FString struct {
	val   string
	len   int
	klass IKlass
}

func (me FString) ValToString() {

}

func (me*FString) GetKlass() IKlass {
	return me.klass
}


func (me *FString) GetLen() int{
	return me.len;
}

func (me *FString) GetVal() string{
	return me.val;
}

func (me *FString) SetVal(val string) *FString {
	me.val = val;
	return me;
}


/**
绑定klass
*/
func (me* FString) SetKlass(klass IKlass) {
	me.klass = klass
}


func (me* FString) Build() *FString {
	//设置class
	me.SetKlass(  Universe.KlassBean["nada.lib.StringKlass"]  )
	return me
}


