package Vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	Utils "utils"

	"os"
)

/**
Class在解析的过程中，
需要把常量转化为实际的数据类型
这样可以减少解释器执行过程中对常量转换类型的消耗
数据类型请看 type_  开头的文件
class的解析并不是多神秘的事情，照着规范一个个字节处理就行了
没啥技术含量，就是细致活儿量大而已...
by mike
*/

type ClassByteFile struct {
	constanPool *Map

	//当前JAVA字节码读取位置
	readIdx int
	//文件的名字
	fileName string
	//CLASS 字节码
	data []byte

	//当前class常量数量
	constanCount int
	//由于double long 占两字节，导致，常量池数量和实际解析数量不一致，加一个变量来做判断吧
	constantObjSpan int

	//常量池下标
	idx int
	//访问标识
	accessFlags uint16
	//自己的类
	thisClass uint16
	//父类
	superClass uint16
	//接口数量
	itfCount uint16

	//字段个数
	fieldCount uint16

	//字段数据
	fields []*FieldInfo

	//方法个数
	methodCount uint16

	methods []*MethodInfo
	//这个是整个class文件的属性总数，和fields、methods的不一样!
	attrsCount uint16
	attrs      []*Attributes
}

func (me *ClassByteFile) Load(fileName string) *ClassByteFile {

	//fmt.Print(  CommonConstan.G_OT_DEF[CommonConstan.G_OT.STRING].Name);

	//	var fileName = "E:\\AAAA_CODE\\new-eclipse-workspace\\far-dev\\demo\\TestDemo1.class"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return me
	}

	stat, _ := file.Stat()

	fmt.Print(stat.Size())

	buf := make([]byte, stat.Size())

	//var data []byte;

	for {
		_, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)

			}
		}
		//data = append(data,buf);
		//fmt.Println(length, string(buf));
	}

	me.data = buf
	return me
}

/**

 */
func (me *ClassByteFile) Read(n int) []byte {
	idx := me.readIdx
	me.readIdx += n
	return me.data[idx : idx+n]
}

func (me *ClassByteFile) ReadShort() uint16 {
	d := me.Read(2)
	return binary.BigEndian.Uint16(d)
}

func (me *ClassByteFile) ReadUint32() uint32 {
	d := me.Read(4)

	return binary.BigEndian.Uint32(d)
}

func (me *ClassByteFile) ReadUi8() int {
	//d := me.Read(2);
	//return binary.BigEndian.Uint16(d);
	var x int8
	d := me.Read(1)

	bytesBuffer := bytes.NewBuffer(d)

	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(uint8(x))

}

/**
读取常量池对象
*/
func (me *ClassByteFile) parserConstanInteger(idx int) {

}

/**
方法解析
*/
func (me *ClassByteFile) parserMethodref(idx uint8) {
	//常量池的对象，需要转成VM中的内置类型对象
	mt := new(ConstantMethodref)
	mt.tag = idx
	mt.classIdx = me.ReadShort()
	mt.nameAndType = me.ReadShort()
	me.constanPool.Set(me.idx, mt)
	me.idx++
}

/**
类引用解析
*/
func (me *ClassByteFile) parserClassInfo(idx uint8) {
	cls := new(ConstantClassInfo)
	cls.tag = idx
	cls.nameIdx = me.ReadShort()
	me.constanPool.Set(me.idx, cls)
	me.idx++
}

/**
接口解析
*/
func (me *ClassByteFile) parserItfMethodInfo(idx uint8) {
	itf := new(ConstantInterfaceMethodRefInfo)
	itf.tag = idx
	//2字节
	itf.itfClsIdx = me.ReadShort()
	//2字节
	itf.nameAndTypeIdx = me.ReadShort()
	me.constanPool.Set(me.idx, itf)
	me.idx++
}

