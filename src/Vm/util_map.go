package Vm

import (
	"container/list"
	"fmt"
)
/**
	这个map，实现潦草，性能一般，本着“先有”的战略定位，先临时用着
	temporary implementation MAP
 */
type MapEntry struct {
	key interface{}
	val interface{}
}

type Map struct {
	data *list.List
}

func toStr(e interface{}) string{
	return fmt.Sprintf("%v", e)
}

func (me* Map ) Set(k interface{}, v interface{})  {
	e:=new(MapEntry)
	e.key = k
	e.val = v
	me.data.PushBack(e);
}

func (me* Map ) Get(k interface{})  interface{}{

	//fmt.Print(toStr(k))

	for e:=me.data.Front();e!=nil; e= e.Next(){
		if ele , ok := (e.Value).(*MapEntry); ok==true{
			if toStr(ele.key) == toStr(k){
				return ele.val
			}
		}
	}
	return nil
}

func (me* Map ) Filter()  {

}

func (me* Map ) ForEach(  call func( entry *MapEntry )   )  {
	for e:=me.data.Front();e!=nil; e= e.Next(){
		if ele , ok := (e.Value).(*MapEntry); ok==true{
			call(ele)
		}
	}
}

func (me* Map ) Size() int  {
	return me.data.Len()
}

func (me* Map ) Remove()  {

}

func (me* Map ) Clean()  {

}

func (me* Map ) has( call func( interface{} ) bool  )  {

}


func NewMap() *Map{
	e:=new(Map)
	e.data = list.New()
	return e
}


