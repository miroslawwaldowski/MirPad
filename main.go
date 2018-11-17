package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var aplikacja *widgets.QApplication

func main() {
	aplikacja = widgets.NewQApplication(len(os.Args), os.Args)

	core.QCoreApplication_SetApplicationName("MirPad")
	core.QCoreApplication_SetApplicationVersion("1.0")

	var glowneOkno = initEdytor()

	glowneOkno.SetMinimumSize2(600, 450)

	glowneOkno.Show()
	aplikacja.Exec()
}
