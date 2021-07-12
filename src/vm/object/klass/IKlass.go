package VmObjectKlass

import VmObject "vm/object"

/**
	Klass接口
 */
type IKlass interface {
	//获取类型对象
	GetTypeObject()  *VmObject.TypeObject
	SetTypeObject(*VmObject.TypeObject)

	//每个Klass都应该有分配具体实例的方法
	Alloc() *VmObject.IFObject

	Init() *MethodKlass

}