package Vm

//方法引用
type ConstantMethodref struct {

	tag uint8
	classIdx uint16
	nameAndType uint16

}

func (me* ConstantMethodref ) getTag() uint8  {
	return 0
}

//-----------------------------------------------


//类引用
type ConstantClassInfo struct {
	tag uint8
	nameIdx uint16
}

func (me* ConstantClassInfo ) getTag() uint8  {
	return me.tag
}

//------------------------------------------------



//接口 tag=11
type ConstantInterfaceMethodInfo struct {

	tag uint8
	itfClsIdx uint16
	nameAndTypeIdx uint16
}

func (me* ConstantInterfaceMethodInfo ) getTag() uint8  {
	return me.tag
}

//--------------------------------------------------



//接口 tag=12
type ConstantNameAndType struct {
	tag uint8
	nameIdx uint16
	descIdx uint16
}

func (me* ConstantNameAndType ) getTag() uint8  {
	return me.tag
}

//-----------------------------------------------------


//utf8字符串 tag=1
type ConstantUtf8 struct {
	tag uint8
	len uint16
	bytes []byte
}

func (me* ConstantUtf8 ) getTag() uint8  {
	return me.tag
}
//------------------------------------------------



//string info tag=8
type ConstantStringInfo struct {
	tag uint8
	//指向UTF8常量对象
	strIdx uint16
}

func (me* ConstantStringInfo ) getTag() uint8  {
	return me.tag
}









