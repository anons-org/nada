package ftype

type TypeKlass struct {
	//klass名称
	name string

}


func (me *TypeKlass)GetName() string{
	return me.name
}


func (me *TypeKlass)Build() *TypeKlass {
	me.name="TypeKlass"
	return me
}
