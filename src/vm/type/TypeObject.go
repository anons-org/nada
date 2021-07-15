package ftype

import "vm/klass"

/**
	类型对象
 */
type TypeObject struct {
	//所属klass
	onwKlass klass.IKlass
}


func  (me *TypeObject)SetOnwKlass(onwKlass klass.IKlass) {
	me.onwKlass = onwKlass
}

func  (me *TypeObject)GetOnwKlass() klass.IKlass {
	return me.onwKlass
}


