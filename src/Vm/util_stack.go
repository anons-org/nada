package Vm



type Stack struct {
	data []interface{}
}



func (me* Stack ) push( o interface{}) *Stack  {
	me.data = append(me.data,o);
	return me;
}

func (me* Stack ) pop() interface{}{

	return nil
}





func (me* Stack ) Size() int  {
	return len(me.data)
}

func (me* Stack ) Remove()  {

}

func (me* Stack ) Clean()  {

}






