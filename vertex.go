package graph

import "fmt"

func NewVertex(data interface{}) *Vertex {
	return &Vertex{
		data:        data,
		color:       White,
		predecessor: nil,
		meta:        nil,
	}
}

func (v *Vertex) SetColor(c Color) bool {
	if v == nil {
		return false
	}

	v.color = c
	return true
}

func (v *Vertex) Color() Color {
	if v == nil {
		return White
	}

	return v.color
}

func (v *Vertex) Data() interface{} {
	if v == nil {
		return nil
	}
	return v.data
}

func (v *Vertex) String() string {
	if v == nil {
		return "nil"
	}
	return fmt.Sprintf("[%v<=%v]::%v", v.data, v.predecessor, v.color.String())
}
