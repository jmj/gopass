package ui

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
)

type ViewManager struct {
	views      map[string]View
	app        tui.UI
	tabManager *SwitchingFocusChain
}

func NewViewManager(app tui.UI) *ViewManager {
	vm := &ViewManager{
		app:        app,
		views:      make(map[string]View),
		tabManager: NewSwitchingFocusChain(),
	}

	return vm
}

func (v *ViewManager) AddView(view View) error {
	if _, ok := v.views[view.Name()]; ok {
		return fmt.Errorf("View with name %s already exists", view.Name())
	}
	v.views[view.Name()] = view
	v.tabManager.Set(view.Name(), view.GetTabOrder()...)
	return nil
}

func (v *ViewManager) GetView(view string) View {
	if _, ok := v.views[view]; ok {
		return v.views[view]
	}
	return nil
}

func (v *ViewManager) ActivateView(view string) error {
	if _, ok := v.views[view]; !ok {
		return fmt.Errorf("View %s does not exist", view)
	}

	v.app.SetWidget(v.views[view].ViewWidget())
	if err := v.tabManager.SetActiveSet(view); err != nil {
		return err
	}

	return nil
}
