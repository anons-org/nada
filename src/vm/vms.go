package Vm


type Vms struct {

	//所有的Klass元类
	metaKlass *FMap

	//项目的class目录 默认当前的lib目录
	classPath string

	//所有的解释器

}

func (me *Vms ) start()  {


	//初始化参数 klass等


	var fileName = "E:\\AAAA_CODE\\new-eclipse-workspace\\far-dev\\demo\\lib\\Test.class"
	classN  := NewClassByteFile()
	classN.Load(fileName).Parser();
	main:=me.metaKlass.get("nada.lib.TestKlass").(IKlass)
	f:=main.getStaticDict("main").(*FMethod)

	interp:=new(Interpreter)
	interp.run(f)

}




func (me *Vms ) Build() *Vms{

	maps:=make(map[string]IKlass)
	maps["BUILTN.NADA_TYPE"] 	= new(TypeKlass).Init()

	me.metaKlass = NewFMap()
	me.metaKlass.set(BUILTN.NADA_TYPE,   	maps["BUILTN.NADA_TYPE"])
	me.metaKlass.set(BUILTN.NADA_NATIVE, 	new(NativeMethodKlass).Init())
	me.metaKlass.set(BUILTN.NADA_VIRTUAL,	new(MethodKlass).Init())
	me.metaKlass.set(BUILTN.NADA_STRING,    new(StringKlass).Init())
	me.metaKlass.set(BUILTN.NADA_FLOAT,    new(FloatKlass).Init())



	return me
}



