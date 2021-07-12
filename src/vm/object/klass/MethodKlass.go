package VmObjectKlass

import VmObject "vm/object"



type MethodKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *VmObject.TypeObject
}


/**
	设置类型对象
 */
func (me *MethodKlass)SetTypeObject(to *VmObject.TypeObject) {
	me.typeObject = to

}

func (me *MethodKlass)GetTypeObject() *VmObject.TypeObject{
	return me.typeObject
}

func (me *MethodKlass)GetName() string{
	return me.name
}

func (me *MethodKlass)Alloc() *VmObject.IFObject{
	return nil
}


/**
	初始化Klass
 */
func (me *MethodKlass)Init() *MethodKlass{
	me.name = "MethodKlass"
	typeObj := new(VmObject.TypeObject)
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}

