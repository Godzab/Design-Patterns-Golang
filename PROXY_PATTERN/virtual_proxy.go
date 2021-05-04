package main

import "fmt"

type Image interface {
	Draw()
}


type BitMap struct{
	fileName string
}

func NewBitMap(fileName string) *BitMap {
	fmt.Println("Loading image from", fileName)
	return &BitMap{fileName: fileName}
}

type LazyBitMap struct{
	fileName string
	bitMap *BitMap
}

func (lb *LazyBitMap) Draw(){
	if lb.bitMap == nil{
		lb.bitMap = NewBitMap(lb.fileName)
	}
	lb.bitMap.Draw()
}

func NewLazyBitMap(fileName string) *LazyBitMap {
	return &LazyBitMap{fileName: fileName}
}

func (b *BitMap) Draw(){
	fmt.Println("We are drawing", b.fileName)
}

func DrawImage(image Image){
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

func main(){
	bmp := NewLazyBitMap("Demo.png")
	DrawImage(bmp)
}