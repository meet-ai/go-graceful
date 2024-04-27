// init.go
package main

import (
	"time"

	"github.com/changsongl/gevent"
)

var ge gevent.GEvent


// user.go
func main() {
    ge = gevent.NewGEvent()
    ge.AddObserver("account-create", "user", func(a int, b int) {
		println(a+b)
       // do something for user
    }, true)
    ge.TriggerEvent("account-create", 1, 2)
	time.Sleep(2*time.Second)
}