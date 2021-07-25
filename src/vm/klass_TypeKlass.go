package Vm

type TypeKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject

}


func (me *TypeKlass)GetName() string{
	return me.name
}


func (me *TypeKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me *TypeKlass) GetDict(k string) IFObject {
	return nil
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


func (me *TypeKlass) getStaticDict(k string) IFObject {
	return nil
}
