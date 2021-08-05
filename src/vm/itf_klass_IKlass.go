package vm

/**
	Klass接口
 */
type IKlass interface {
	//获取KLASS的名字
	getName() string
	//获取类型对象
	GetTypeObject()  *TypeObject
	SetTypeObject(	*TypeObject)

	setStaticMethod(name string, method IFObject) IKlass

	//添加静态属性
	setStaticField(name string, field IFObject) IKlass


	setInsMethod(x string , y IFObject) IKlass

	setInsField(x string , y IFObject) IKlass



	getInsMethod(k string) IFObject

	getInsField(k string) IFObject


	getStaticMethod(name string) IFObject

	//添加静态属性
	getStaticField(name string) IFObject




	//给实例的操作 实例是没有静态方法和静态属性的!!------------
	setObMethod(x IFObject,  y IFObject, z IFObject) IKlass
	//这方法 所有的KLASS 都应该必须实现
	setObField(x IFObject,  y IFObject,z IFObject) IKlass

	getObMethod (x IFObject,  y IFObject) IFObject
	getObField (x IFObject,  y IFObject) IFObject


	//-----------------对外的接口
	GetStaticMethod(name string) IFObject





//每个Klass都应该有分配具体实例的方法
	Alloc() *IFObject

	Init() IKlass




}