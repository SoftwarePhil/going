package moving

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)


func Init() *sync.WaitGroup {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	return &wg
}

type Worker struct {
	Point
	Name string
	M    map[string]Point
}

type Point struct {
	X     int
	Y     int
	ATime string
}

func (w *Worker) StartMoving (wg *sync.WaitGroup){
	wg.Add(1)
	go w.move(wg)
	wg.Wait()
}

func (w *Worker) move(wg *sync.WaitGroup) {
	//for i := 0; i < 100; i++{
	defer wg.Done()
	for i := 0; i < 6; i++ {
		
		if(!w.checkIfVisited()){
			w.X = w.X + selectValue()
			w.Y = w.Y + selectValue()
			w.checkIfVisited()
		}
		aTime := time.Now().String()
		strPlace := ("x : " + strconv.Itoa(w.X) + " " + "y : " + strconv.Itoa(w.Y) + " " + w.Name + " " + aTime)
		place := strconv.Itoa(w.X) + " " + strconv.Itoa(w.Y)
		w.addPlace(place, aTime)
		fmt.Println(strPlace)
		time.Sleep(500 * time.Millisecond)
	}
}

func (w * Worker) checkIfVisited() bool{
	if _, ok := w.M[strconv.Itoa(w.X)+" "+strconv.Itoa(w.Y)]; ok {
			return false
		}
		return true
}
func (w *Worker) addPlace(s, t string) {
	w.ATime = t
	w.M[s] = w.Point
}

func (w *Point) getLocation() string {
	return "(" + strconv.Itoa(w.X) + "," + strconv.Itoa(w.Y) + ")"
}

func (w *Worker) ConstructPath() string {
	s := " "
	f := s
	for key, value := range w.M {
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
/*
var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	wg.Add(2)
	go a.move(&wg)
	go b.move(&wg)
	wg.Wait()

	fmt.Println(a.constructPath())
	fmt.Println(b.constructPath())
*/
