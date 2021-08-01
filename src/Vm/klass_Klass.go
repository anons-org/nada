package Vm


type Klass struct {

	//klass名称
	name *FString
	//类型对象
	typeObject *TypeObject



	//静态方法
	staticMethod map[string]IFObject
	//动态方法
	insMethod map[string]IFObject

	//属性
	staticField map[string]IFObject

	//动态属性
	insField map[string]IFObject


	//
	//类的限定符 x.x.x.Test
	qualifer string

}

func (me *Klass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *Klass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *Klass) getStaticMethod(k string) IFObject {
	return me.staticMethod[k]
}

func (me *Klass) getStaticField(k string) IFObject {
	return me.staticField[k]
}

func (me *Klass) getInsMethod(k string) IFObject {
	return me.staticMethod[k]
}

func (me *Klass) getInsField(k string) IFObject {
	return me.staticField[k]
}





/**
	设置类型对象
*/
func (me *Klass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
	获取类型对象
 */
func (me *Klass)GetTypeObject() *TypeObject {
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



func  (me *Klass)setStaticMethod(name string, method IFObject) IKlass{
	me.staticMethod[name] = method
	return me
}

//添加静态属性
func  (me *Klass)setStaticField(name string, field IFObject) IKlass{
	me.staticField[name] = field
	return me
}

//动态属性
func  (me *Klass)setInsMethod(x string ,y IFObject) IKlass{
	me.insMethod[x] = y
	return me
}
//动态属性
func  (me *Klass)setInsField(x string ,y IFObject) IKlass{
	me.insField[x] = y
	return me
}


//供实例调用的方法
func  (me *Klass)setObMethod(x IFObject,  y IFObject, z IFObject) IKlass{
	x.getMethodDict().(*FMap).set(y.(*FString).GetVal(),z)
	return me
}

//供实例调用的方法
func  (me *Klass)setObField(x IFObject,  y IFObject, z IFObject) IKlass{
	x.getFieldDict().(*FMap).set(y.(*FString).GetVal(),z)
	return me
}