/**

名称和类型解析

程序员最大的问题就是容易认真，认真起来就喜欢抬杠
今天同事问我：老大，车间用两台电脑直接通过C#的scoket互联保证稳定性行的通不(被车间那帮家伙打报告了)，我就简单说了两句，另外一个同事，马上质疑如此这般
<其实这同事也就2年JAVA开发经验，别说操作系统和更底层的玩意儿，应该连socket的有多少中模式都还不清楚,哪儿来的的勇气质疑，请看段子开头>
对于他的质疑，我没有反驳，感觉没必要，思维不在一个层面，所以避重就轻的说：技术都是为业务服务的，只要你能保证车间工序正常运作，不出现生产事故就行

朋友千个少，敌人一个多


*/
func (me *ClassByteFile) parserNameAndType(idx uint8) {
	nt := new(ConstantNameAndType)
	nt.tag = idx
	nt.nameIdx = me.ReadShort()
	nt.descIdx = me.ReadShort()
	me.constanPool.Set(me.idx, nt)
	me.idx++
}

/**
utf8解析
*/
func (me *ClassByteFile) parserUtf8(idx uint8) {

	utf := new(ConstantUtf8)
	utf.tag = idx
	utf.len = me.ReadShort()
	//读取字节码 未完成，人啊，这辈子，除了生死，都是小事...先回家吃个饭 睡觉..
	utf.bytes = me.Read(int(utf.len))
	me.constanPool.Set(me.idx, utf)
	me.idx++
}

/**
string字符解析
*/
func (me *ClassByteFile) parserStringInfo(idx uint8) {

	str := new(ConstantStringInfo)
	str.tag = idx
	str.strIdx = me.ReadShort()
	me.constanPool.Set(me.idx, str)
	me.idx++
}

/**
字段解析
*/
func (me *ClassByteFile) parserFieldInfo() {

	//字段个数
	me.fieldCount = me.ReadShort()

	if me.fieldCount <= 0 { //没有字段就不处理了...
		return
	}

	for i := 0; i < int(me.fieldCount); i++ {

		field := new(FieldInfo)
		field.accessFlag = me.ReadShort()
		field.nameIdx = me.ReadShort()
		field.descIdx = me.ReadShort()
		field.attrCount = me.ReadShort()

		for i := 0; i < int(field.attrCount); i++ {
			attr := new(Attributes)
			attr.nameIdx = me.ReadShort()
			attr.len = int(me.ReadUint32())
			attr.bytes = me.Read(attr.len)
			field.attriButes = append(field.attriButes, attr)
		}
		me.fields = append(me.fields, field)
	}

}

//class文件的attr属性解析
func (me *ClassByteFile) parserAttrsInfo() {
	me.attrsCount = me.ReadShort()

	if me.attrsCount <= 0 { //没有字段就不处理了...
		return
	}

	for i := 0; i < int(me.attrsCount); i++ {
		attr := new(Attributes)
		attr.nameIdx = me.ReadShort()
		attr.len = int(me.ReadUint32())
		attr.bytes = me.Read(attr.len)
		me.attrs = append(me.attrs, attr)
	}

}

/**
方法解析
*/
func (me *ClassByteFile) parserMethodInfo() {

	me.methodCount = me.ReadShort()

	if me.methodCount <= 0 { //没有字段就不处理了...
		return
	}

	for i := 0; i < int(me.methodCount); i++ {

		method := new(MethodInfo)
		method.accessFlag = me.ReadShort()
		method.nameIdx = me.ReadShort()
		method.name = ctConvStr(me, method.nameIdx)
		method.descIdx = me.ReadShort()
		method.attrCount = me.ReadShort()
		//方法的属性解析
		for i := 0; i < int(method.attrCount); i++ {
			//先读出属性名
			nameIdx := me.ReadShort()
			aName := ctConvStr(me, nameIdx)
			if aName == "Code" {
				attr := new(CodeAttr)
				attr.nameIdx = nameIdx
				attr.name = aName
				attr.len = int(me.ReadUint32())
				attr.maxStack = me.ReadShort()
				attr.maxLocals = me.ReadShort()
				attr.codeLen = me.ReadUint32()
				attr.code = me.Read(int(attr.codeLen))
				attr.expTableLen = me.ReadShort()
				attr.attrCount = me.ReadShort()
				//code自己附加的属性
				for i := 0; i < int(attr.attrCount); i++ {
					nameIdx = me.ReadShort()
					aName := ctConvStr(me, nameIdx)
					if aName == "LineNumberTable" {
						lAttr := new(LineNumberTableAttr)
						lAttr.nameIdx = nameIdx
						lAttr.name = aName
						lAttr.attrLen = me.ReadUint32()
						//读出来，不处理
						me.Read(int(lAttr.attrLen))
					} else if aName == "StackMapTable" {
						//先偷懒不处理。。。。
						attrLen := me.ReadUint32()
						me.Read(int(attrLen))
					}
				}
				method.attriButes = append(method.attriButes, attr)
			} else {
				aa := 0
				aa++
			}

		}
		me.methods = append(me.methods, method)
	}

}

