package main

import (
	"database/sql"
	"ecohortapp/repository"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	_ "github.com/glebarez/go-sqlite"
)

// Creamos un struct con todas las configuraciones de la app

type Config struct {
	App fyne.App //Almacenaremos la GUI de fyne
	InfoLog *log.Logger //Log de Acciones del usuario
	ErrorLog *log.Logger //Log de Errores
	MainWindow fyne.Window //finestra principal
	ClimaDadesContainer *fyne.Container // Contenidor del clima
	HTTPClient http.Client
	PronosticGraficContainer *fyne.Container
	DB repository.Repository
	Registres [][]interface{}
	RegistresTable *widget.Table
	AfergirRegistresDataRegistreEntrada *widget.Entry
	AfergirRegistresPrecipitacioEntrada *widget.Entry
	AfergirRegistresTempMaximaEntrada *widget.Entry
	AfergirRegistresTempMinimaEntrada *widget.Entry
	AfergirRegistresHumitatEntrada *widget.Entry
	

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
	sqlDB, err := myApp.connectSql()
	if err!=nil{
		log.Panic(err)
	}
	// Crearemos un repositorio de base de datos
	myApp.setupDB(sqlDB)
	//Crearemos y definiremos el tamaño de la pantalla de la aplicación
	myApp.MainWindow = fyneApp.NewWindow("Eco Hort App")
	myApp.MainWindow.Resize(fyne.NewSize(800,500)) //tamaño ventana principal
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster() //Establecemos que es la ventana principal

	myApp.makeUI()
	
	//Mostraremos y ejecutaremos la app
	myApp.MainWindow.ShowAndRun()
}
func (app *Config) connectSql()(*sql.DB, error){
	path := ""
	if os.Getenv("DB_PATH") != ""{
		path = os.Getenv("DB_PATH")
	}else{
		path = app.App.Storage().RootURI().Path()+"/sql.db"
		app.InfoLog.Println("db in: ", path)
		
	}
	db, err :=sql.Open("sqlite", path)
		if err != nil{
			return nil, err
		}
		return db, nil
	
}

func (app *Config) setupDB(sqlDB *sql.DB){
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
	

}