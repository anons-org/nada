package VmObject


type FString struct {
	val string
	len int
}

func (me *FString) GetLen() int{
	return me.len;
}

func (me *FString) GetVal() string{
	return me.val;
}

func (me FString) SetVal(val string) FString{
	me.val = val;
	return me;
}