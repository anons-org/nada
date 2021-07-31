package Vm

type ObjectKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *ObjectKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me *ObjectKlass) GetDict(k string) IFObject {
	return me.dict[k]
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


func (me *ObjectKlass) getStaticDict(k string) IFObject {
	return nil
}


