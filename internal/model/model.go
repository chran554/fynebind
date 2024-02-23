package model

type SubData struct {
	Data string
}

type MyData struct {
	Name string
	Info string
	Sub  SubData
}
