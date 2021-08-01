package Vm

type ArrayKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *ArrayKlass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *ArrayKlass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *ArrayKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *ArrayKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *ArrayKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *ArrayKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *ArrayKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *ArrayKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *ArrayKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *ArrayKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *ArrayKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *ArrayKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
设置类型对象
*/
func (me *ArrayKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *ArrayKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *ArrayKlass)GetName() string{
	return me.name
}

func (me *ArrayKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *ArrayKlass)Init() IKlass {
	me.name = "ArrayKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}




