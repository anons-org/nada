package Vm

func NewFString(str string)  *FString {

	n := new(FString).Build().SetVal(str)
	return  n
}

func NewFInt()  {

}

func NewFMethod(t int,name string)  *FMethod {
	n := new(FMethod).Build(t)
	return  n
}





