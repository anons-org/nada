package Vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)
/**
	Class在解析的过程中，
	需要把常量转化为实际的数据类型
	这样可以减少解释器执行过程中对常量转换类型的消耗
	数据类型请看 type_  开头的文件
	by mike
 */


type ClassByteFile struct {

	constanPool []IConstantObject

	//当前JAVA字节码读取位置
	readIdx int
	//文件的名字
	fileName string
	//CLASS 字节码
	data []byte

	//当前class常量数量
	constanCount int
}

func (me *ClassByteFile) Load(fileName string) *ClassByteFile {

	//fmt.Print(  CommonConstan.G_OT_DEF[CommonConstan.G_OT.STRING].Name);

//	var fileName = "E:\\AAAA_CODE\\new-eclipse-workspace\\far-dev\\demo\\TestDemo1.class"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666);
	if err != nil {
		fmt.Println("Open file error!", err);
		return me;
	}

	stat, _ := file.Stat();

	fmt.Print(stat.Size());

	buf := make([]byte, stat.Size());

	//var data []byte;



	for {
		_, err := file.Read(buf);
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

	me.data = buf;
	return  me;
}
/**

 */
func (me *ClassByteFile) Read(n int) []byte {
	idx := me.readIdx;
	me.readIdx	+=n;
	return me.data[idx:idx+n];
}

func (me *ClassByteFile) ReadShort() uint16 {
	d := me.Read(2);
	return binary.BigEndian.Uint16(d);
}

func (me *ClassByteFile) ReadUi8() int {
	//d := me.Read(2);
	//return binary.BigEndian.Uint16(d);
	var x int8
	d := me.Read(1);

	bytesBuffer := bytes.NewBuffer(d)

	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(uint8(x))

}


/**
读取常量池对象
 */
func (me *ClassByteFile) parserConstanInteger(idx int) {
	//integer := new(Constans.ConstanInteger)
	//me.constanPool = append(me.constanPool,integer)
}

/**
	方法解析
 */
func (me *ClassByteFile) parserConstanMethod(idx uint8) {
	//常量池的对象，需要转成VM中的内置类型对象

	//mt := new( VmObject.FMethod ).Build();

	mt:=new(ConstantMethodref)
	mt.tag = idx
	mt.classIdx = me.ReadShort()
	mt.nameAndType = me.ReadShort()
	me.constanPool = append(me.constanPool,mt)


}

/**
	类引用解析
 */
func (me *ClassByteFile) parserConstanClassInfo(idx uint8) {
	cls:=new(ConstantClassInfo)
	cls.tag 		= idx
	cls.nameIdx 	= me.ReadShort()
	me.constanPool 	= append(me.constanPool,cls)
}


/**
常量池内容解析
 */
func (me *ClassByteFile) parserConstansPool() {

	x := int( me.ReadUi8() )

	switch x {
	case 3:

	case 4:

	case 5:

	case 6:

	case 7:
		me.parserConstanClassInfo(uint8(x))
	case 8:

	case 9:

	case 10:
		me.parserConstanMethod(uint8(x));
	default:
		fmt.Print( "常量池解析错误" );
		os.Exit(1);

	}

}

/**
解析class
 */
func (me *ClassByteFile) Parser() {

	data := me.Read(4);
	//能正常使用
	fmt.Println( 	binary.BigEndian.Uint32(data)==0xcafebabe )
	version 		:= me.ReadShort();
	masterVersion 	:= me.ReadShort();
	constanCount 	:= me.ReadShort();
	me.constanCount = int(constanCount);

	fmt.Print( fmt.Sprintf("主版本%d,副版本%d,常量池%d\n",masterVersion, version, constanCount) );
	//不知道这多两个字节是什么鬼.. 难道是常量池内容所占的字节数？
	d := me.ReadShort();


	me.parserConstansPool();

	fmt.Print(d);


}


func (me *ClassByteFile) get(n int){

}





