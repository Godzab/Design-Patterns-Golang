package main

type Sized interface {
	GetWidth() int
	SetWidth() int
	GetHeight() int
	SetHeight() int
}

type Rectangle struct {
	width int
	height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func main(){

}