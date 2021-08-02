package Vm

type DoubleKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject


	//静态方法
	staticMethod map[string]IFObject
	//动态方法
	insMethod map[string]IFObject

	//属性
	staticField map[string]IFObject

	//动态属性
	insField map[string]IFObject

}

func (me *DoubleKlass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *DoubleKlass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *DoubleKlass) setStaticMethod(name string, method IFObject) IKlass {
	me.staticMethod[name]=method;
	return me
}

func (me *DoubleKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *DoubleKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *DoubleKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *DoubleKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *DoubleKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *DoubleKlass) getStaticMethod(name string) IFObject {
	return me.staticMethod[name]
}

func (me *DoubleKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *DoubleKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *DoubleKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
	设置类型对象
*/
func (me *DoubleKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
	获取类型对象
 */
func (me *DoubleKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *DoubleKlass)GetName() string{
	return me.name
}

func (me *DoubleKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *DoubleKlass)Init() IKlass {
	me.name = "DoubleKlass"
	me.insMethod = make(map[string]IFObject)
	me.staticMethod = make(map[string]IFObject)
	me.insField = make(map[string]IFObject)
	me.staticField = make(map[string]IFObject)
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}




