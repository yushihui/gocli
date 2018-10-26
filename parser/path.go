package parser

type Path struct {
	Success bool
	Hops []Hop
}

type Hop struct {
	Name string
	North string
	South string
}