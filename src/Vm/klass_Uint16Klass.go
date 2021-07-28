package Vm

type Uint16Klass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject

	//方法集合
	dict map[string]IFObject

}



func (me *Uint16Klass) Dict(k string) IKlass {
	panic("implement me")
}


func (me *Uint16Klass) GetDict(k string) IFObject {
	return me.dict[k]
}



/**
设置类型对象
*/
func (me *Uint16Klass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *Uint16Klass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *Uint16Klass)GetName() string{
	return me.name
}

func (me *Uint16Klass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *Uint16Klass)Init() IKlass {
	me.name = "Uint16Klass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}


func (me *Uint16Klass) getStaticDict(k string) IFObject {
	return nil
}


