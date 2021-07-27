package Vm

/**
	Klass接口
 */
type IKlass interface {

	//获取类型对象
	GetTypeObject()  *TypeObject
	SetTypeObject(	*TypeObject)

    getStaticDict(k string) IFObject



//每个Klass都应该有分配具体实例的方法
	Alloc() *IFObject

	Init() IKlass

	//每个Klass应该有自己的字典
	Dict(k string) IKlass

 	GetDict(k string) IFObject



}