package VmObject

import (
	"./klass"
)

type FMethod struct {
	//方法名
	mName string
	//方法字节码
	code []byte
	klass *VmObjectKlass.MethodKlass


}


func (me *FMethod) ValToString(){

}


/**
	设置方法名
 */
func (me *FMethod) SetMethod(name string) {

}

/**
	绑定klass
 */
func (me *FMethod) SetKlass(klass *VmObjectKlass.MethodKlass) {
	me.klass = klass;
}


/**
	初始化方法
 */
func (me *FMethod) Build() *FMethod {
	//设置class
	me.SetKlass(  new( VmObjectKlass.MethodKlass ).Init()  )
	return me
}




