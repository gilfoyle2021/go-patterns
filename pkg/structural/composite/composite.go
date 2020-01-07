package composite

import "fmt"

// 组合模式

//Swim swim interface
type Swimer interface {
	Swim()
}

//Train train interface
type Trainer interface {
	Train()
}

type Swimimpl struct {
}

func (sw *Swimimpl) Swim() {
	fmt.Println("Swimming...")
}

type Trainimpl struct {
}

func (sw *Trainimpl) Train() {
	fmt.Println("Training...")
}

type Athelete struct {
	Swimer
	Trainer
}

func (ath *Athelete) Exec() {

}
