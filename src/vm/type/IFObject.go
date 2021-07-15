package ftype

import "vm/klass"

/**
	所有内置对象需要实现的接口
	在C++中，实际上是所有内置对象的父类
 */
type IFObject interface {

	GetKlass() klass.IKlass
	SetKlass(klass.IKlass)

}


