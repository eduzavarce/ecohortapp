package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolBar(window fyne.Window) *widget.Toolbar{
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
	widget.NewToolbarAction(theme.DocumentCreateIcon(), func(){}),
	widget.NewToolbarAction(theme.ViewRefreshIcon(),func(){app.actualizarDatosClimaContenido()} ),
	widget.NewToolbarAction(theme.SettingsIcon(), func(){}),
	)
	return toolBar
}