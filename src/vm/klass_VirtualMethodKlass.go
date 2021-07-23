package Vm

type VirtualMethodKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
}

func (me *VirtualMethodKlass) GetDict(k string) IFObject {
	panic("implement me")
}


func (me *VirtualMethodKlass) Dict(k string) IKlass {
	panic("implement me")
}

/**
	设置类型对象
 */
func (me *VirtualMethodKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

func (me *VirtualMethodKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *VirtualMethodKlass)GetName() string{
	return me.name
}

func (me *VirtualMethodKlass)Alloc() *IFObject {
	return nil
}


/**
	初始化Klass
 */
func (me *VirtualMethodKlass)Init() IKlass {
	me.name = "VirtualMethodKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}

