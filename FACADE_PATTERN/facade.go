package main

import "fmt"

type Buffer struct {
	width, height int
	buffer []rune
}

type ViewPort struct{
	buffer *Buffer
	offset int
}

func NewViewPort(buffer *Buffer) *ViewPort {
	return &ViewPort{buffer: buffer}
}

func (v *ViewPort) GetCharacterAt(i int) rune {
	return v.buffer.At(v.offset + i)
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{width: width, height: height, buffer: make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune{
	return b.buffer[index]
}

type Console struct{
	buffers []*Buffer
	viewPorts []*ViewPort
	offset int
}

func NewConsole()*Console{
	b := NewBuffer(200, 150)
	v := NewViewPort(b)
	return &Console{[]*Buffer{b}, []*ViewPort{v}, 0}
}

func (c *Console) GetCharacterAt(i int)rune{
	return c.viewPorts[0].GetCharacterAt(i)
}

func main(){
	c := NewConsole()
	fmt.Println(c)
}
