package main

import (
	"github.com/g3n/engine/gui"
)

func init() {
	TestMap["gui.tabbar"] = &GuiTabBar{}
}

type GuiTabBar struct {
	tb *gui.TabBar
}

func (t *GuiTabBar) Initialize(ctx *Context) {

	// Button for adding tabs
	b1 := gui.NewButton("Add Tab")
	b1.SetPosition(10, 10)
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Info("button 1 OnClick")
	})
	ctx.Gui.Add(b1)

	// Creates TabBar
	t.tb = gui.NewTabBar(0, 0)
	tby := b1.Position().Y + b1.Height() + 10
	t.tb.SetPosition(b1.Position().X, tby)
	ctx.Gui.Add(t.tb)

	// Resizes TabBar when main window resizes
	ctx.Gui.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		t.tb.SetSize(ctx.Gui.ContentWidth()-t.tb.Position().X-10, ctx.Gui.ContentHeight()-tby-10)
	})
	ctx.Gui.Dispatch(gui.OnResize, nil)
}

func (t *GuiTabBar) Render(ctx *Context) {
}
