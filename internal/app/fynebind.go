package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fynebind/internal/model"
	"image/color"
	"strconv"
)

func ApplicationContent() *fyne.Container {
	doubleLabelUntypedBindContainer := doubleLabelUntypedBind()
	doubleLabelListUntypedBindContainer := doubleLabelListUntypedBind()
	doubleLabelUntypedBindListContainer := doubleLabelUntypedBindInList()

	split := container.NewHSplit(doubleLabelListUntypedBindContainer, doubleLabelUntypedBindListContainer)

	mainLayout := container.NewBorder(nil, doubleLabelUntypedBindContainer, nil, nil, split)
	return mainLayout
}

func createCaptionPanel(header, description string) *fyne.Container {
	caption := widget.NewLabel(header)
	caption.TextStyle.Bold = true
	caption.Importance = widget.HighImportance
	subCaption := widget.NewLabel("(" + description + ")")
	subCaption.Importance = widget.HighImportance
	captionPanel := container.NewVBox(caption, subCaption)
	return captionPanel
}

func getBackgroundPanel(rgba color.RGBA) *canvas.Rectangle {
	bg := canvas.NewRectangle(rgba)
	bg.CornerRadius = theme.SelectionRadiusSize()
	bg.StrokeColor = rgba
	bg.StrokeWidth = 2.0
	return bg
}

func createMyData(i int) model.MyData {
	id := strconv.Itoa(i)
	return model.MyData{Name: "Name " + id, Info: "Info " + id, Sub: model.SubData{Data: "Data " + id}}
}

func DoubleLabelMyDataContentExtractorFn(boundData binding.ExternalUntyped) (string, string) {
	value, _ := boundData.Get()
	data := value.(model.MyData)
	return data.Name, data.Sub.Data
}
