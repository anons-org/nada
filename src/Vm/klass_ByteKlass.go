package Vm

type ByteKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *ByteKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *ByteKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *ByteKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *ByteKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *ByteKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *ByteKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *ByteKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *ByteKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *ByteKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *ByteKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
设置类型对象
*/
func (me *ByteKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *ByteKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *ByteKlass)GetName() string{
	return me.name
}

func (me *ByteKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *ByteKlass)Init() IKlass {
	me.name = "ByteKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}




