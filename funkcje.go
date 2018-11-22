package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func (e *Edytor) Nowy() {
	e.oknoEdytor.Clear()
	// dodać metodę sprawdzającą czy aktywny plik był modyfikowany
}

func (e *Edytor) Otworz() {
	var oknoDialog = widgets.NewQFileDialog2(e, "Otwórz plik", "", "pliki tekstowe (*.txt)")
	oknoDialog.SetAcceptMode(widgets.QFileDialog__AcceptOpen)
	oknoDialog.SetFileMode(widgets.QFileDialog__ExistingFile)

	oknoDialog.Exec()

	var plikNazwa = oknoDialog.SelectedFiles()[0] //wybrany plik w oknie

	if core.QFile_Exists(plikNazwa) { //sprawdza czy plik istnieje

		e.oknoEdytor.SetPlainText("wprowadzić zamiast tego tekst z pliku")
	}

}
