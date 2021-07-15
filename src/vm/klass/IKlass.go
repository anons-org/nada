package klass

import (
	"vm/type"
)

/**
	Klass接口
 */
type IKlass interface {

	//获取类型对象
	GetTypeObject()  *ftype.TypeObject
	SetTypeObject(	*ftype.TypeObject )

	//每个Klass都应该有分配具体实例的方法
	Alloc() *ftype.IFObject

	Init() IKlass

	//每个Klass应该有自己的字典
	Dict(k string) IKlass

 	GetDict(k string) ftype.IFObject

	//获取指定的方法 可能是原生
	GetKlassMethod(k string) ftype.IFObject

}