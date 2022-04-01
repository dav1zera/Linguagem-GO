package data

import (
	"errors"
)

//Structs

type Date struct {
	year  int
	month int
	day   int
}

//Getters

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

//Setters

func (d *Date) SetYear(year int) {
	d.year = year
}

//Tratamento de erros

func (d *Date) SetMonth(month int) error {
	if month > 12 || month < 1 {
		return errors.New("Mes inválido")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) {
	d.day = day
}
