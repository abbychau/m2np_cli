package shell

import (
	"bufio"
	"fmt"
	"m2np_cli/ctx"
	"os"
)

type shell struct {
	in       *bufio.Reader
	ctx      *ctx.M2npContext
	position *directory
}

func New() *shell {
	return &shell{
		in:       bufio.NewReader(os.Stdin),
		ctx:      &ctx.M2npContext{},
		position: root,
	}
}

func (s *shell) Start() {
	for {
		fmt.Printf("[%s@m2np %s]# ", s.ctx.User.Username, s.position.name)
		inputs, err := s.readline()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		res, err := s.position.action(s, inputs)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Print(res)
	}
}

func (s *shell) readline() ([]byte, error) {
	inputs, err := s.in.ReadBytes('\n')
	return inputs[0 : len(inputs)-1], err
}
