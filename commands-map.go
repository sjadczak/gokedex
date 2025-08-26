package main

import (
	"fmt"

	"github.com/sjadczak/gokedex/internal/pokeapi"
)

type locationState struct {
	page int
}

func newLState() *locationState {
	ls := &locationState{
		page: 0,
	}
	return ls
}

func (ls *locationState) inc() {
	ls.page++
}

func (ls *locationState) dec() {
	ls.page -= 2
	if ls.page < 0 {
		ls.page = 0
	}
}

func makeMapCommands(client *pokeapi.Client) (func() error, func() error) {
	ls := newLState()

	cm := func() error {
		s := ls.page*20 + 1
		e := s + 20
		for i := s; i < e; i++ {
			l, err := client.LocationArea(i)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", l.Name)
		}
		ls.inc()
		return nil
	}

	cmb := func() error {
		ls.dec()
		s := ls.page*20 + 1
		e := s + 20
		for i := s; i < e; i++ {
			l, err := client.LocationArea(i)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", l.Name)
		}
		return nil
	}

	return cm, cmb
}
