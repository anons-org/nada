package VmObject


type String struct {
	val string
	len int
}

func (me *String) GetLen() int{
	return me.len;
}

func (me *String) GetVal() string{
	return me.val;
}

func (me String) SetVal(val string) String{
	me.val = val;
	return me;
}