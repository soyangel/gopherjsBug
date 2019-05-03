package psjs

import (
	"testBug/psgo"

	"github.com/gopherjs/gopherjs/js"
)

type Msg struct {
	*js.Object
	To string `js:"to"`
}

var idCnt = 0
var subs = map[int]*psgo.Subscriber{}

func init() {
	ob := js.Global.Get("Object").New()
	js.Global.Set("psgo", ob)
	ob.Set("newSubscriber", newSubscriber)
	ob.Set("subscribe", subscribe)
	ob.Set("publish", publish)
}

func newSubscriber(f func(m *Msg)) int {
	wf := func(m *psgo.Msg) {
		f(nil)
		///
		// UNCOMMENTING THE FOLLOWING log.Println FIXES THE BUG!!!!!!
		///
		//log.Println("IT WORKS!")
	}
	idCnt++
	subs[idCnt] = psgo.NewSubscriber(wf)
	return idCnt
}

func subscribe(id int, paths ...string) {
	subs[id].Subscribe(paths...)
}

func publish(msg *Msg) int {
	m := &psgo.Msg{To: msg.To}
	go psgo.Publish(m)
	return 0
}
