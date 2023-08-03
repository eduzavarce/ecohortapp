package main

import (
	"ecohortapp/repository"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) registresTab() *fyne.Container {
	app.RegistresTable = app.getRegistresTable()

	registresContenidor := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1,app.RegistresTable),
	)
	return registresContenidor
}

func (app *Config) getRegistresTable() *widget.Table {
		data := app.getRegistresSlice()

		app.Registres = data		//Definim l'estructura del widget per crear una nova taula amb fyne
		t := widget.NewTable(
			func() (int, int) {
				return len(app.Registres), len(app.Registres[0])
			},
			func() fyne.CanvasObject {
				ctr := container.NewVBox(widget.NewLabel(""))
				return ctr
			},
			func(i widget.TableCellID, o fyne.CanvasObject) {
				if i.Col == (len(app.Registres[0])-1) && i.Row != 0 {
					//Ultima cel.la - situa un botò
					w := widget.NewButtonWithIcon("Borrar", theme.DeleteIcon(), func() {
						//Presentem un dialeg de confirmació
						dialog.ShowConfirm("Borrar?", "", func(deleted bool) {
							if deleted {
								id, _ := strconv.Atoi(app.Registres[i.Row][0].(string)) //Transformem el identificador a decimal sencer
								err := app.DB.BorrarRegistre(int64(id))        //Invoquem el metode per borrar a partir d'un id
								//Capturem possibles errors
								if err != nil {
									app.ErrorLog.Println(err)
								}
							}
							//Forcem el refresc de la taula
							app.actualitzarRegistresTable()
						}, app.MainWindow)
					})
					//Creem un widget d'alta importancia per mostrar un missatge destacat
					w.Importance = widget.HighImportance
	
					//Definim el contenidor a on situarem el objecte corresponent a el boto.
					o.(*fyne.Container).Objects = []fyne.CanvasObject{
						w,
					}
				} else {
					//situarem la informació rebuda en el slice, recordem que primer gestiona la fila i després la columna
					o.(*fyne.Container).Objects = []fyne.CanvasObject{
						widget.NewLabel(app.Registres[i.Row][i.Col].(string)),
					}
				}
			})
	
		//Establim el ample de les diferents celdes
		colWidths := []float32{50, 100, 100, 100, 100, 100, 110}
		//Executem una estructura for per aplicar cada un de els amples amb el metode SetColumnWidth
		for i := 0; i < len(colWidths); i++ {
			t.SetColumnWidth(i, colWidths[i])
		}
	
		return t
}

func (app *Config) getRegistresSlice() [][]interface{} {
	var slice [][]interface{}
	//llamada al metodo inferior
	registres, err:= app.registresActuals()
	if err != nil {
		app.ErrorLog.Println(err)
	}
	//append 1 fila enunciados
	slice = append(slice,[]interface{}{"ID", "Data", "Precipitacio", "T.Máxima", "t.Mínima", "Humitat", "Opcions"})
	// ejecutamos un for para poblar las filas
	for _,x := range registres{
		//creamos una interficie vacia por columne
		var filaActual []interface{}

		// Anem vinculant els valors en cada una de les columnes
		filaActual = append(filaActual, strconv.FormatInt(x.ID, 10)) 
		filaActual = append(filaActual, x.Data.Format("2006-01-02"))
		filaActual = append(filaActual, fmt.Sprintf("%d%%", x.Precipitacio))
		filaActual = append(filaActual, fmt.Sprintf("%d", x.TempMaxima))
		filaActual = append(filaActual, fmt.Sprintf("%d", x.TempMinima))
		filaActual = append(filaActual, fmt.Sprintf("%d%%", x.Humitat))
		filaActual = append(filaActual, widget.NewButton("Borrar", func(){}))

		//Afegir la fila al slice
		slice = append(slice, filaActual)
	}



	return slice
}

func (app *Config) registresActuals() ([]repository.Registres, error) {

registres, err := app.DB.ObtenirTotsRegistres()
if err != nil {
	app.ErrorLog.Println(err)
	return nil, err
}
return registres, nil
}