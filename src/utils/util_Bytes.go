package Utils

type Bytes struct {
	data []byte
	idx int
}

func (me *Bytes) read(n int)  []byte{
	idx := me.idx;
	me.idx	+=n;
	return me.data[idx:idx+n];
}


func (me *Bytes) Read(n int)  []byte{
	return me.read(n);
}

func (me *Bytes) GetIdx()  int{
	return me.idx;
}




func NewBytes(d []byte) *Bytes {
	b:=new(Bytes)
	b.data  = d
	return b
}