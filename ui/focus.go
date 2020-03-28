package ui

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
)

type SwitchingFocusChain struct {
	widgets   map[string][]tui.Widget
	activeSet string
}

func NewSwitchingFocusChain() *SwitchingFocusChain {
	FocusChain := SwitchingFocusChain{
		widgets: make(map[string][]tui.Widget),
	}
	return &FocusChain
}

func (c *SwitchingFocusChain) ActiveSet() string {
	return c.activeSet
}

func (c *SwitchingFocusChain) SetActiveSet(a string) error {
	if _, ok := c.widgets[a]; ok {
		c.activeSet = a
		return nil
	}
	return fmt.Errorf("%s not defined", a)
}

func (c *SwitchingFocusChain) Set(name string, ws ...tui.Widget) {
	if len(c.widgets) < 1 {
		c.activeSet = name
	}
	c.widgets[name] = ws
}

func (c *SwitchingFocusChain) FocusNext(current tui.Widget) tui.Widget {
	for i, w := range c.widgets[c.activeSet] {
		if w != current {
			continue
		}
		if i < len(c.widgets[c.activeSet])-1 {
			return c.widgets[c.activeSet][i+1]
		}
		return c.widgets[c.activeSet][0]
	}
	// If we get here, then the active set may have chaged
	return c.FocusDefault()
}

// FocusPrev returns the widget in the ring that is before the given widget.
func (c *SwitchingFocusChain) FocusPrev(current tui.Widget) tui.Widget {
	for i, w := range c.widgets[c.activeSet] {
		if w != current {
			continue
		}
		if i <= 0 {
			return c.widgets[c.activeSet][len(c.widgets[c.activeSet])-1]
		}
		return c.widgets[c.activeSet][i-1]
	}
	// If we get here, then the active set may have chaged
	return c.FocusDefault()
}

// FocusDefault returns the default widget for when there is no widget
// currently focused.
func (c *SwitchingFocusChain) FocusDefault() tui.Widget {
	if len(c.widgets[c.activeSet]) == 0 {
		return nil
	}
	return c.widgets[c.activeSet][0]
}
