package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)
func (app *Config) getClimaText()(*canvas.Text, *canvas.Text, *canvas.Text, *canvas.Text,){
	var g Diaria
	var precipitacio, tempMax, tempMin, humitat *canvas.Text
	prediccio, err := g.GetPrediccions()
	if err != nil {
		//color texto
		gris := color.RGBA{R:155,G:155,B:155,A:255}
		precipitacio = canvas.NewText("Precipitació: No definit", gris)
		tempMax = canvas.NewText("Temperatura máxima: No definit", gris)
		tempMin = canvas.NewText("Temperatura mínima: No definit", gris)
		humitat = canvas.NewText("Humedad relativa: No definit", gris)
	}else{
		colorDefecte := color.RGBA{R:0,G:180,B:0,A:255}
		 // filtro si hay menos de 50
		 if prediccio.ProbPrecipitacio <50 {
			colorDefecte = color.RGBA{R:180,G:0,B:0,A:255}
		 }
		 //preparamos los strimgs
		 precipitacioText := fmt.Sprintf("Precipitacion: %d%%", prediccio.ProbPrecipitacio)
		tempMaxText := fmt.Sprintf("Temporatura max: %d%%", prediccio.TemperaturaMax)
		tempMinText := fmt.Sprintf("Temporatura max: %d%%", prediccio.TemperaturaMax)
		humitatText := fmt.Sprintf("Precipitacion: %d%%", prediccio.HumitatRelativa)

		precipitacio = canvas.NewText(precipitacioText, colorDefecte)
		tempMax = canvas.NewText(tempMaxText, colorDefecte)
		tempMin = canvas.NewText(tempMinText, colorDefecte)
		humitat = canvas.NewText(humitatText, colorDefecte)

	}
// Alineacion del texto
precipitacio.Alignment = fyne.TextAlignLeading
tempMax.Alignment = fyne.TextAlignCenter
tempMin.Alignment = fyne.TextAlignCenter
humitat.Alignment= fyne.TextAlignTrailing

return precipitacio, tempMax, tempMin, humitat
}	

