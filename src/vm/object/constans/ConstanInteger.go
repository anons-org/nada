package Constans

/**
常量池对象 int
 */
type ConstanInteger struct {
	val int
}


func (me *ConstanInteger) Get() interface{} {
	return nil
}



func (me *ConstanInteger) SetVal(v int)  {
	me.val = v;
}
