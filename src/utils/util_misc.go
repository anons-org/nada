package Utils

import (
	"fmt"
	"net"
)

var miscUtilIns *Misc = nil


type Misc struct {
	ins *Misc
}

type MethodArgAndTypeToken struct {
	idx int
	size int
	content []rune

	tokens []*_methodArgType
}

/**
	java 方法的参数和返回值的类型
 */
type _methodArgType struct {
	t int  //1是基本类型 2 是引用类型
	v string
}


func (me* MethodArgAndTypeToken) nextChar() rune {
	i:=me.idx;
	me.idx++;
	if( me.idx > len(me.content)) {
		return 0
	}
	return me.content[i]
}






func testTcp() {

	var(
		lists []*net.TCPConn
		myconn *net.TCPConn
		err error
	)
	service := ":2000"                                 //设置服务器监听的端口
	tcpAddr, _ := net.ResolveTCPAddr("tcp", service) //创建 tcpAddr数据
	fmt.Println("tcpAddr  ")
	mylistener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		myconn, err = mylistener.AcceptTCP()
		fmt.Println(err)
		myconn.SetReadBuffer(1)

		var b []byte = make([]byte, 1)
		myconn.Read(b)
		lists = append(lists,myconn)
	}
}



func (me* MethodArgAndTypeToken) parser(s string) {
	//(DLjava/lang/String;)D
	//str:=strings.Split(s,")");
	//
	//str0:= strings.ReplaceAll(str[0],"(","")
	//retArg:= strings.ReplaceAll(str[1],")","")
	//fmt.Println(retArg)

	me.content =[]rune(s)
	for{
		 if char := me.nextChar();char!=0{
			 if string(char) =="L"{
				 var oks string
				 for{
					 if c := me.nextChar(); string(c)!=";" && c!=0{
						 oks+=string(c);
					 }else{

						 val:=&_methodArgType{
							 t:2,
							 v:oks,
						 }
						 me.tokens = append(me.tokens,val);
						 break
					 }
				 }
			 }else if string(char) =="V"{
			 	continue
			 }else{
				 val:=&_methodArgType{
					 t:1,
					 v:string(char),
				 }
				 me.tokens = append(me.tokens, val);
			 }
		 }else{ break}
	}

	fmt.Println(fmt.Sprintf("method arg type %s",me.tokens))

}


func (me *Misc) GetDescStrData(d string)   *MethodArgAndTypeToken {
	tk:=&MethodArgAndTypeToken{}
	tk.parser(d);
	tk.size = len(tk.tokens)
	return tk
}


func (me *MethodArgAndTypeToken) GetSize()   int {
	return me.size
}

func (me *MethodArgAndTypeToken) GetData()   []*_methodArgType {
	return me.tokens
}





func MiscIns()  *Misc {
	if miscUtilIns ==nil{
		miscUtilIns = new(Misc)
		return miscUtilIns;
	}
	return miscUtilIns
}
