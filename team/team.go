package team

import (
	"fmt"

	"github.com/pattack/rcss"
)

type janstun struct {
	name  string
	side  rcss.Side
	match rcss.Match
}

func (j janstun) Name() string {
	return j.name
}

func (j *janstun) Init(match rcss.Match, side rcss.Side, unum rcss.UniformNumber, mode rcss.PlayMode) {
	j.side = side
	j.match = match

	fmt.Printf("Side: %c, Uniform Number: %d, Mode: %s\n", side, unum, mode)
}

func (j *janstun) ServerParam(sp rcss.ServerParameters) {
	fmt.Printf("Server Parameters: %#v\n", sp)
}

func (j *janstun) PlayerParam(pp rcss.PlayerParameters) {
	fmt.Printf("Player Parameters: %#v\n", pp)
}

func (j *janstun) PlayerType(pt rcss.PlayerType) {
	fmt.Printf("Player Type: %#v\n", pt)
}

func (j *janstun) See() {

}

func (j *janstun) Hear() {

}

func (j *janstun) SenseBody() {

}

func (j *janstun) Score() {

}

func NewTeam(name string) rcss.Team {
	return &janstun{
		name: name,
	}
}
