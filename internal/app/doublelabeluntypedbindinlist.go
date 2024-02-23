package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"fynebind/internal/model"
	widget2 "fynebind/pkg/widget"
	"image/color"
	"math/rand"
	"time"
)

func doubleLabelUntypedBindInList() *fyne.Container {
	var boundDataList []binding.ExternalUntyped

	for i := 0; i < 5; i++ {
		myData := createMyData(rand.Int())
		boundDataList = append(boundDataList, binding.BindUntyped(&myData))
	}

	list := widget.NewList(
		func() int {
			return len(boundDataList)
		},
		func() fyne.CanvasObject {
			return widget2.NewDoubleLabelWithData(nil, DoubleLabelMyDataContentExtractorFn)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			boundData := boundDataList[id]
			object.(*widget2.DoubleLabel).Bind(boundData)
		},
	)

	button1 := widget.NewButton("Change data #1", func() {
		if len(boundDataList) >= 1 {
			boundValue := boundDataList[0]
			value, _ := boundValue.Get()
			myData := value.(model.MyData)
			myData.Sub.Data = time.Now().Format("15:04:05")
			boundValue.Set(myData)
		}
	})

	button2 := widget.NewButton("Change data #2", func() {
		if len(boundDataList) >= 2 {
			boundValue := boundDataList[1]
			value, _ := boundValue.Get()
			myData := value.(model.MyData)
			myData.Sub.Data = time.Now().Format("15:04:05")
			boundValue.Set(myData)
		}
	})

	captionPanel := createCaptionPanel("List with bound items", "items in list listens to changes in bound data")
	buttonPanel := container.NewVBox(button1, button2)
	mainPanel := container.NewPadded(container.NewBorder(captionPanel, buttonPanel, nil, nil, list))
	bg := getBackgroundPanel(color.RGBA{R: 128, G: 16, B: 16, A: 24})

	return container.NewStack(bg, mainPanel)
}
