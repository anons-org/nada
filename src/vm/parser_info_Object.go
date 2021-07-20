package Vm


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

	descIdx uint16
	//字段所包含的属性的个数
	attrCount uint16

	//属性信息 仅仅是把属性存的是字节
	attriButes []*Attributes

}