func (me *ClassByteFile) parserIntegerInfo(idx uint8) {
	o := new(ConstantIntegerInfo)
	o.tag = idx
	o.bytes = me.ReadUint32()
	me.constanPool.Set(me.idx, o)
	me.idx++
}

func (me *ClassByteFile) parserFloatInfo(idx uint8) {
	o := new(ConstantFloatInfo)
	o.tag = idx
	o.bytes = me.ReadUint32()
	me.constanPool.Set(me.idx, o)
	me.idx++
}

func (me *ClassByteFile) parserDoubleInfo(idx uint8) {

	o := new(ConstantDoubleInfo)
	o.tag = idx
	o.hByte = me.Read(4)
	o.lByte = me.Read(4)
	me.constanPool.Set(me.idx, o)
	me.idx += 2
	me.constantObjSpan++
}

func (me *ClassByteFile) parserLongInfo(idx uint8) {

	o := new(ConstantLongInfo)
	o.tag = idx
	o.hByte = me.Read(4)
	o.lByte = me.Read(4)
	me.constanPool.Set(me.idx, o)
	me.idx += 2
	me.constantObjSpan++
}

func (me *ClassByteFile) parserFieldrefInfo(idx uint8) {
	o := new(ConstantFieldrefInfo)
	o.tag = idx
	o.clsIdx = me.ReadShort()
	o.nameAndTypeIdx = me.ReadShort()
	me.constanPool.Set(me.idx, o)
	me.idx++

}

func (me *ClassByteFile) parserInvokeDynInfo(idx uint8) {
	o := new(ConstantInvokeDynInfo)
	o.tag = idx
	o.bootMethodAttrIdx = me.ReadShort()
	o.nameAndTypeIdx = me.ReadShort()
	me.constanPool.Set(me.idx, o)
	me.idx++
}

func (me *ClassByteFile) parserMethodHandleInfo(idx uint8) {
	o := new(ConstantMethodHandleInfo)
	o.tag = idx
	o.refKind = uint8(me.ReadUi8())
	o.refIdx = me.ReadShort()
	me.constanPool.Set(me.idx, o)
	me.idx++
}

func (me *ClassByteFile) parserMethodTypeInfo(idx uint8) {
	o := new(ConstantMethodTypeInfo)
	o.tag = idx
	o.descIdx = me.ReadShort()
	me.constanPool.Set(me.idx, o)
	me.idx++
}

