package main

import (
	"fmt"
	"strings"
)

type proxyPresident struct {
	realPresident President
}

func NewProxyPresident(realPresident President) President {
	return &proxyPresident{realPresident: realPresident}
}

func (p *proxyPresident) ListenTrouble(trouble string) {

	fmt.Println("We will deal with it immediately...")

	important := extract(trouble)
	if important {
		p.realPresident.ListenTrouble(trouble)
	}

}

func (p *proxyPresident) FindJob(s string) {

	fmt.Println("We will try to help...")

}

func extract(trouble string) bool {

	if strings.Contains(trouble, "please") {
		return true
	}

	return false
}
