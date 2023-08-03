package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolBar(window fyne.Window) *widget.Toolbar{
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
	widget.NewToolbarAction(theme.DocumentCreateIcon(), func(){
		app.addRegistreDialog()
	}),
	widget.NewToolbarAction(theme.ViewRefreshIcon(),func(){app.actualizarDatosClimaContenido()} ),
	widget.NewToolbarAction(theme.SettingsIcon(), func(){}),
	)
	return toolBar
}

func (app *Config) addRegistreDialog() dialog.Dialog{
	dataRegistreEntrada := widget.NewEntry()
	precipitacioRegistreEntrada := widget.NewEntry()
	tempMaximaRegistreEntrada := widget.NewEntry()
	tempMinimaRegistreEntrada := widget.NewEntry()
	humitatRegistreEntrada := widget.NewEntry()

	app.AfergirRegistresDataRegistreEntrada = dataRegistreEntrada
	app.AfergirRegistresPrecipitacioEntrada = precipitacioRegistreEntrada
	app.AfergirRegistresTempMaximaEntrada = tempMaximaRegistreEntrada
	app.AfergirRegistresTempMinimaEntrada = tempMinimaRegistreEntrada
	app.AfergirRegistresHumitatEntrada = humitatRegistreEntrada
	//placeholders

	dataRegistreEntrada.PlaceHolder = "YYYY-MM-DD"

	addForm := dialog.NewForm(
		"Afegir registre", 
		"Afegir", 
		"Cancelar",
		//etiquetas
		[]*widget.FormItem{
			{Text: "Data Registre", Widget: dataRegistreEntrada},
			{Text: "Probabilitat de precipitacio", Widget: precipitacioRegistreEntrada},
			{Text: "Tempreatura maxima", Widget: tempMaximaRegistreEntrada}, 
			{Text: "Temperatura minima", Widget: tempMinimaRegistreEntrada},
			{Text: "Humitat", Widget: humitatRegistreEntrada},
		},
		func(valid bool){
			if valid{
				//pend de validacio
			}
		},
		app.MainWindow)
		addForm.Resize(fyne.Size{Width: 400})
		addForm.Show()
		return addForm
}