package Vm

type FFunction struct {
	//函数名
	name string
	//函数字节码
	byteCode string


}

func (me *FFunction) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FFunction) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FFunction) GetKlass() IKlass {
	panic("implement me")
}

func (me *FFunction) SetKlass(klass IKlass) {
	panic("implement me")
}

func (me *FFunction) getField(k string) IFObject {
	panic("implement me")
}

func (me *FFunction) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *FFunction) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *FFunction) setMethod(k string, v IFObject) {
	panic("implement me")
}

func (me *FFunction) test() IFObject{
	return me
}


