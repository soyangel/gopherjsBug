package psgo

import (
	"fmt"
	"sync"
)

type Msg struct {
	To string
}

type Subscriber struct {
	subs map[string]bool
	f    func(msg *Msg)
}

var oldMessages = map[string]*Msg{}
var psLock sync.Mutex

func NewSubscriber(f func(msg *Msg)) *Subscriber {
	return &Subscriber{subs: map[string]bool{}, f: f}
}

func (su *Subscriber) Subscribe(paths ...string) {
	//
	// COMMENTING THE FOLLOWING 2 LINES FIXES THE BUG!!!!!!
	psLock.Lock()
	defer psLock.Unlock()

	for _, path := range paths {
		msg := oldMessages[path]
		if msg != nil {
			fmt.Println("BEFORE EXEC GOROUTINE")
			go su.f(msg)
			fmt.Println("AFTER EXEC GOROUTINE")
		}
	}
}

func Publish(msg *Msg) {
	oldMessages[msg.To] = msg
}
