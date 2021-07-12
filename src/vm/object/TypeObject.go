package VmObject

import VmObjectKlass "vm/object/klass"
/**
	类型对象
 */
type TypeObject struct {
	//所属klass
	onwKlass VmObjectKlass.IKlass
}


func  (me *TypeObject)SetOnwKlass(onwKlass VmObjectKlass.IKlass) {
	me.onwKlass = onwKlass
}

func  (me *TypeObject)GetOnwKlass() VmObjectKlass.IKlass {
	return me.onwKlass
}


