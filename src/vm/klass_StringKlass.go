package Vm

import (
	"fmt"
)

type StringKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *TypeObject

	//方法集合
	dict map[string]IFObject

}

/**
	获取方法
 */
func (me StringKlass) GetKlassMethod(k string) IFObject {
	return me.dict[k]
}

func (me StringKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me StringKlass) GetDict(k string) IFObject {
	return me.dict[k]
}



/**
	设置类型对象
*/
func (me StringKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
	获取类型对象
 */
func (me StringKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me StringKlass)GetName() string{
	return me.name
}

func (me StringKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me StringKlass)Init() IKlass {
	me.name = "StringKlass"
	typeObj := new(TypeObject)
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	me.dict = make(map[string]IFObject)
	me.dict["indexOf"] = new(FMethod).Build(1).SetCall(indexOf)
	me.dict["print"] = new(FMethod).Build(1).SetCall(print)

	return me
}



/**
String print
*/
func print(args []IFObject) IFObject {
	fmt.Print("测试原生方法.....")
	return nil
}



/**
	String indexOf
 */
func indexOf(args []IFObject) IFObject {

	return nil
}

/**
String len
*/
func len(args []IFObject) IFObject {

	return nil
}

/**
String split
*/
func split(args []IFObject) IFObject {

	return nil
}

/**
String join
*/

func join(args []IFObject) IFObject {

	return nil
}