package main

import "going/lib"
import "fmt"
import "time"

func main(){
	a := moving.Worker{moving.Point{1, 1, time.Now().String()}, "a", map[string]moving.Point{}}
	b := moving.Worker{moving.Point{1, 1, time.Now().String()}, "b", map[string]moving.Point{}}
	c := moving.Worker{moving.Point{1, 1, time.Now().String()}, "c", map[string]moving.Point{}}
	wg := moving.Init()
	go a.StartMoving(wg)
	go b.StartMoving(wg)
	go c.StartMoving(wg)
	time.Sleep(5000 * time.Millisecond)
	fmt.Print(a.ConstructPath())
}
