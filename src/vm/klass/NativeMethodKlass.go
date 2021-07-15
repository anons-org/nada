package klass

import (
	"vm/type"
)

type NativeMethodKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *ftype.TypeObject
}

func (me NativeMethodKlass) GetDict(k string) ftype.IFObject {
	panic("implement me")
}

func (me NativeMethodKlass) GetKlassMethod(k string) ftype.IFObject {
	return nil
}

func (me NativeMethodKlass) Dict(k string) IKlass {
	panic("implement me")
}

/**
	设置类型对象
 */
func (me NativeMethodKlass)SetTypeObject(to *ftype.TypeObject) {
	me.typeObject = to
}

func (me NativeMethodKlass)GetTypeObject() *ftype.TypeObject {
	return me.typeObject
}

func (me NativeMethodKlass)GetName() string{
	return me.name
}

func (me NativeMethodKlass)Alloc() *ftype.IFObject {
	return nil
}


/**
	初始化Klass
 */
func (me NativeMethodKlass)Init() IKlass {
	me.name = "NativeMethodKlass"
	typeObj := new(ftype.TypeObject)
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}

