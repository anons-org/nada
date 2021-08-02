package Vm

func NewFString(str string)  *FString {

	n := new(FString).Build().SetVal(str)
	return  n
}


func NewFFLoat(v float64)  *FFloat {
	n := new(FFloat).Build().setVal(v)
	return  n
}


func NewFDouble(v float64)  *FDouble {
	n := new(FDouble).Build().setVal(v)
	return  n
}


func NewFByte(v []byte)  *FByte {
	n := new(FByte).Build().setVal(v)
	return  n
}



func NewFInt(v int)  *FInt {
	n := new(FInt).Build().setVal(v)
	return  n
}


func NewFArray()  *FArray {
	n := new(FArray).Build()
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


func NewNativeFMethod(name string, m func(arg *FArray) IFObject ) *FMethod {
	n := new(FMethod).Build(METHOD_TYPE_NATIVE)
	n.mName = name
	n.call = m
	return  n;
}



func NewFMap()  *FMap {
	n := new(FMap).Build()
	return  n
}

func NewFObject()  IFObject {
	n := new(FObject).Build()
	return  n
}


func NewFStack()  *FStack {
	n := new(FStack).Build()
	return  n
}



func  NewFrame ( e IFObject) *Frame{

	me:=new(Frame)
	me.pc = -1;
	me.fn = e;
	me.localVal = NewFMap()
	me.statck = NewFStack()
	me.code = e.(*FMethod).code;
	me.codes = e.(*FMethod).codes;


	///**
	// * 		本地变量是有类型的，在实际测试中
	// * 		会发现本地变量有重复的槽号，但类型不同
	// */
	//localVar.put( OT.INT, 	new HashMap<Integer, IFObject>() );
	//localVar.put( OT.LONG, 	new HashMap<Integer, IFObject>() );
	//localVar.put( OT.REF, 	new HashMap<Integer, IFObject>() );

	return me
}






