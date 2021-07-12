package VmObject

/**
	所有内置对象需要实现的接口
	在C++中，实际上是所有内置对象的父类
 */
type IFObject interface {
	ValToString()
	GetKlass()
}


