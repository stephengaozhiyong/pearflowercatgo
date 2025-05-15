

import "fmt"

type Gallons float64

func (g Gallons) ToLitters() Litters {
	return Litters(g * 3.785)
}

func (g *Gallons) Double() {
	*g *= 2
}

type Litters float64

func (l Litters) ToGallons() Gallons {
	return Gallons(l / 3.785)
}

type MilliLitters float64

func (m MilliLitters) ToGallons() Gallons {
	return Gallons(m / 3.785 / 1000)
}

func main() {

	carFule := Gallons(10)
	carFule.Double()
	fmt.Println(carFule)
}
