package math

type Match struct{
	A int
	B int
}

func (m Match) add() int{
	return m.A + m.B
}