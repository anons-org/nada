package Vm



/**
	所有内置对象需要实现的接口
	在C++中，实际上是所有内置对象的父类
 */
type IFObject interface {

	GetKlass() IKlass
	SetKlass(IKlass)

	//优先查找自己的 没有查找Klass的
	getField(k string) IFObject
	//优先查找自己的 没有查找Klass的
	getMethod(k string) IFObject

	//设置实例自己的字段
	setField(k string, v IFObject)
	//设置实例自己的方法
	setMethod(k string, v IFObject)

	getFieldDict() IFObject

	getMethodDict() IFObject

}


