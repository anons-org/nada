package Vm

import "fmt"

func NewFString(str string)  *FString {

	n := new(FString).Build().SetVal(str)
	return  n
}


func NewFFLoat(v float64)  *FFloat {
	n := new(FFloat).Build().setVal(v)
	return  n
}


func NewFInt(v int)  *FInt {
	n := new(FInt).Build().setVal(v)
	return  n
}

func NewFUint16(v uint16)  *FUint16 {
	n := new(FUint16).Build().setVal(v)
	return  n
}

func NewFMethod(t MethodType,name string)  *FMethod {
	n := new(FMethod).Build(t)
	n.mName=name
	return  n
}


func NewFMap()  *FMap {
	n := new(FMap).Build()
	return  n
}


func  NewFrame ( e IFObject) *Frame{

	xxx:=e.GetKlass()
	fmt.Print(xxx)
	me:=new(Frame)
	me.fn = e;
	me.code = e.(*FMethod).code;


	///**
	// * 		本地变量是有类型的，在实际测试中
	// * 		会发现本地变量有重复的槽号，但类型不同
	// */
	//localVar.put( OT.INT, 	new HashMap<Integer, IFObject>() );
	//localVar.put( OT.LONG, 	new HashMap<Integer, IFObject>() );
	//localVar.put( OT.REF, 	new HashMap<Integer, IFObject>() );

	return me
}






