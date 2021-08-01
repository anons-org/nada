package Vm

type NativeMethodKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
}

func (me *NativeMethodKlass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *NativeMethodKlass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *NativeMethodKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *NativeMethodKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *NativeMethodKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *NativeMethodKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *NativeMethodKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *NativeMethodKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *NativeMethodKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *NativeMethodKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *NativeMethodKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *NativeMethodKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
	设置类型对象
 */
func (me *NativeMethodKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

func (me *NativeMethodKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *NativeMethodKlass)GetName() string{
	return me.name
}

func (me *NativeMethodKlass)Alloc() *IFObject {
	return nil
}


/**
	初始化Klass
 */
func (me *NativeMethodKlass)Init() IKlass {
	me.name = "NativeMethodKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}


