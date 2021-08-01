package Vm

type StackKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *StackKlass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *StackKlass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *StackKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *StackKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *StackKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *StackKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *StackKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *StackKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *StackKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *StackKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *StackKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *StackKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
设置类型对象
*/
func (me *StackKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *StackKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *StackKlass)GetName() string{
	return me.name
}

func (me *StackKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *StackKlass)Init() IKlass {
	me.name = "StackKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}



