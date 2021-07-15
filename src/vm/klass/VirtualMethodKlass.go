package klass

import (
	"vm/type"
)

type VirtualMethodKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *ftype.TypeObject
}

func (me *VirtualMethodKlass) GetDict(k string) ftype.IFObject {
	panic("implement me")
}

func (me *VirtualMethodKlass) GetKlassMethod(k string) ftype.IFObject {
	return nil
}

func (me *VirtualMethodKlass) Dict(k string) IKlass {
	panic("implement me")
}

/**
	设置类型对象
 */
func (me *VirtualMethodKlass)SetTypeObject(to *ftype.TypeObject) {
	me.typeObject = to
}

func (me *VirtualMethodKlass)GetTypeObject() *ftype.TypeObject {
	return me.typeObject
}

func (me *VirtualMethodKlass)GetName() string{
	return me.name
}

func (me *VirtualMethodKlass)Alloc() *ftype.IFObject {
	return nil
}


/**
	初始化Klass
 */
func (me *VirtualMethodKlass)Init() IKlass {
	me.name = "VirtualMethodKlass"
	typeObj := new(ftype.TypeObject)
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}

