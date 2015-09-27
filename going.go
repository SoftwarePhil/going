package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	a := worker{point{1, 1, time.Now().String()}, "a", map[string]point{}}
	b := worker{point{5, 5, time.Now().String()}, "b", map[string]point{}}
	wg.Add(2)
	go a.move(&wg)
	go b.move(&wg)
	wg.Wait()

	fmt.Println(a.constructPath())
	fmt.Println(b.constructPath())
}

type worker struct {
	point
	name string
	m    map[string]point
}

type point struct {
	x     int
	y     int
	aTime string
}

func (w *worker) move(wg *sync.WaitGroup) {
	//for i := 0; i < 100; i++{
	defer wg.Done()
	for i := 0; i < 6; i++ {
		w.x = w.x + selectValue()
		w.y = w.y + selectValue()
		if _, ok := w.m[strconv.Itoa(w.x)+" "+strconv.Itoa(w.y)]; ok {
			w.x = w.x + selectValue()
			w.y = w.y + selectValue()
		}
		aTime := time.Now().String()
		strPlace := ("x : " + strconv.Itoa(w.x) + " " + "y : " + strconv.Itoa(w.y) + " " + w.name + " " + aTime)
		place := strconv.Itoa(w.x) + " " + strconv.Itoa(w.y)
		w.addPlace(place, aTime)
		fmt.Println(strPlace)
		time.Sleep(500 * time.Millisecond)
	}
}

func (w *worker) addPlace(s, t string) {
	w.aTime = t
	w.m[s] = w.point
}

func (w *point) getLocation() string {
	return "(" + strconv.Itoa(w.x) + "," + strconv.Itoa(w.y) + ")"
}

func (w *worker) constructPath() string {
	s := " "
	f := s
	for key, value := range w.m {
		f = key
		s = s + value.getLocation()
	}
	f = ""
	return s + f
}

func selectValue() int {
	if rand.Intn(100) > 50 {
		return rand.Intn(2) * (-1)
	}

	return rand.Intn(2)
}
