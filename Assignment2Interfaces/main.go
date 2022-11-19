package main

type triangle struct {
	sideLength float64
	getArea
}

type square struct {
	hight float64
	base  float64
	getArea
}
type shape interface {getArea() float64}

func main() {

}

func (t triangle) getArea() float64
func (s square) getArea() float64
func getArea
