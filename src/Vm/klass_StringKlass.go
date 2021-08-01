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

func (me *StringKlass) getObMethod(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *StringKlass) getObField(x IFObject, y IFObject) IFObject {
	panic("implement me")
}

func (me *StringKlass) setStaticMethod(name string, method IFObject) IKlass {
	panic("implement me")
}

func (me *StringKlass) setStaticField(name string, field IFObject) IKlass {
	panic("implement me")
}

func (me *StringKlass) setInsMethod(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *StringKlass) setInsField(x string, y IFObject) IKlass {
	panic("implement me")
}

func (me *StringKlass) getInsMethod(k string) IFObject {
	panic("implement me")
}

func (me *StringKlass) getInsField(k string) IFObject {
	panic("implement me")
}

func (me *StringKlass) getStaticMethod(name string) IFObject {
	panic("implement me")
}

func (me *StringKlass) getStaticField(name string) IFObject {
	panic("implement me")
}

func (me *StringKlass) setObMethod(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

func (me *StringKlass) setObField(x IFObject, y IFObject, z IFObject) IKlass {
	panic("implement me")
}

/**
	设置类型对象
*/
func (me *StringKlass)SetTypeObject(to *TypeObject) {
	me.typeObject = to
}

/**
	获取类型对象
 */
func (me *StringKlass)GetTypeObject() *TypeObject {
	return me.typeObject
}

func (me *StringKlass)GetName() string{
	return me.name
}

func (me *StringKlass)Alloc() *IFObject {
	return nil
}


/**
初始化Klass
*/
func (me *StringKlass)Init() IKlass {
	me.name = "StringKlass"
	typeObj := new(TypeObject).Build()
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