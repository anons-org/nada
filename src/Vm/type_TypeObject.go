package Vm



/**
	类型对象
 */
type TypeObject struct {
	//所属klass
	onwKlass IKlass
	klass IKlass
}

func (me *TypeObject) getFieldDict() IFObject {
	panic("implement me")
}

func (me *TypeObject) getMethodDict() IFObject {
	panic("implement me")
}

func  (me *TypeObject)SetOnwKlass(onwKlass IKlass) {
	me.onwKlass = onwKlass
}

func  (me *TypeObject)GetOnwKlass() IKlass {
	return me.onwKlass
}

//IFObject----------------------------------------
func (me *TypeObject) GetKlass() IKlass{
	return me.klass
}


func (me *TypeObject) SetKlass(klass IKlass) {
	me.klass = klass
}


func (me *TypeObject) Build() *TypeObject{
	me.SetKlass( vms.metaKlass.get(BUILTN.NADA_TYPE).(*TypeKlass) )
	return me
}

