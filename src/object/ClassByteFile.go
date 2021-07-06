package object

import (
	"../etc"
	"../vm/object/constans"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)



type ClassByteFile struct {

	//当前JAVA字节码读取位置
	readIdx int
	//文件的名字
	fileName string
	//CLASS 字节码
	data []byte
	constanPool []Constans.ConstansInterface
}

func (me *ClassByteFile) Load(fileName string){


	fmt.Print(  CommonConstan.G_OT_DEF[CommonConstan.G_OT.STRING].Name);


//	var fileName = "E:\\AAAA_CODE\\new-eclipse-workspace\\far-dev\\demo\\TestDemo1.class"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666);
	if err != nil {
		fmt.Println("Open file error!", err);
		return
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
				return
			}
		}
		//data = append(data,buf);
		//fmt.Println(length, string(buf));
	}

	me.data = buf;
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


/**
读取常量池对象
 */
func (me *ClassByteFile) parserConstanInteger() {
	integer := new(Constans.ConstanInteger)
	me.constanPool = append(me.constanPool,integer)
}




func (me *ClassByteFile) get(n int){

}





