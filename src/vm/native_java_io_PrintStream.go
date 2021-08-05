package vm

import (
	"fmt"
)

func printStreamPrint(args *FArray) IFObject {
	fmt.Println(args.get(0))
	return NONE
}


func doubleValueOf(args *FArray) IFObject {
	v:=args.get(0).(*FDouble).getVal()
	return NewFDouble(v)
}