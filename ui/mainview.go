package ui

type MainView struct {
	BaseView
}

func NewMainView() *MainView {
	mv := MainView{}
	mv.BaseView.widgets = make(map[string]*Widget)

	return &mv
}
