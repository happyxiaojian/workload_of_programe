package goFunTest

import (
	"fmt"
	"testing"
)

type PlayState interface {
	Play()
}

type Robot struct{}

func (r *Robot)Play(){
	fmt.Println("Robot play")
}

type User struct{}

func (u *User)Play(){
	fmt.Println("Player play")
}

// Flag  0-robot 1-user
type Player struct{
	State PlayState
	Flag int
}

func NewPlayer(flag int)*Player{
	p := &Player{
		Flag:  flag,
	}
	p.changeState()
	return p
}

func (p *Player) changeState() {
	if p.Flag == 1{
		p.State = &User{}
	}else {
		p.State = &Robot{}
	}
}

func (p *Player)SetFlag(flag int){
	p.Flag = flag
	p.changeState()
}


func TestFunc(t *testing.T){

	p := NewPlayer(0)
	p.State.Play()

	p.SetFlag(1)
	p.State.Play()

	p.SetFlag(0)
	p.State.Play()
}

//================================================


