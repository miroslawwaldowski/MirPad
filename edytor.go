package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

//struktura edytor, która opisuje co bedzie zaiwerało góówne okno
type Edytor struct {
	widgets.QMainWindow // elementy z "klasy" QMainWindow

	oknoEdytor *widgets.QTextEdit //okno tekstowe
}

// funkcja inicjująca
func initEdytor() *Edytor {
	//tworzy strukturę okno typu Edytor
	var okno = NewEdytor(nil, 0)
	// w strukturze okno inicjuje funkcją nowe oknoEdytor
	okno.oknoEdytor = widgets.NewQTextEdit(okno)
	//utworzony oknoEdytor w strukturze okno typu Edytor przypisuje do zmiennej oknoE
	var oknoE = okno.oknoEdytor
	//ustawia oknoE centralnie w okno
	okno.SetCentralWidget(oknoE)

	okno.SetToolButtonStyle(core.Qt__ToolButtonFollowStyle) //ustawia styl paska narzędzi zgodnie z systemem
	//metody tworzące menu plik i edytuj
	okno.setupPlik()
	okno.setupEdytuj()
	return okno
}

func (e *Edytor) setupPlik() {
	//tworzy menu plik

	var menu = e.MenuBar().AddMenu2("Plik")
	var pasek = e.AddToolBar3("Plik") //pasek narzedzi
	//
	var nowyIkona = gui.QIcon_FromTheme2("nowy", gui.NewQIcon5(":/qml/ikony/new.png")) //deklaruje ikone
	var a = menu.AddAction2(nowyIkona, "Nowy")                                         //dodaje do menu
	a.SetShortcuts2(gui.QKeySequence__New)                                             //ustala skrót klawiszowy
	pasek.QWidget.AddAction(a)                                                         //dodaje do paska narzędzi
	//
	var otworzIkona = gui.QIcon_FromTheme2("otwórz", gui.NewQIcon5(":/qml/ikony/open.png"))
	a = menu.AddAction2(otworzIkona, "Otwórz")
	a.SetShortcuts2(gui.QKeySequence__Open)
	pasek.QWidget.AddAction(a)
	//
	var zapiszIkona = gui.QIcon_FromTheme2("zapisz", gui.NewQIcon5(":/qml/ikony/save.png"))
	a = menu.AddAction2(zapiszIkona, "Zapisz")
	a.SetShortcuts2(gui.QKeySequence__Save)
	pasek.QWidget.AddAction(a)
	//
	var drukujIkona = gui.QIcon_FromTheme2("drukuj", gui.NewQIcon5(":/qml/ikony/print.png"))
	a = menu.AddAction2(drukujIkona, "Drukuj")
	a.SetShortcuts2(gui.QKeySequence__Print)
	pasek.QWidget.AddAction(a)
	//
	var wyjscieIkona = gui.QIcon_FromTheme2("wyjście", gui.NewQIcon5(":/qml/ikony/close.png"))
	a = menu.AddAction2(wyjscieIkona, "Wyjście")
	a.SetShortcuts2(gui.QKeySequence__Quit) //sprawdzić skrót

}

func (e *Edytor) setupEdytuj() {
	//tworzy menu edytuj

	var menu = e.MenuBar().AddMenu2("Edytuj")
	var pasek = e.AddToolBar3("Edytuj")
	//
	var wytnijIkona = gui.QIcon_FromTheme2("wytnij", gui.NewQIcon5(":/qml/ikony/cut.png"))
	var a = menu.AddAction2(wytnijIkona, "Wytnij")
	a.SetShortcuts2(gui.QKeySequence__Cut)
	pasek.QWidget.AddAction(a)
	//
	var kopiujIkona = gui.QIcon_FromTheme2("kopiuj", gui.NewQIcon5(":/qml/ikony/copy.png"))
	a = menu.AddAction2(kopiujIkona, "Kopiuj")
	a.SetShortcuts2(gui.QKeySequence__Copy)
	pasek.QWidget.AddAction(a)
	//
	var wklejIkona = gui.QIcon_FromTheme2("wklej", gui.NewQIcon5(":/qml/ikony/paste.png"))
	a = menu.AddAction2(wklejIkona, "Wklej")
	a.SetShortcuts2(gui.QKeySequence__Paste)
	pasek.QWidget.AddAction(a)
	//
	var cofnijIkona = gui.QIcon_FromTheme2("cofnij", gui.NewQIcon5(":/qml/ikony/undo-arrow.png"))
	a = menu.AddAction2(cofnijIkona, "Cofnij")
	a.SetShortcuts2(gui.QKeySequence__Undo)
	pasek.QWidget.AddAction(a)
	//
	var ponowIkona = gui.QIcon_FromTheme2("ponów", gui.NewQIcon5(":/qml/ikony/redo-arrow.png"))
	a = menu.AddAction2(ponowIkona, "Ponów")
	a.SetShortcuts2(gui.QKeySequence__Redo)
	pasek.QWidget.AddAction(a)

}
