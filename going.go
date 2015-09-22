package main
import (
"fmt"
"time"
"math/rand"
"sync"
"strconv"
)
func main() {
	emptyMap := map[int]Point{}
	var wg sync.WaitGroup
	a := Worker{Point{1,1},"a", emptyMap}
	b := Worker{Point{5,5},"b", emptyMap}
	wg.Add(2)
	go a.move(&wg)
	go b.move(&wg)
	wg.Wait()
	fmt.Println(a.constructPath())
}

type Worker struct {
	Point
	name string
	m map[int]Point
	}

type Point struct {
	x int
	y int
	}
	

func (w *Worker) move(wg *sync.WaitGroup) {
	//for i := 0; i < 100; i++{
	defer wg.Done()
	for i := 0; i < 50; i++{
	w.x = w.x + selectValue()
	w.y = w.y + selectValue()
	place := ("x : " + strconv.Itoa(w.x) +" "+ "y : " + strconv.Itoa(w.y) + " " + w.name + " " + time.Now().String())
	w.addPlace(i)
	fmt.Println(place)
	time.Sleep(500 * time.Millisecond)
	}
}

func (w*Worker) addPlace(i int){
		w.m[i] = w.Point
	}
	
func (w*Worker) constructPath() string{
	places := ""
	for i := 0; i < 50; i++{
		places = places + " " +"(" +strconv.Itoa(w.m[i].x) + " " + strconv.Itoa(w.m[i].y) + ")"
		}
		
		return places
	}

func selectValue() int{
		if (rand.Intn(100) > 50) {
			return rand.Intn(2)*(-1)
		}
		
		return rand.Intn(2)
} 
