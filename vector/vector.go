package vector

type Vector interface {
	Scale(float64) Vector
	Magnitude() float64
	Distance(Vector) float64
	Add(Vector) Vector
}
