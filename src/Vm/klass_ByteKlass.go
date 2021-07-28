package Vm

type ByteKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *ByteKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me *ByteKlass) GetDict(k string) IFObject {
	return me.dict[k]
}

/**
设置类型对象
*/
func (me *ByteKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
获取类型对象
*/
func (me *ByteKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *ByteKlass)GetName() string{
	return me.name
}

func (me *ByteKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *ByteKlass)Init() IKlass {
	me.name = "ByteKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}


func (me *ByteKlass) getStaticDict(k string) IFObject {
	return nil
}


