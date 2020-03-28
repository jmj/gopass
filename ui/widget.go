package ui

import (
	"github.com/marcusolsson/tui-go"
)

type Widget struct {
	widget tui.Widget
	name   string
}

func NewWidget(n string, w tui.Widget) *Widget {
	return &Widget{
		name:   n,
		widget: w,
	}
}

func (w *Widget) Widget() tui.Widget {
	return w.widget
}

func (w *Widget) Name() string {
	return w.name
}
