package Vm


type Interpreter struct {

	//当前执行的栈帧
	frame *Frame;
	//虚拟机实例
	vms *Vms;
}



func (me *Interpreter)eval(){

}

/**
	设置栈帧
 */
func (me *Interpreter)setFrame(f *Frame){

	me.frame = f;
}


func (me *Interpreter)run(m FMethod){
	me.frame = 1
	//frame = new FrameObject(method);

	// 执行代码
	me.exec();
}



/**
	运行指令
 */
func (me *Interpreter)exec(){

	for me.frame.hasMoreOpcode(){

	}

}