package Vm


type Klass struct {

	//klass名称
	name *FString
	//类型对象
	typeObject *TypeObject

	//作废
	dict map[string]IFObject

	//静态属性和静态方法
	staticDict map[string]IFObject
	//属性
	staticField map[string]IFObject

	//动态属性和动态方法
	insDict map[string]IFObject
	insField  map[string]IFObject
	//类的限定符 x.x.x.Test
	qualifer string

}


func (me Klass) Dict(k string) IKlass {
	panic("implement me")
}


func (me Klass) GetDict(k string) IFObject {
	return me.dict[k]
}


func (me *Klass) getStaticDict(k string) IFObject {
	return me.staticDict[k]
}

func (me *Klass) getInsDict(k string) IFObject {
	return me.staticDict[k]
}




/**
	设置类型对象
*/
func (me Klass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
	获取类型对象
 */
func (me Klass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *Klass)GetName() *FString{
	return me.name
}

func (me *Klass)Alloc() *IFObject {
	return nil
}



/**
初始化Klass
*/
func (me *Klass)Init() IKlass {
	return me
}






//添加静态虚拟方法
//func  (me *Klass)addStaticVirtualMethod(name string, code []byte) *Klass{
//	m:=NewFMethod(2,name)
//	me.staticDict[name] = m
//	return me
//}

func  (me *Klass)addStaticMethod(name string, method IFObject) *Klass{
	me.staticDict[name] = method
	return me
}

//添加静态属性
func  (me *Klass)addStaticField(name string, field IFObject) *Klass{
	me.staticField[name] = field
	return me
}


//添加动态虚拟方法
func  (me *Klass)addInsMethod(name string, method IFObject) *Klass{
	me.insDict[name] = method
	return me
}


func  (me *Klass)addInsField(name string, field IFObject) *Klass{
	me.insField[name] = field
	return me
}


//添加原生方法
func  (me *Klass)addNativeMethod(name string){

}


