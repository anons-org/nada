package klass

import (
	"fmt"
	"vm/type"
)

type StringKlass struct {
	//klass名称
	name string
	//类型对象
	typeObject *ftype.TypeObject

	//方法集合
	dict map[string]ftype.IFObject

}

/**
	获取方法
 */
func (me StringKlass) GetKlassMethod(k string) ftype.IFObject {
	return me.dict[k]
}

func (me StringKlass) Dict(k string) IKlass {
	panic("implement me")
}


func (me StringKlass) GetDict(k string) ftype.IFObject {
	return me.dict[k]
}



/**
	设置类型对象
*/
func (me StringKlass)SetTypeObject(to *ftype.TypeObject) {
	me.typeObject = to
}

/**
	获取类型对象
 */
func (me StringKlass)GetTypeObject() *ftype.TypeObject {
	return me.typeObject
}

func (me StringKlass)GetName() string{
	return me.name
}

func (me StringKlass)Alloc() *ftype.IFObject {
	return nil
}


/**
初始化Klass
*/
func (me StringKlass)Init() IKlass {
	me.name = "StringKlass"
	typeObj := new(ftype.TypeObject)
	typeObj.SetOnwKlass(me)
	me.SetTypeObject( typeObj )
	me.dict = make(map[string]ftype.IFObject)
	me.dict["indexOf"] = new(ftype.FMethod).Build(1).SetCall(indexOf)
	me.dict["print"] = new(ftype.FMethod).Build(1).SetCall(print)

	return me
}



/**
String print
*/
func print(args []ftype.IFObject) ftype.IFObject {
	fmt.Print("测试原生方法.....")
	return nil
}



/**
	String indexOf
 */
func indexOf(args []ftype.IFObject) ftype.IFObject {

	return nil
}

/**
String len
*/
func len(args []ftype.IFObject) ftype.IFObject {

	return nil
}

/**
String split
*/
func split(args []ftype.IFObject) ftype.IFObject {

	return nil
}

/**
String join
*/

func join(args []ftype.IFObject) ftype.IFObject {

	return nil
}