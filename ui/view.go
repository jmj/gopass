package ui

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
)

type View interface {
	GetWidget(string) tui.Widget
	Name() string
	ViewWidget() tui.Widget
	GetTabOrder() []tui.Widget
}

type BaseView struct {
	widgets  map[string]*Widget
	name     string
	tabOrder []tui.Widget
}

func (v *BaseView) GetTabOrder() []tui.Widget {
	return v.tabOrder
}

func (v *BaseView) ViewWidget() tui.Widget {
	if _, ok := v.widgets["__top__"]; ok {
		return v.widgets["__top__"].Widget()
	}
	return nil
}

func (v *BaseView) Name() string {
	return v.name
}

func (v *BaseView) GetWidget(name string) *Widget {
	if v, ok := v.widgets[name]; ok {
		return v
	}
	return nil
}

func (v *BaseView) AppendWidget(parent string, widgets ...*Widget) error {
	if parent == "" {
		if len(widgets) > 1 {
			return fmt.Errorf("Only one parentless witget allowed per view")
		}
		if len(v.widgets) > 0 {
			return fmt.Errorf("Only one parentless witget allowed per view")
		}
		v.widgets["__top__"] = widgets[0]
	}

	if _, ok := v.widgets[parent]; !ok {
		return fmt.Errorf("Parent not found")
	}

	var err error
	for _, w := range widgets {
		name := w.Name()

		if _, ok := v.widgets[name]; ok {
			err = fmt.Errorf("Widget with name %s already exists", name)
			break
		}
	}
	if err != nil {
		return err
	}

	switch p := v.widgets[parent].Widget().(type) {
	case *tui.Box:
		for _, w := range widgets {
			p.Append(w.Widget())
			v.widgets[w.Name()] = w
		}
	case *tui.Grid:
		widges := make([]tui.Widget, 0, 0)
		for _, w := range widgets {
			widges = append(widges, w.Widget())
		}
		p.AppendRow(widges...)
		for _, w := range widgets {
			v.widgets[w.Name()] = w
		}
	}
	return nil
}
