package vm

type FByte struct {

	klass IKlass
	//val 长度
	size int
	val []byte

}

func (me *FByte) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FByte) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FByte) getField(k string) IFObject {
	panic("implement me")
}

func (me *FByte) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *FByte) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *FByte) setMethod(k string, v IFObject) {
	panic("implement me")
}

func (me *FByte) getVal() []byte {
	return me.val
}

func (me *FByte) setVal(v []byte) *FByte {
	me.val = v;
	return me
}

func (me *FByte) getSize() int {
	return len(me.val)
}


func (me *FByte) GetKlass() IKlass {
	return me.klass
}

/**
	绑定klass
 */
func (me *FByte) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me*FByte) Build() *FByte {
	//设置class
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_BYTE).(*ByteKlass)  )
	return me
}
func (me *FByte) test() IFObject {
	return me
}
