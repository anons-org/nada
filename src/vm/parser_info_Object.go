package Vm


/**
	所有 的class中的 属性都要实现
 */
type IAttributes interface {
	getName() string
}



type Attributes struct {
	//u2
	nameIdx uint16
	//u4 属性的数据长度
	len int
	//当前属性的分类的英文描述
	typeName string
	//数据 数据长度为len 这玩意儿里面的包含的属性类型太多了。。。
	//具体参考资料吧....
	bytes []byte
}




type FieldInfo struct {
	//访问标志
	accessFlag uint16
	//
	nameIdx uint16

	descIdx uint16
	//字段所包含的属性的个数
	attrCount uint16

	//属性信息 仅仅是把属性存的是字节
	attriButes []*Attributes

}


type MethodInfo struct {
	//访问标志
	accessFlag uint16
	//
	nameIdx uint16

	name string

	descIdx uint16
	//字段所包含的属性的个数
	attrCount uint16

	//属性信息
	attriButes []IAttributes

}

/**
	code属性位于方法的属性表中
 */
type CodeAttr struct {

	//u2
	nameIdx uint16
	//u4 属性的数据长度
	len int
	//栈最大深度
	maxStack uint16

	//变量最大个数
	maxLocals uint16

	//代码长度
	codeLen uint32

	//具体指令
	code []byte

	//异常表长度
	expTableLen uint16

	//code的附加属性
	attrCount uint16

	attr []IAttributes

	//当前属性的分类的英文描述
	name string

}

func (me *CodeAttr ) getName()  string{
	return me.name
}

//---------------------------------------------------------
type LineNumberTableAttr struct {

	nameIdx uint16
	//当前属性的长度，不包括 nameIdx和attrLen的长度，这么做是方便虚拟机不处理当前属性，直接截取字节就行了
	attrLen uint32
	//属性的长度
	lineLen uint16
	name string
}

func (me *LineNumberTableAttr ) getName()  string{
	return me.name
}