package main

import "fmt"

type ConnectionMap map[Connection]*Dot

type Dot struct {
	DotID       string
	Connections []string
	M           ConnectionMap
}

type Connection struct {
	Joint    [2]*Dot
	Distance int
	JointID  string
}

func main() {
	aDot := createEmptyDot("my dot")
	aDot2 := createEmptyDot("my dot2")

	addConnection(aDot, aDot2, 5, "c1")
	aDot.getConnection("c1")
	fmt.Printf("%d", aDot.getConnection("c1").Distance)
}

func createEmptyDot(s string) *Dot {
	dot := new(Dot)
	dot.DotID = s
	dot.M = make(ConnectionMap)
	return dot
}

func addConnection(d, other *Dot, distance int, id string) {

	a := [2]*Dot{d, other}
	c := Connection{a, distance, id}
	d.M[c] = other
	other.M[c] = d

	d.Connections = append(d.Connections, other.DotID)
	other.Connections = append(other.Connections, d.DotID)
}

func (d *Dot) getConnection(id string) *Connection {
	for key := range d.M {
		if key.JointID == id {
			return &key
		}
	}
	return nil
}

func (d *Dot) PrintDot() {
	fmt.Print(d.DotID)
}
