package Vm

type IntKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *IntKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *IntKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *IntKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *IntKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *IntKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *IntKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *IntKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *IntKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *IntKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *IntKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
设置类型对象
*/
func (me *IntKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *IntKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *IntKlass)GetName() string{
	return me.name
}

func (me *IntKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *IntKlass)Init() IKlass {
	me.name = "IntKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}