/**

常量池内容解析
注意 Long double 这些类型会占用多个字节
导致常量池编号(下标)会移位!!
这在各种介绍JVM的资料中都有介绍,但没有明确说明会导致这个问题

*/
func (me *ClassByteFile) parserConstansPool() bool {

	x := int(me.ReadUi8())

	fmt.Printf("当前常量编号%d\n", me.idx)

	fg := true
	switch x {
	case 1:
		me.parserUtf8(uint8(x))

	case 3:
		me.parserIntegerInfo(uint8(x))
	case 4:
		me.parserFloatInfo(uint8(x))
	case 5:
		me.parserLongInfo(uint8(x))
	case 6:
		me.parserDoubleInfo(uint8(x))
	case 7:
		me.parserClassInfo(uint8(x))
	case 8:
		me.parserStringInfo(uint8(x))
	case 9:
		me.parserFieldrefInfo(uint8(x))
	case 10:
		me.parserMethodref(uint8(x))
	case 11:
		me.parserItfMethodInfo(uint8(x))
	case 12:
		me.parserNameAndType(uint8(x))

	case 15:
		me.parserMethodHandleInfo(uint8(x))
	case 16:
		me.parserMethodTypeInfo(uint8(x))

	case 18:
		me.parserInvokeDynInfo(uint8(x))

	default:
		fg = false

	}
	return fg

}

/**
解析class
*/
func (me *ClassByteFile) Parser() {

	data := me.Read(4)
	//能正常使用
	fmt.Println(binary.BigEndian.Uint32(data) == 0xcafebabe)
	version := me.ReadShort()
	masterVersion := me.ReadShort()
	constanCount := me.ReadShort()
	me.constanCount = int(constanCount)

	fmt.Print(fmt.Sprintf("主版本%d,副版本%d,常量池%d\n", masterVersion, version, constanCount))

	//常量池下标是从零开始的，需要加一个空值
	me.constanPool.Set(me.idx, struct{}{})
	me.idx++

	for i := 0; i < me.constanCount-1; i++ {

		if me.constanPool.Size()+me.constantObjSpan > me.constanCount-1 {
			break
		}
		me.parserConstansPool()
	}

	//for ;me.parserConstansPool();{
	//
	//}

	me.print()

	//跳过10个字节 parserInterfaceAndSuper
	//访问标志、类索引(当前类索引)(this结构)、父类索引、接口索引集合。
	//查阅了资料 这10字节分别是class结构中的:
	// u2 access_flags
	// u2 this_class
	// u2 super_class
	// u2 interface_count 如果 count为0 下面的值就不应该存在了 所以当前的测试的class文件只读取8字节
	// u2 interface[interface_count]

	me.Read(6)

	//接扣数量
	me.itfCount = me.ReadShort()

	//解析字段 目前还有问题
	me.parserFieldInfo()
	me.parserMethodInfo()
	//属性段
	me.parserAttrsInfo()

	//生成KLASS
	me.buildKlass()

	s := fmt.Sprintf("pool sieze %d", me.constanPool.Size())
	fmt.Println(s)
	fmt.Print("解析class完成")

	//fmt.Print(d);
}

