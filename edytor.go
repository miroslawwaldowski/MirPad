package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

//struktura edytor, która opisuje co bedzie zaiwerało góówne okno
type Edytor struct {
	widgets.QMainWindow

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

	okno.SetToolButtonStyle(core.Qt__ToolButtonFollowStyle)
	//metody tworzące menu plik i edytuj
	okno.setupPlik()
	okno.setupEdytuj()
	return okno
}

func (e *Edytor) setupPlik() {
	//tworzy menu plik
	var menu = e.MenuBar().AddMenu2("&Plik")

	var nowyIkona = gui.QIcon_FromTheme2("nowy", gui.NewQIcon5(":/qml/ikony/new.png"))
	var a = menu.AddAction2(nowyIkona, "&Nowy")
	a.SetShortcuts2(gui.QKeySequence__New)
	menu.AddSeparator()

}

func (e *Edytor) setupEdytuj() {
	//tworzy menu edytuj
	var menu = e.MenuBar().AddMenu2("Edytuj")
	var wytnijIkona = gui.QIcon_FromTheme2("wytnij", gui.NewQIcon5(":/qml/ikony/cut.png"))
	var a = menu.AddAction2(wytnijIkona, "&Wytnij")
	a.SetShortcuts2(gui.QKeySequence__Cut)
	menu.AddSeparator()

}
