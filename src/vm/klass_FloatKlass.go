package Vm

type FloatKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
	//方法集合
	dict map[string]IFObject

}

func (me *FloatKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me *FloatKlass) GetDict(k string) IFObject {
	return me.dict[k]
}



/**
	设置类型对象
*/
func (me *FloatKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
	获取类型对象
 */
func (me *FloatKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *FloatKlass)GetName() string{
	return me.name
}

func (me *FloatKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *FloatKlass)Init() IKlass {
	me.name = "FloatKlass"
	typeObj := new(TypeObject).Build()
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )


	return me
}


func (me *FloatKlass) getStaticDict(k string) IFObject {
	return nil
}


