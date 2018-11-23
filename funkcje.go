package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/printsupport"
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

	if oknoDialog.Exec() == int(widgets.QDialog__Accepted) {

		var plikNazwa = oknoDialog.SelectedFiles()[0] //nazwa wybranego plik w oknie

		if core.QFile_Exists(plikNazwa) { //sprawdza czy plik istnieje

			var Qplik = core.NewQFile2(plikNazwa) //tworzy Qfile z wybranego pliku

			if Qplik.Open(core.QIODevice__ReadOnly) { //czy do odczytu

				var qbTablica = Qplik.ReadAll()                            //przekazuje z Qpliku do QByteArray
				var qtkodowanie = core.QTextCodec_CodecForHtml2(qbTablica) // odczytuje kodowanie
				var tekstString = qtkodowanie.ToUnicode(qbTablica)         // na string unicode

				e.oknoEdytor.SetPlainText(tekstString) //w e wypisuje surowy tekst przekazany w stringu
			}
		}
	}

}

func (e *Edytor) Zapisz() {
	var oknoDialog = widgets.NewQFileDialog2(e, "Zapisz jako", "", "pliki tekstowe (*.txt)")
	oknoDialog.SetAcceptMode(widgets.QFileDialog__AcceptSave)

	if oknoDialog.Exec() == int(widgets.QDialog__Accepted) {

		var plikNazwa = oknoDialog.SelectedFiles()[0]                           //nazwa wybranego plik w oknie
		var plik = gui.NewQTextDocumentWriter3(plikNazwa, core.NewQByteArray()) // tworzy plik tekstowy
		plik.Write(e.oknoEdytor.Document())                                     //przekazuje zawartosc edytora i zapisuje
	}

}

func (e *Edytor) Drukuj() {

	var drukarka = printsupport.NewQPrinter(printsupport.QPrinter__HighResolution)
	var oknoDialog = printsupport.NewQPrintDialog2(e)

	if oknoDialog.Exec() == int(widgets.QDialog__Accepted) {

		e.oknoEdytor.Print(drukarka)
	}

}
