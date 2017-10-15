package model

import "fmt"

//import "time"

type Tag struct{
	Id string
	Name string
	//Created time.Time
	//Changed time.Time
}

func (Tag) ToString() {
	fmt.Println("Xablau")
}