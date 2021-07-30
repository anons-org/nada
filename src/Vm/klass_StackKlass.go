package Vm

type StackKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *StackKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me *StackKlass) GetDict(k string) IFObject {
	return me.dict[k]
}



/**
设置类型对象
*/
func (me *StackKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *StackKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *StackKlass)GetName() string{
	return me.name
}

func (me *StackKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *StackKlass)Init() IKlass {
	me.name = "StackKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}


func (me *StackKlass) getStaticDict(k string) IFObject {
	return nil
}


