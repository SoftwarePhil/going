package main
import (
"fmt"
"time"
"math/rand"
"sync"
)
func main() {
	var wg sync.WaitGroup
	a := Worker{1,1,"a"}
	b := Worker{5,5,"b"}
	wg.Add(2)
	go a.move(&wg)
	go b.move(&wg)
	wg.Wait()
}

type Worker struct {
	x int
	y int
	name string
	}
	

func (w *Worker) move(wg *sync.WaitGroup) {
	//for i := 0; i < 100; i++{
	defer wg.Done()
	for i := 0; i < 50; i++{
	w.x = w.x + selectValue()
	w.y = w.y + selectValue()
	fmt.Println("x : ", w.x, "y : ", w.y, w.name," ", time.Now())
	time.Sleep(500 * time.Millisecond)
	}
}

func selectValue() int{
		if (rand.Intn(100) > 50) {
			return rand.Intn(2)*(-1)
		}
		
		return rand.Intn(2)
} 
