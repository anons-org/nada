package Vm

type ArrayKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *ArrayKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me *ArrayKlass) GetDict(k string) IFObject {
	return me.dict[k]
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


func (me *ArrayKlass) getStaticDict(k string) IFObject {
	return nil
}


