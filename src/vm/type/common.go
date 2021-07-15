package ftype

func NewFString()  *FString{
	n := new(FString).Build()
	return  n
}

func NewFInt()  {

}

func NewFMethod(t int)  *FMethod{
	n := new(FMethod).Build(t)
	return  n
}




