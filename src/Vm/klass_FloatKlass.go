package Vm

type FloatKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *FloatKlass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *FloatKlass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *FloatKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *FloatKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *FloatKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *FloatKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *FloatKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *FloatKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *FloatKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *FloatKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *FloatKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *FloatKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
	设置类型对象
*/
func (me *FloatKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
	获取类型对象
 */
func (me *FloatKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *FloatKlass)GetName() string{
	return me.name
}

func (me *FloatKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *FloatKlass)Init() IKlass {
	me.name = "FloatKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}




