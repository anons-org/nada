package Vm

//方法引用
type ConstantMethodref struct {

	tag uint8
	classIdx uint16
	nameAndType uint16

	//所属的类
	rtClsName string
	//方法名
	rtFnName string
	//方法的参数 返回值类型
	rtFnType string

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
type ConstantInterfaceMethodRefInfo struct {

	tag uint8
	itfClsIdx uint16
	nameAndTypeIdx uint16
}

func (me* ConstantInterfaceMethodRefInfo ) getTag() uint8  {
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

//-------------------------------------------
//ConstantStringInfo info tag=6
type ConstantDoubleInfo struct {
	tag uint8
	//指向UTF8常量对象
	hByte []byte
	lByte []byte
}


func (me* ConstantDoubleInfo ) getTag() uint8  {
	return me.tag
}


//-------------------------------------------
//ConstantLongInfo info tag=5
type ConstantLongInfo struct {
	tag uint8
	//指向UTF8常量对象
	hByte []byte
	lByte []byte
}


func (me* ConstantLongInfo ) getTag() uint8  {
	return me.tag
}



//-------------------------------------------
//ConstantIntegerInfo info tag=3
type ConstantIntegerInfo struct {
	tag uint8
	bytes uint32
}


func (me* ConstantIntegerInfo ) getTag() uint8  {
	return me.tag
}


//-------------------------------------------
//ConstantFloatrInfo info tag=4
type ConstantFloatInfo struct {
	tag uint8
	bytes uint32
}


func (me* ConstantFloatInfo ) getTag() uint8  {
	return me.tag
}



//-------------------------------------------
//ConstantFieldrefInfo info tag=9
type ConstantFieldrefInfo struct {
	tag uint8
	clsIdx uint16
	nameAndTypeIdx uint16

}


func (me* ConstantFieldrefInfo ) getTag() uint8  {
	return me.tag
}


//-------------------------------------------
//ConstantInvokeDynInfo info tag=18
type ConstantInvokeDynInfo struct {
	tag uint8
	//class文件中的bootstrap_methods中的项
	bootMethodAttrIdx uint16
	nameAndTypeIdx uint16
}


func (me* ConstantInvokeDynInfo ) getTag() uint8  {
	return me.tag
}


//-------------------------------------------
//ConstantMethodHandleInfo info tag=15
type ConstantMethodHandleInfo struct {
	tag uint8
	refKind uint8
	refIdx uint16

}


func (me* ConstantMethodHandleInfo ) getTag() uint8  {
	return me.tag
}


//-------------------------------------------
//ConstantMethodTypeInfo info tag=16
type ConstantMethodTypeInfo struct {
	tag uint8
	//class文件中的bootstrap_methods中的项
	descIdx uint16
}


func (me* ConstantMethodTypeInfo ) getTag() uint8  {
	return me.tag
}













