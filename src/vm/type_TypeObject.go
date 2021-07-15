package Vm



/**
	类型对象
 */
type TypeObject struct {
	//所属klass
	onwKlass IKlass
}


func  (me *TypeObject)SetOnwKlass(onwKlass IKlass) {
	me.onwKlass = onwKlass
}

func  (me *TypeObject)GetOnwKlass() IKlass {
	return me.onwKlass
}


