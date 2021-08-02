package Vm

import "fmt"

func printStreamPrint(args *FArray) IFObject {
	fmt.Print(args.get(0))
	return NONE
}


func doubleValueOf(args *FArray) IFObject {
	v:=args.get(0).(*FDouble).getVal()
	return NewFDouble(v)
}