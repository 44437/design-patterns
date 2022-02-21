package main

import "fmt"

type realPresident struct {
}

func NewRealPresident() President {
	return &realPresident{}
}

func (r *realPresident) ListenTrouble(trouble string) {
	fmt.Printf("'%s' your trouble is my trouble.\n", trouble)
}

func (r *realPresident) FindJob(s string) {
	fmt.Println("I can not do this")
}
