package main

import (
	"fmt"
	"strconv"
)

type ConnectionMap map[*Dot]*Connection

type Dot struct {
	DotID       string
	Connections []ConnectionInfo
	M           ConnectionMap
	full        bool
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
	a := MakeHexShape(3)
	for i := range a {
		a[i].PrintDot()
	}
}

func createEmptyDot(s string) *Dot {
	dot := new(Dot)
	dot.DotID = s
	dot.M = make(ConnectionMap)
	dot.full = false
	return dot
}

func (d *Dot) addConnection(other *Dot, distance int, id string) {

	if _, ok := d.M[other]; !ok {
		if len(d.Connections) < 6 {
			a := [2]*Dot{d, other}
			c := &Connection{a, distance, id}
			d.M[other] = c
			other.M[d] = c

			d.Connections = append(d.Connections, ConnectionInfo{id, other.DotID})
			other.Connections = append(other.Connections, ConnectionInfo{id, d.DotID})
		} else {
			d.full = true
			fmt.Println("no room!")
		}
	} else {
		fmt.Println("Already connected")
	}
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

func getDot(d *Dot, id string) *Dot {
	for value := range d.M {
		if value.DotID == id {
			return value
		}
	}
	a := new(Dot)
	a.DotID = "nil"
	return a
}

func MakeHexShape(distance int) []*Dot {
	dots := make([]*Dot, 7)

	for count := range dots {
		dots[count] = createEmptyDot("dot" + strconv.Itoa(count))
	}
	dots[0].addConnection(dots[1], distance, "0c1")
	dots[0].addConnection(dots[2], distance, "0c2")
	dots[0].addConnection(dots[3], distance, "0c3")
	dots[0].addConnection(dots[4], distance, "0c4")
	dots[0].addConnection(dots[5], distance, "0c5")
	dots[0].addConnection(dots[6], distance, "0c6")

	dots[1].addConnection(dots[2], distance, "1c6")
	dots[1].addConnection(dots[6], distance, "1c6")

	dots[2].addConnection(dots[3], distance, "2c6")

	dots[3].addConnection(dots[4], distance, "3c6")

	dots[4].addConnection(dots[5], distance, "4c6")

	dots[5].addConnection(dots[6], distance, "5c6")
	return dots
}

func (d *Dot) PrintDot() {
	s := ""
	for x := range d.Connections {
		s = s + "\n	Name: " + d.Connections[x].ConnectionDot + " Distance: " + strconv.Itoa(getConnection(d, d.Connections[x].connectionID).Distance) //how to find distance with map?
	}
	fmt.Println("\nDot id: " + d.DotID + "\nConnections: " + s)
}

/*
func MakeHex(distance, size int) []*Dot {
	dots := make([]*Dot, size)
	for count := range dots {
		dots[count] = createEmptyDot("dot" + strconv.Itoa(count))
	}
	for count := range dots {
		dots[count] = fillDot(count, distance, dots)
	}
	return dots
}

func fillDot(i, distance int, dots []*Dot) *Dot {
	count := 1
	for !dots[i].full && (count+i) < len(dots) {
		dots[i].addConnection(dots[i+count], distance, "c"+strconv.Itoa(i)+":"+strconv.Itoa(count))
		fmt.Println(strconv.Itoa(i + count))
		count++
	}
	return dots[i]
}
*/
