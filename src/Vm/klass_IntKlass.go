package Vm

type IntKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *IntKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me *IntKlass) GetDict(k string) IFObject {
	return me.dict[k]
}



/**
设置类型对象
*/
func (me *IntKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *IntKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *IntKlass)GetName() string{
	return me.name
}

func (me *IntKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *IntKlass)Init() IKlass {
	me.name = "IntKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}


func (me *IntKlass) getStaticDict(k string) IFObject {
	return nil
}


