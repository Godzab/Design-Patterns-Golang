package main

import "fmt"

var entryCount = 0

type Journal struct{
	entries []string
}

func (j Journal) AddEntry(txt string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, txt)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j Journal) RemoveEntry(index int) {

}

// SaveToFile  This BREAKS the single responsibility pattern
func (j Journal) SaveToFile(filename string) {

}
func (j Journal) Load(filename string) {

}

// SaveToFile  This PROMOTES the single responsibility pattern
func SaveToFile(j *Journal, filename string) {

}
