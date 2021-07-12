package ConstansParserInter

type IConstanParser interface {
	Parser(idx int) interface{}
}