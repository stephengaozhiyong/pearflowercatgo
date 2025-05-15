package liquidunit

import "fmt"

type Gallons float64

func (g *Gallons) ToLitters() Litters {
	return Litters(*g * 3.785)
}

func (g *Gallons) Double() {
	*g *= 2
}

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f Gallons.", g)
}

type Litters float64

func (l *Litters) ToGallons() Gallons {
	return Gallons(*l / 3.785)
}

func (l Litters) String() string {
	return fmt.Sprintf("%0.2f Litters.", l)
}

type MilliLitters float64

func (m *MilliLitters) ToGallons() Gallons {
	return Gallons(*m / 3.785 / 1000)
}

func (m MilliLitters) String() string {
	return fmt.Sprintf("%0.2f MilliLitters.", m)
}
