package storage

import "github.com/Dev-TempusFugit/API/"

type Memory struct {
	CurrentID int
	Personas map[int]model.Person
}
