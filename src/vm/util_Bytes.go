package Vm

type Bytes struct {
	data []byte
	idx int
}

func (me *Bytes ) read(n int)  []byte{
	idx := me.idx;
	me.idx	+=n;
	return me.data[idx:idx+n];
}

func NewBytes(d []byte) *Bytes {
	b:=new(Bytes)
	b.data  = d
	return b
}