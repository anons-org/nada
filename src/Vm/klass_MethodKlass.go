package Vm

type MethodKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject

	//静态属性和静态方法
	staticDict map[string]IFObject
	//动态属性和动态方法
	insDict map[string]IFObject


}

func (me *MethodKlass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *MethodKlass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *MethodKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *MethodKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *MethodKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *MethodKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *MethodKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *MethodKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *MethodKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *MethodKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *MethodKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *MethodKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
	设置类型对象
 */
func (me *MethodKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

func (me *MethodKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *MethodKlass)GetName() string{
	return me.name
}

func (me *MethodKlass)Alloc() *IFObject {
	return nil
}


/**
	初始化Klass
 */
func (me *MethodKlass)Init() IKlass {
	me.name = "MethodKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}

