package domain

type Vector struct {
	X float64
	Y float64
}

func AddVectors(vectors ...Vector) Vector {
	sum := Vector{}

	for _, v := range vectors {
		sum.X += v.X
		sum.Y += v.Y
	}

	return sum
}
