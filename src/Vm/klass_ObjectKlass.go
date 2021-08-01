package Vm

type ObjectKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *ObjectKlass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *ObjectKlass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *ObjectKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *ObjectKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *ObjectKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *ObjectKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *ObjectKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *ObjectKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *ObjectKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *ObjectKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *ObjectKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *ObjectKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
设置类型对象
*/
func (me *ObjectKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *ObjectKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *ObjectKlass)GetName() string{
	return me.name
}

func (me *ObjectKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *ObjectKlass)Init() IKlass {
	me.name = "ObjectKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}




