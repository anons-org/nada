package CommonConstan

//OPCODE
var G_OPCODE_DEF = []struct {
	Idx int
	Id   int
	Name string
}{
	{1,1, "测试"},
	{2, 1,"测试"},
}

var G_OPCODE = []struct {
	Idx int
	Id   int
	Name string
}{
	{1,1, "测试"},
	{2, 1,"测试"},
}




//对象类型---------------------
var G_OT_DEF = []struct {
	Id   int
	Name string
}{
	{0, "int"},
	{1, "string"},
	{2, "bool"},
	{3, "function"},
	{4, "method"},
}

var G_OT = struct {
	INT   int
	STRING int
}{
	INT:1,STRING: 2,
}
