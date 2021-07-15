package Vm

type NativeMethodKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject
}

func (me NativeMethodKlass) GetDict(k string) IFObject {
	panic("implement me")
}

func (me NativeMethodKlass) GetKlassMethod(k string) IFObject {
	return nil
}

func (me NativeMethodKlass) Dict(k string) IKlass {
	panic("implement me")
}

/**
	设置类型对象
 */
func (me NativeMethodKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

func (me NativeMethodKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me NativeMethodKlass)GetName() string{
	return me.name
}

func (me NativeMethodKlass)Alloc() *IFObject {
	return nil
}


/**
	初始化Klass
 */
func (me NativeMethodKlass)Init() IKlass {
	me.name = "NativeMethodKlass"
	typeObj := new(TypeObject)
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	return me
}

