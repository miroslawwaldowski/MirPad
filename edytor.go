package main

import (
	"github.com/therecipe/qt/widgets"
)

type Edytor struct {
	widgets.QMainWindow

	oknoEdytor *widgets.QTextEdit
}

func initEdytor() *Edytor {
	var okno = NewEdytor(nil, 0)

	okno.oknoEdytor = widgets.NewQTextEdit(okno)
	var oknoE = okno.oknoEdytor

	okno.SetCentralWidget(oknoE)

	return okno
}
