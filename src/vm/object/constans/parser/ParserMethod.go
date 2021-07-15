package ConstansParser

import (
	"vm/type"
)
/**
常量池对象 int
 */
type ParserMethod struct {
	val int
}


func (me *ParserMethod) Parser(idx int) _type.IFObject {
	return nil;
}


