package main

import "github.com/Bakarseck/file"



func main()  {
	a, err := file.Read("../a.txt", " ")
	file.Check(err)

	file.Write("../b.txt", a, ",")
}