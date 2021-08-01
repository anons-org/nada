package Vm

type Uint16Klass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject

	//方法集合
	dict map[string]IFObject

}

func (me *Uint16Klass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *Uint16Klass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *Uint16Klass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *Uint16Klass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *Uint16Klass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *Uint16Klass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *Uint16Klass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *Uint16Klass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *Uint16Klass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *Uint16Klass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
设置类型对象
*/
func (me *Uint16Klass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *Uint16Klass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *Uint16Klass)GetName() string{
	return me.name
}

func (me *Uint16Klass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *Uint16Klass)Init() IKlass {
	me.name = "Uint16Klass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}




