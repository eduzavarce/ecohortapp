package repository

import (
	"errors"
	"time"
)

var(
	errUpdateField = errors.New("l'actualització ha fallat")
	errDeleteField = errors.New("el borrat ha fallat")
)

type Repository interface {
	Migrate() error
	InsertRegistre(h Registres)(*Registres, error)
	ObtenirTotsRegistres() ([]Registres, error)
	ObtenirRegistresPerID(id int64)(*Registres, error)
	ActualitzarRegistre(id int64, actualitzar Registres)error
	BorrarRegistre(id int64)error
}

type Registres struct {
	ID int64 `json:"id"`
	Data time.Time `json:"data_registre"`
	Precipitacio int `json:"precipitacio"`
	TempMaxima int `json:"temp_maxima"`
	TempMinima int `json:"temp_minima"`
	Humitat int `json:"humitat"`
}