func (me *ClassByteFile) print() {

	me.constanPool.ForEach(func(e *MapEntry) {
		var s string = ""
		key, _ := e.key.(int)
		if v, ok := e.val.(*ConstantMethodref); ok {
			s = fmt.Sprintf("#%d = %s					#%d.#%d", key, "Methodref", v.classIdx, v.nameAndType)
		} else if _, ok := e.val.(*ConstantStringInfo); ok {
			s = fmt.Sprintf("#%d = %s					", key, "String")
		} else if v, ok := e.val.(*ConstantClassInfo); ok {
			s = fmt.Sprintf("#%d = %s					#%d", e.key, "Class", v.nameIdx)
		} else if v, ok := e.val.(*ConstantInterfaceMethodRefInfo); ok {
			s = fmt.Sprintf("#%d = %s					#%d.#%d", e.key, "InterfaceMethodref", v.itfClsIdx, v.nameAndTypeIdx)
		} else if v, ok := e.val.(*ConstantNameAndType); ok {
			s = fmt.Sprintf("#%d = %s					#%d.#%d", e.key, "NameAndType", v.nameIdx, v.descIdx)
		} else if utf8, ok := e.val.(*ConstantUtf8); ok {
			s = fmt.Sprintf("#%d = %s					%s", e.key, "Utf8", string(utf8.bytes))
		} else if intOb, ok := e.val.(*ConstantIntegerInfo); ok {
			s = fmt.Sprintf("#%d = %s					%s", e.key, "Integer", int(intOb.bytes))
		} else if floatOb, ok := e.val.(*ConstantFloatInfo); ok {
			s = fmt.Sprintf("#%d = %s					%s", e.key, "FLoat", float64(floatOb.bytes))
		} else if _, ok := e.val.(*ConstantDoubleInfo); ok {
			s = fmt.Sprintf("#%d = %s					%s", e.key, "Double", 1)
		} else if _, ok := e.val.(*ConstantLongInfo); ok {
			s = fmt.Sprintf("#%d = %s					%s", e.key, "Long", 1)
		} else if invOb, ok := e.val.(*ConstantInvokeDynInfo); ok {
			s = fmt.Sprintf("#%d = %s					#%d.#%d", e.key, "InvokeDynamic", invOb.bootMethodAttrIdx, invOb.nameAndTypeIdx)
		} else if invOb, ok := e.val.(*ConstantFieldrefInfo); ok {
			s = fmt.Sprintf("#%d = %s					#%d.#%d", e.key, "Fieldref", invOb.clsIdx, invOb.nameAndTypeIdx)
		} else if invOb, ok := e.val.(*ConstantMethodHandleInfo); ok {
			s = fmt.Sprintf("#%d = %s					#%d.#%d", e.key, "MethodHandle ", invOb.refKind, invOb.refIdx)
		} else if invOb, ok := e.val.(*ConstantMethodTypeInfo); ok {
			s = fmt.Sprintf("#%d = %s					#%d", e.key, "MethodType ", invOb.descIdx)
		}
		fmt.Println(s)
	})

}

func (me *ClassByteFile) get(n int) {

}

/**
initialize operands in advance
*/
func (me *ClassByteFile) bulidImmediate(method *FMethod, code []byte) {
	//codes:=bytes.NewBuffer(code)
	//codes.Read(2);

	data := NewBytes(code)

	for data.idx < len(code) {
		op := data.read(1)[0]
		switch op {
		case OP.ALOAD_0:
			method.codeList[op] = nil
		case OP.INVOKESPECIAL:
			//https://www.cnblogs.com/cfas/p/15063669.html
			mt := me.getCtMethodRef(binary.BigEndian.Uint16(data.read(2)))
			fmt := NewFMethod(METHOD_TYPE_VIRTUAL, mt.rtFnName)
			method.codeList[op] = fmt
		case OP.RETURN:
			method.codeList[op] = nil
		case OP.LDC:
			idx := data.read(1)[0]
			v := me.constanPool.Get(idx)
			switch v.(type) {
			case *ConstantIntegerInfo, *ConstantFloatInfo:
				method.codeList[op] = nil
			case *ConstantStringInfo:
				method.codeList[op] = me.getCtStringToFString(v.(*ConstantStringInfo).strIdx)
			default:
				fmt.Print("The operand of the LDC instruction is not Int or float when the operand is generated")
			}
		case OP.ASTORE_0, OP.ASTORE_1, OP.ASTORE_2, OP.ASTORE_3: //don't need operands
			method.codeList[op] = nil
		case OP.LDC_W:
			idx := binary.BigEndian.Uint16(data.read(2))
			v := me.constanPool.Get(idx)
			switch v.(type) {
			case *ConstantIntegerInfo, *ConstantFloatInfo:
				method.codeList[op] = nil
			case *ConstantStringInfo:
				method.codeList[op] = me.getCtStringToFString(v.(*ConstantStringInfo).strIdx)
			case *ConstantDoubleInfo:
				hb := v.(*ConstantDoubleInfo).hByte
				lb := v.(*ConstantDoubleInfo).lByte
				hb = append(hb, lb...)
				fVal := Utils.Float64frombytes(hb)
				method.codeList[op] = NewFFLoat(fVal)
			case *ConstantLongInfo:

			default:
				fmt.Print("The operand of the LDC instruction is not Int or float when the operand is generated")
			}
		default:
			fmt.Print("failed to initialize operands in advance. procedure...")
		}
	}

}

