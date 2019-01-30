package main

import "fmt"

type Point struct{
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

func (p *Point) Move(dx int, dy int){
	p.X += dx
	p.Y += dy
}

func (s *Square) Move(dx int, dy int){
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int{
	return s.Length * s.Length
}

func NewSquare(X int, Y int, Length int) (*Square, error) {
	if Length <= 0{
		return nil, fmt.Errorf("length must be greater than 0")
	}


	//this was a dumb way to do it...
	/* try:
	square := &Square{
		Center: Point{X, Y},
		Length: Length,
	}

	*/
	 square := new(Square)
	point := new(Point)
	point.X = X
	point.Y = Y
	square.Center = *point
	square.Length = Length

	return square, nil

}

func main(){
	square, err := NewSquare(2,2, 2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(square)
	square.Move(1,1)
	fmt.Println(square)
	area := square.Area()
	fmt.Println(area)
}