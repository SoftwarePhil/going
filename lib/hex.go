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
	a := MakeHex(3, 50)
	fmt.Println(a[0].Connections[0].connectionID)
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
	fmt.Println(strconv.Itoa(i) + "func start")
	for !dots[i].full && (count+i*2) < len(dots) {
		dots[i].addConnection(dots[i+count*2], distance, "c"+strconv.Itoa(i)+":"+strconv.Itoa(count))
		fmt.Println(strconv.Itoa(i + count))
		count++
	}
	return dots[i]
}

func (d *Dot) PrintDot() {
	s := ""
	for x := range d.Connections {
		s = s + "\n	Name: " + d.Connections[x].ConnectionDot + " Distance: " + strconv.Itoa(getConnection(d, d.Connections[x].connectionID).Distance) //how to find distance with map?
	}
	fmt.Println("\nDot id: " + d.DotID + "\nConnections: " + s)
}
