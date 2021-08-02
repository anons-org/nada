package Vm

var TRUE *FString
var FALSE *FString
var NONE *FString

//ACC_PUBLIC	0x0001	方法是否为public
//ACC_PRIVATE	0x0002	方法是否为private
//ACC_PROTECTED	0x0004	方法是否为protected
//ACC_STATIC	0x0008	方法是否为static
//ACC_FINAL	0x0010	方法是否为final
//ACC_SYHCHRONRIZED	0x0020	方法是否为synchronized
//ACC_BRIDGE	0x0040	方法是否是有编译器产生的方法
//ACC_VARARGS	0x0080	方法是否接受参数
//ACC_NATIVE	0x0100	方法是否为native
//ACC_ABSTRACT	0x0400	方法是否为abstract
//ACC_STRICTFP 0x0800	方法是否为	strictfp
//ACC_SYNTHETIC	0x1000	方法是否是有编译器自动产生的

//访问标识
var ACC_PUBLIC = 0x0001
var ACC_PRIVATE = 0x0002
var ACC_STATIC = 0x0008
var ACC_SYNTHETIC = 0x1000

//-----------------------------

type OpcodeStruct struct {

	//指令所在行 有效PC地址
	n int
	//指令
	op uint8
	//操作数
	oper IFObject
}

type NSMethodArgType struct {
	args        []string
	argCount    int
	retArgs     []string
	retArgCount int
}

func (me *NSMethodArgType) getFieldDict() IFObject {
	panic("implement me")
}

func (me *NSMethodArgType) getMethodDict() IFObject {
	panic("implement me")
}

func (me *NSMethodArgType) getField(k string) IFObject {
	panic("implement me")
}

func (me *NSMethodArgType) getMethod(k string) IFObject {
	panic("implement me")
}

func (me *NSMethodArgType) setField(k string, v IFObject) {
	panic("implement me")
}

func (me *NSMethodArgType) setMethod(k string, v IFObject) {
	panic("implement me")
}

func (me *NSMethodArgType) GetKlass() IKlass {
	return nil
}

func (me *NSMethodArgType) SetKlass(kls IKlass) {

}

func (me *NSMethodArgType) test() IFObject {
	return me
}

func (me *NSMethodArgType) test() IFObject {
	return me
}

/**
创建klass
*/
func createKlass(name string, qualifier string, params ...IFObject) *Klass {

	kls := new(Klass)
	kls.name = NewFString(name)
	kls.qualifer = qualifier
	typeObj := new(TypeObject)
	typeObj.SetOnwKlass(kls)
	kls.SetTypeObject(typeObj)
	//静态
	kls.staticMethod = make(map[string]IFObject)
	kls.staticField = make(map[string]IFObject)
	//动态
	kls.insField = make(map[string]IFObject)
	kls.insMethod = make(map[string]IFObject)
	return kls
}

/**
创建实例
*/
func createKlassIns(callable IFObject, args *FArray) IFObject {
	ob := NewFObject()
	ob.SetKlass(callable.(*TypeObject).GetOnwKlass())
	return ob
}

/**
创建实例
*/
func createKlassIns(callable IFObject, args *FArray) IFObject {
	ob := NewFObject()
	ob.SetKlass(callable.(*TypeObject).GetOnwKlass())
	return ob
}

func addKlassToVm(k string, kls IKlass) {
	vms.metaKlass.set(k, kls)
}

func addKlassToVm(k string, kls IKlass) {
	vms.metaKlass.set(k, kls)
}

func CommonInit() {
	vms = new(Vms)
	vms.Build()
	TRUE = NewFString("true")
	FALSE = NewFString("false")
	NONE = NewFString("none")
}

func Start() {

	//fa:=NewFArray()
	//fa.add(NewFString("1"))
	//fa.add(NewFString("2"))
	//
	//fa.pop()

	vms.start()

}
