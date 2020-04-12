package main

import "github.com/gizak/termui/v3/widgets"

// Tab represents the tab widget
type Tab struct {
	Widget *widgets.TabPane
}

// Construct creates the tab widget for
// switching between views
func (self *Tab) Construct() {
	widget := widgets.NewTabPane("🌎 Global", " 🇺  USA", "😷 Protect Yourself", "👌 Credits")
	widget.Border = true

	self.Widget = widget
}
