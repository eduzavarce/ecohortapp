package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLliteRepository struct{
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLliteRepository{
	return &SQLliteRepository{
		Conn: db,
	}
}
func(repo *SQLliteRepository) Migrate() error{
	query :=`
	create table if not exists registres(
		id integer primary key autoincrement,
		data_registre integer not null,
		precipitacio integer not null,
		temp_maxima integer not null,
		temp_minima integer not null,
		humitat integer not null
	)
	`
	_,err := repo.Conn.Exec(query)
	return err
}
func (repo *SQLliteRepository) InsertRegistre(regisrtres Registres) (*Registres, error){
	peticio := "insert into registres (data_registre, precipitacio, temp_maxima, temp_minima, humitat) values (?,?,?,?,?)"
	res, err := repo.Conn.Exec(peticio, regisrtres.Data.Unix(), regisrtres.Precipitacio, regisrtres.TempMaxima, regisrtres.TempMinima, regisrtres.Humitat)
	if err != nil{
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	regisrtres.ID = id

	return &regisrtres, nil
}
func (repo *SQLliteRepository) ObtenirTotsRegistres() ([]Registres, error){
	consulta := "select * from registres order by data_registre"
	rows, err := repo.Conn.Query(consulta)
	if err != nil{
		return nil, err
	}
	//trancament de la conecció a la bd
	defer rows.Close()

	var conjuntValors []Registres

	for rows.Next(){
		var a Registres
		var unixTime int64
		rows.Scan(
			&a.ID,
			&unixTime,
			&a.Precipitacio,
			&a.TempMaxima,
			&a.TempMinima,
			&a.Humitat,		
		)
		if err != nil{
			return nil, err
		}
		a.Data = time.Unix(unixTime, 0)
		conjuntValors = append(conjuntValors, a)
	}
	return conjuntValors, nil
}
func (repo *SQLliteRepository) ObtenirRegistresPerID(id int64) (*Registres, error){
	fila := repo.Conn.QueryRow("select * from registres where id = ?", id)

	var regisrtre Registres
	var unixTime int64
	err := fila.Scan(
		&regisrtre.ID,
		&unixTime,
		&regisrtre.Precipitacio,
		&regisrtre.TempMaxima,
		&regisrtre.TempMinima,
		&regisrtre.Humitat,		
	)
	if err != nil{
		return nil, err
	}
	regisrtre.Data = time.Unix(unixTime, 0)

	return &regisrtre, nil
}

func (repo *SQLliteRepository) ActualitzarRegistre(id int64, actulitzar Registres) (error){
	if id== 0{
		return errors.New("la Id indicada és incorrecte")
	}
	consulta := "Update registres set data_registre = ?, precipitacio = ?, temp_maxima = ?, temp_minima = ?, humitat =? where id = ?"
	res, err:=repo.Conn.Exec(consulta, actulitzar.Data.Unix(), actulitzar.Precipitacio, actulitzar.TempMaxima, actulitzar.TempMinima, actulitzar.Humitat, id)
	if err != nil{
		return  err
	}
	rowsAffected, err:=res.RowsAffected()
	if err != nil{
		return  err
	}
	if rowsAffected == 0{
		return errUpdateField
	}
	return nil
}
func (repo *SQLliteRepository) BorrarRegistre(id int64)error{
	
	res, err:= repo.Conn.Exec("delete from registres where id = ?", id)
	if err != nil{
		return err
	}
	rowsAffected, err :=res.RowsAffected()
	if err != nil{
		return err
	}
	if rowsAffected == 0{
		return errDeleteField
	}
	return nil
}
