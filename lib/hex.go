package main

import (
	"fmt"
	"strconv"
)

type ConnectionMap map[*Dot]*Connection

type Dot struct {
	DotID       string
	Connections []ConnectionInfo
	Connection
	M ConnectionMap
}

type Connection struct {
	Joint    [2]*Dot
	Distance int
	JointID  string
}

type ConnectionInfo struct {
	connectionID  string
	ConnectionDot string
}

func main() {
	aDot := createEmptyDot("myDot")
	aDot2 := createEmptyDot("myDot2")
	aDot3 := createEmptyDot("myDot3")

	addConnection(aDot, aDot2, 5, "c1")
	addConnection(aDot3, aDot, 2, "c2")
	addConnection(aDot2, aDot3, 3, "c3")

	fmt.Printf("%d", getConnection(aDot, aDot.Connections[0].connectionID).Distance)

	aDot.PrintDot()
	aDot2.PrintDot()
	aDot3.PrintDot()
}

func createEmptyDot(s string) *Dot {
	dot := new(Dot)
	dot.DotID = s
	dot.M = make(ConnectionMap)
	return dot
}

func addConnection(d, other *Dot, distance int, id string) {

	a := [2]*Dot{d, other}
	c := &Connection{a, distance, id}
	d.M[other] = c
	other.M[d] = c

	d.Connections = append(d.Connections, ConnectionInfo{id, other.DotID})
	other.Connections = append(other.Connections, ConnectionInfo{id, d.DotID})
}

func getConnection(d *Dot, id string) *Connection {
	for _, value := range d.M {
		if value.JointID == id {
			return value
		}
	}
	a := new(Connection)
	a.JointID = "nil"
	return a
}

func (d *Dot) PrintDot() {
	s := ""
	for x := range d.Connections {
		s = s + "\n	Name: " + d.Connections[x].ConnectionDot + " Distance: " + strconv.Itoa(getConnection(d, d.Connections[x].connectionID).Distance)
	}
	fmt.Println("\nDot id: " + d.DotID + "\nConnections: " + s)
}
