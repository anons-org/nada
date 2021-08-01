package Vm

type TypeKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject

}

func (me *TypeKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *TypeKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *TypeKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *TypeKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *TypeKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *TypeKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *TypeKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *TypeKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *TypeKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *TypeKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *TypeKlass)GetName() string{
	return me.name
}







/**
设置类型对象
*/
//IKlass-------------------------------------------------
func (me *TypeKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/

func (me *TypeKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *TypeKlass)Alloc() *IFObject {
	return nil
}



/**
初始化Klass
*/
func (me *TypeKlass)Init() IKlass {
	me.name = "TypeKlass"
	//此处的TypeObject不能Build，否则错误
	typeObj := new(TypeObject)
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}


