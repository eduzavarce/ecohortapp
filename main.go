package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// Creamos un struct con todas las configuraciones de la app

type Config struct {
	App fyne.App //Almacenaremos la GUI de fyne
	InfoLog *log.Logger //Log de Acciones del usuario
	ErrorLog *log.Logger //Log de Errores
}

var myApp Config 

func main () {
	//Crearemos una app de Fyne (un canvas sovre el que trabajaremos)
	fyneApp := app.NewWithID("cat.cibernarium.ecohortapp")
	myApp.App = fyneApp

	//Crearemos nuestros Logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) 
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile)
	//Conectamos con al base de datos

	// Crearemos un repositorio de base de datos

	//Crearemos y definiremos el tamaño de la pantalla de la aplicación
	win := fyneApp.NewWindow("Eco Hort App")
	win.ShowAndRun()
	//Mostraremos y ejecutaremos la app

}