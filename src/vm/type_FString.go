package vm

type FString struct {
	val   string
	len   int
	klass IKlass
}

func (me *FString) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FString) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FString) getField(k string) IFObject {
	panic("implement me")
}

func (me *FString) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *FString) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *FString) setMethod(k string, v IFObject) {
	panic("implement me")
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
func (me*FString) SetKlass(klass IKlass) {
	me.klass = klass
}


func (me*FString) Build() *FString {
	//设置class
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_STRING).(*StringKlass)  )
	return me

}

func (me *FString) test1() IFObject {
	return me
}



