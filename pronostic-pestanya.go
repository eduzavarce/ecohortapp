package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pronosticTab() *fyne.Container {
	grafic := app.obtenirGrafic()
	contenidorGrafic := container.NewVBox(grafic)
	app.PronosticGraficContainer = contenidorGrafic
return contenidorGrafic
}

func (app *Config)obtenirGrafic() *canvas.Image {
apiURL := fmt.Sprint("https://my.meteoblue.com/visimage/meteogram_web_hd?look=KILOMETER_PER_HOUR%2CCELSIUS%2CMILLIMETER%2Cdarkmode&apikey=5838a18e295d&winddirection=3char&temperature=C&windspeed=kmh&precipitationamount=mm&city=Abrera&iso2=es&lat=41.5168&lon=1.901&asl=111&tz=Europe%2FMadrid&lang=en&sig=7e20e2dc955af53a0b9cb36d86b5afb8")
err := app.descarregarArxiu(apiURL, "pronostic.png")
var image *canvas.Image


	if err != nil{
		//imagen por errores
		image = canvas.NewImageFromResource(resourceNodisponiblePng)
	}else {
		//generamos imagen
		image =canvas.NewImageFromFile("pronostic.png")
	}
	image.SetMinSize(fyne.Size{
		Width: 770,
		Height: 480,
	})

	// Determinem la situacio

	image.FillMode = canvas.ImageFillContain



	return image

}
func (app *Config) descarregarArxiu(URL, nomArxiu string) error{

	res, err := app.HTTPClient.Get(URL)
	if err != nil{
		return err
	}
	if res.StatusCode != 200{
		return errors.New("codigo de respuesta erroneo")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil{
		return err
	}
	defer res.Body.Close()

	//decodificar imagen

	img, _, err:= image.Decode(bytes.NewReader(body))
	if err != nil{
		return err
	}
	//gerneramos la salida indicada

	sortida, err := os.Create(fmt.Sprintf("./%s", nomArxiu))
	if err != nil{
		return err
	}
	//codificar png
	png.Encode(sortida,img)

	return nil

}



/*

https://my.meteoblue.com/visimage/meteogram_web_hd?look=KILOMETER_PER_HOUR%2CCELSIUS%2CMILLIMETER%2Cdarkmode&apikey=5838a18e295d&winddirection=3char&temperature=C&windspeed=kmh&precipitationamount=mm&city=Abrera&iso2=es&lat=41.5168&lon=1.901&asl=111&tz=Europe%2FMadrid&lang=en&sig=7e20e2dc955af53a0b9cb36d86b5afb8

*/