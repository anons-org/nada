package Vm

/**
	Klass接口
 */
type IKlass interface {

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




	//给实例的操作 ------------
	setObMethod(x IFObject,  y IFObject, z IFObject) IKlass
	//这方法 所有的KLASS 都应该必须实现
	setObField(x IFObject,  y IFObject,z IFObject) IKlass









//每个Klass都应该有分配具体实例的方法
	Alloc() *IFObject

	Init() IKlass




}