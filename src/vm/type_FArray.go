package vm

type FArray struct {
	klass IKlass
	val   []IFObject
	//最后一个元素的位置
	size int
}

func (me *FArray) getFieldDict() IFObject {
	panic("implement me")
}

func (me *FArray) getMethodDict() IFObject {
	panic("implement me")
}

func (me *FArray) getField(k string) IFObject {
	panic("implement me")
}

func (me *FArray) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *FArray) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *FArray) setMethod(k string, v IFObject) {
	panic("implement me")
}

func (me *FArray) getVal() []IFObject {
	return me.val
}

func (me *FArray) setVal(v []IFObject) *FArray {
	me.val = v;
	return me;
}

func (me *FArray) GetKlass() IKlass {
	return me.klass;
}


/**
	绑定klass
*/
func (me *FArray) SetKlass(klass IKlass) {
	me.klass = klass;
}


func (me*FArray) Build() *FArray {
	//设置class
	me.SetKlass(  vms.metaKlass.get(BUILTN.NADA_ARRAY).(*ArrayKlass)  );
	return me;

}

//添加数组元素
func (me *FArray) add(v IFObject) *FArray {
	me.val = append(me.val, v);
	me.size++;
	return me;
}


func (me *FArray) Add(v IFObject) *FArray {
	me.add(v);
	return me;
}


//虚拟机内部使用的GET
func (me *FArray) get(idx int) IFObject {
	if idx>me.size-1{
		return NONE;
	}
	return me.val[idx]
}



func  (me *FArray) remove( start, end int)  {
	me.val = append(me.val[:start], me.val[end:]...)
}



//需要删除
func (me *FArray) pop() IFObject {
	me.size--;
	if( me.size<0 ){
		me.size=0;
		return NONE;
	}
	ob:= me.val[me.size]
	me.val = me.val[:me.size]

	return ob
}


func (me *FArray) Pop() IFObject {
	me.size--;
	if( me.size<0 ){
		me.size=0;
		return NONE;
	}
	ob:= me.val[me.size]
	me.val = me.val[:me.size]

	return ob
}





func (me *FArray) test() IFObject {
	return me
}