func (me *ClassByteFile) buildKlass() {

	kls := createKlass("")
	kls.name = NewFString("Test")
	//each method
	for _, v := range me.methods {

		desc := ctConvStr(me, v.descIdx)
		desc += ""
		name := ctConvStr(me, v.nameIdx)
		mt := NewFMethod(METHOD_TYPE_VIRTUAL, name)
		//each method att
		for _, va := range v.attriButes {
			//deal with code attr
			if vaOb, ok := va.(*CodeAttr); ok {
				mt.setCode(vaOb.code)
				//advance processing immediate may optimize performance
				me.bulidImmediate(mt, vaOb.code)
			}
		}
		//deal with of accessFlag
		if int(v.accessFlag) == ACC_PUBLIC|ACC_STATIC || int(v.accessFlag) == ACC_PRIVATE|ACC_STATIC || int(v.accessFlag) == ACC_PRIVATE|ACC_STATIC|ACC_SYNTHETIC {
			kls.addStaticMethod(name, mt)
		} else {
			kls.addInsMethod(name, mt)
		}
	}

	me.addKlass(kls)

}

func NewClassByteFile() *ClassByteFile {
	e := new(ClassByteFile)
	e.constanPool = NewMap()
	return e
}

/**
添加Klass
*/
func (me *ClassByteFile) addKlass(kls *Klass) {
	vms.metaKlass.set("nada.lib.TestKlass", kls)
}

/**
获取常量池中的Methodref
*/
func (me *ClassByteFile) getCtMethodRef(idx uint16) *ConstantMethodref {
	val := me.constanPool.Get(idx)
	if mt, ok := val.(*ConstantMethodref); ok {
		//完善该对象的值
		cls := me.getCtClass(mt.classIdx)
		mt.rtClsName = ctConvStr(me, cls.nameIdx)
		nt := me.getCtNameAndType(mt.nameAndType)
		mt.rtFnName = ctConvStr(me, nt.nameIdx)
		mt.rtFnType = ctConvStr(me, nt.descIdx)
		return mt
	}
	return nil
}

/**
获取常量池中的Class
*/
func (me *ClassByteFile) getCtClass(idx uint16) *ConstantClassInfo {
	val := me.constanPool.Get(idx)
	if cls, ok := val.(*ConstantClassInfo); ok {
		return cls
	}
	return nil
}

/**
获取常量池中的NameAndType
*/
func (me *ClassByteFile) getCtNameAndType(idx uint16) *ConstantNameAndType {
	val := me.constanPool.Get(idx)
	if nt, ok := val.(*ConstantNameAndType); ok {
		return nt
	}
	return nil
}

func (me *ClassByteFile) getCtStringToFString(idx uint16) *FString {
	s := me.getCtString(idx)
	s2 := NewFString(s)
	return s2
}

func (me *ClassByteFile) getCtString(idx uint16) string {
	val := me.constanPool.Get(idx)
	if utf8, ok := val.(*ConstantUtf8); ok {
		return string(utf8.bytes)
	}
	return ""
}

/**
常量对象转string
*/
func ctConvStr(me *ClassByteFile, idx uint16) string {
	val := me.constanPool.Get(idx)
	if utf8, ok := val.(*ConstantUtf8); ok {
		return string(utf8.bytes)
	}
	return ""
}

func ct2FString(me *ClassByteFile, idx uint16) *FString {
	s := ctConvStr(me, idx)
	s2 := NewFString(s)
	return s2
}

//func Uint16(b []byte) uint16 {
//	_ = b[1]
//	return uint16(b[0]) | uint16(b[1])<<8
//}
//
//
//func BytesToUint32(b []byte) uint32 {
//	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
//	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
//}
//
//
//func Uint64(b []byte) uint64 {
//	_ = b[7]
//	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |        uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
//}
//
//
