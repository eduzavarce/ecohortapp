package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// Obtener los datos de la api
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	//Insretamos los datos en el contenedor
	climaDadesContent := container.NewGridWithColumns(4,precipitacio, tempMax, tempMin, humitat )

	app.ClimaDadesContainer = climaDadesContent

	//barra de herramientas
	toolBar := app.getToolBar(app.MainWindow)
	pronosticTabContenidor := app.pronosticTab()
	// pestanyes de l'aplicacio
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Pronostic", theme.HomeIcon(),pronosticTabContenidor),
		container.NewTabItemWithIcon("Diari Metorologic", theme.InfoIcon(),canvas.NewText("hjkaslhdkljasdfh", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	//Agregamos el contenedor a la ventana
	finalContent :=container.NewVBox(climaDadesContent, toolBar, tabs)

	//cridar la pagina principal y afegirem el contenidor
	app.MainWindow.SetContent((finalContent))

	//Go routine
	go func() {
		for range time.Tick(time.Second * 30){
			app.actualizarDatosClimaContenido()
		}
	}()
}

func (app *Config) actualizarDatosClimaContenido(){
	app.InfoLog.Println("refrecando los datos")
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	app.ClimaDadesContainer.Objects = []fyne.CanvasObject{precipitacio, tempMax, tempMin, humitat}
	app.ClimaDadesContainer.Refresh()

	image2 := app.obtenirGrafic()
	app.PronosticGraficContainer.Objects = []fyne.CanvasObject{image2}
	app.PronosticGraficContainer.Refresh()
}

/* api key aemet
eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJlZHV6YXZhcmNlQGdtYWlsLmNvbSIsImp0aSI6ImIxYzZiYmFhLTcyY2MtNGJiZS1iZmZkLTI3ZDU2ZjBlZjQ3ZiIsImlzcyI6IkFFTUVUIiwiaWF0IjoxNjkwNDc0NTY1LCJ1c2VySWQiOiJiMWM2YmJhYS03MmNjLTRiYmUtYmZmZC0yN2Q1NmYwZWY0N2YiLCJyb2xlIjoiIn0.6XDOSB6_eFLQwS28nsnDbf8DlPWuIxlk7FXVLwmmeHc
*/