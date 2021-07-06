package object

import "fmt"

type Nation struct {
	Name string
	//友善度
	FriendOpt int
	//等级
	Lev int
	//盟国
	Ally []*Nation

	//工业
	//农业
	//教育
	//军事
}

func (this *Nation) AddAlly(al *Nation) {
	this.Ally = append(this.Ally, al)

	if al.Ally == nil {
		al.AddAlly(this)
	}

	//al.AddAlly(this)
}

func (this *Nation) Init() {
	this.Name = "x"
	this.FriendOpt = 1
}

func (this *Nation) Print() {
	fmt.Println("输出属性为")
	fmt.Println(this.Name)
	fmt.Println(this.FriendOpt)
}

func (this *Nation) act() {
	fmt.Println("输出属性为")
	fmt.Println(this.Name)
	fmt.Println(this.FriendOpt)
}

func Test() {
	fmt.Println("输出完成")
}
