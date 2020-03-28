package main

import (
	"github.com/marcusolsson/tui-go"
)

var (
	detailFields map[string]*tui.Label
	views        map[string]tui.Widget
	mainTabOrder []tui.Widget
)

func init() {
	detailFields = make(map[string]*tui.Label)
	views = make(map[string]tui.Widget)
	mainTabOrder = make([]tui.Widget, 0, 0)

}

func ui(app tui.UI) tui.Widget {

	mainWin := tui.NewGrid(0, 0)
	mainWin.SetRowStretch(0, 10)
	mainWin.SetRowStretch(1, 1)
	mainWin.SetBorder(true)

	mainWin.AppendRow(detailSelectBox(app))
	mainWin.AppendRow(btnBox(app))
	views["main"] = mainWin
	FocusChain.Set("main", mainTabOrder...)

	addWin := addDialog(app)
	views["add"] = addWin

	return mainWin
}

func addDialog(app tui.UI) tui.Widget {
	ne := tui.NewEntry()
	ue := tui.NewEntry()
	pe := tui.NewEntry()

	genBtn := tui.NewButton("[ Generate ]")
	genBtn.OnActivated(func(b *tui.Button) {
		pe.SetText(GenPasswordDefault())
	})

	saveBtn := tui.NewButton("[ Save ]")
	saveBtn.OnActivated(func(b *tui.Button) {
		// implement later
	})

	canBtn := tui.NewButton("[ Cancel ]")
	canBtn.OnActivated(func(b *tui.Button) {
		app.SetWidget(views["main"])
		FocusChain.SetActiveSet("main")
	})

	fieldsBox := tui.NewVBox(
		tui.NewHBox(tui.NewLabel("Name: "), ne, tui.NewSpacer()),
		tui.NewHBox(tui.NewLabel("URL: "), ue, tui.NewSpacer()),
		tui.NewHBox(tui.NewLabel("Password: "), pe, tui.NewSpacer()),
	)
	fieldsBox.SetBorder(true)

	vb := tui.NewVBox(
		fieldsBox,
		tui.NewSpacer(),
		tui.NewHBox(genBtn, saveBtn, canBtn, tui.NewSpacer()),
	)
	vb.SetBorder(true)

	FocusChain.Set("add", ne, ue, pe, genBtn, saveBtn, canBtn)

	return vb
}

func detailSelectBox(app tui.UI) tui.Widget {
	detailSelectBox := tui.NewHBox()

	db := detailBox(app)
	detailSelectBox.Append(selectBox(app))
	detailSelectBox.Append(db)
	return detailSelectBox

}

func detailBox(app tui.UI) tui.Widget {
	detailFields["name"] = tui.NewLabel("")
	detailFields["url"] = tui.NewLabel("")
	detailFields["password"] = tui.NewLabel("")

	vLayout := tui.NewVBox(
		tui.NewHBox(tui.NewLabel("Name: "), detailFields["name"]),
		tui.NewHBox(tui.NewLabel("URL: "), detailFields["url"]),
		tui.NewHBox(tui.NewLabel("Password: "), detailFields["password"]),
		tui.NewSpacer(),
	)

	box := tui.NewVBox(vLayout)
	box.SetBorder(true)
	return box
}

func selectBox(app tui.UI) tui.Widget {
	n := make([]string, len(passwords))

	indx := 0
	for k := range passwords {
		n[indx] = k
		indx++
	}

	lb := tui.NewList()
	lb.AddItems(n...)

	lb.OnSelectionChanged(func(l *tui.List) {
		item := l.SelectedItem()
		if _, ok := passwords[item]; ok {
			detailFields["name"].SetText(passwords[item].Name)
			detailFields["url"].SetText(passwords[item].URL)
			detailFields["password"].SetText(passwords[item].Password)
		}
	})

	mainTabOrder = append(mainTabOrder, lb)

	box := tui.NewVBox(lb)
	box.SetBorder(true)
	return box
}

func btnBox(app tui.UI) tui.Widget {
	btnBox := tui.NewHBox()

	btnBox.Insert(0, quitButton(app))
	btnBox.Insert(1, addButton(app))
	btnBox.Insert(2, tui.NewSpacer())
	return btnBox
}

func quitButton(app tui.UI) tui.Widget {

	quitBtn := tui.NewButton("[ Quit ]")
	quitBtn.OnActivated(func(b *tui.Button) {
		app.Quit()
	})

	mainTabOrder = append(mainTabOrder, quitBtn)
	return tui.NewPadder(1, 0, quitBtn)
}

func addButton(app tui.UI) tui.Widget {
	addBtn := tui.NewButton("[ Add ]")
	addBtn.OnActivated(func(b *tui.Button) {
		app.SetWidget(views["add"])
		FocusChain.SetActiveSet("add")
	})

	mainTabOrder = append(mainTabOrder, addBtn)
	return tui.NewPadder(1, 0, addBtn)
}
