package main

import (
	"errors"
	"fmt"
)

type locationState struct {
	prev *string
	next *string
}

func newLState() *locationState {
	dft := "/location-area"
	ls := &locationState{
		prev: nil,
		next: &dft,
	}
	return ls
}

func llDo(cfg *config, endpoint string) error {
	ll, err := cfg.client.LocationList(endpoint)
	if err != nil {
		return err
	}
	cfg.ls.prev = &ll.Previous
	cfg.ls.next = &ll.Next

	for _, l := range ll.Results {
		fmt.Printf("%s\n", l.Name)
	}

	return nil
}

func commandMap(cfg *config, params ...string) error {
	return llDo(cfg, *cfg.ls.next)
}

func commandMapb(cfg *config, params ...string) error {
	if *cfg.ls.prev == "" {
		//lint:ignore ST1005 Being written to user
		return errors.New("You're on the first page!")
	}

	return llDo(cfg, *cfg.ls.prev)
}
