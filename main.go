package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//deklaruję główną aplikację jako strukturę typu QApplication
var aplikacja *widgets.QApplication

func main() {

	//funkcja tworzy aplikację typu QApplication
	aplikacja = widgets.NewQApplication(len(os.Args), os.Args)

	//funkcje wprowadzają informacje do aplikacji
	core.QCoreApplication_SetApplicationName("MirPad")
	core.QCoreApplication_SetApplicationVersion("1.0")

	//deklaruję główne okno aplikacji, które jest tworzone przez funkcję inicjującą
	//jako struktura typuEedytor
	var glowneOkno = initEdytor()
	//ustawia rozmiar
	glowneOkno.SetMinimumSize2(600, 450)
	//pokazuje okno
	glowneOkno.Show()
	//uruchamia stworzoną aplikacje
	aplikacja.Exec()
}
