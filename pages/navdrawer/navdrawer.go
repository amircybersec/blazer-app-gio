package navdrawer

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"blazerproxy.org/icon"
	page "blazerproxy.org/pages"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Page holds the state for a page demonstrating the features of
// the NavDrawer component.
type Page struct {
	widget.List
	*page.Router
	udpCheckbox            widget.Bool
	tcpCheckbox            widget.Bool
	swtch               widget.Bool
	radioButtonsGroup widget.Enum
	disableBtn        widget.Clickable
	saveButton        widget.Clickable
}

// New constructs a Page with the provided router.
func New(router *page.Router) *Page {
	return &Page{
		Router: router,
	}
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Settings",
		Icon: icon.SettingsIcon,
	}
}

var lineEditor = &widget.Editor{
	SingleLine: true,
	Submit:     true,
}

var resolverEditor = &widget.Editor{
	SingleLine: true,
	Submit:     true,
}

var domainEditor = &widget.Editor{
	SingleLine: true,
	Submit:     true,
}

var reporURLEditor = &widget.Editor{
	SingleLine: true,
	Submit:     true,
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	widgets := []layout.Widget{
		func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.H6(th, "Local Address").Layout),
				layout.Rigid(func(gtx C) D {
					e := material.Editor(th, lineEditor, "localhost:8080")
					e.Font.Style = font.Italic
					border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
					return border.Layout(gtx, func(gtx C) D {
						return layout.UniformInset(unit.Dp(10)).Layout(gtx, e.Layout)
					})
				}),
			)
		},
		func(gtx C) D {
			gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(5))
			gtx.Constraints.Max.Y = gtx.Constraints.Min.Y

			dr := image.Rectangle{Max: gtx.Constraints.Min}
			paint.LinearGradientOp{
				Stop1:  layout.FPt(dr.Min),
				Stop2:  layout.FPt(dr.Max),
				Color1: color.NRGBA{R: 0x10, G: 0xff, B: 0x10, A: 0xFF},
				Color2: color.NRGBA{R: 0x10, G: 0x10, B: 0xff, A: 0xFF},
			}.Add(gtx.Ops)
			defer clip.Rect(dr).Push(gtx.Ops).Pop()
			paint.PaintOp{}.Add(gtx.Ops)
			return layout.Dimensions{
				Size: gtx.Constraints.Max,
			}
		},
		func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.H6(th, "Test Domain").Layout),
				layout.Rigid(func(gtx C) D {
					e := material.Editor(th, domainEditor, "example.com")
					e.Font.Style = font.Italic
					border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
					return border.Layout(gtx, func(gtx C) D {
						return layout.UniformInset(unit.Dp(10)).Layout(gtx, e.Layout)
					})
				}),
			)
		},
		func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.H6(th, "Resolver").Layout),
				layout.Rigid(func(gtx C) D {
					e := material.Editor(th, resolverEditor, "8.8.8.8")
					e.Font.Style = font.Italic
					border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
					return border.Layout(gtx, func(gtx C) D {
						return layout.UniformInset(unit.Dp(10)).Layout(gtx, e.Layout)
					})
				}),
			)
		},
		func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.H6(th, "Report URL").Layout),
				layout.Rigid(func(gtx C) D {
					e := material.Editor(th, reporURLEditor, "report.example.com")
					e.Font.Style = font.Italic
					border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
					return border.Layout(gtx, func(gtx C) D {
						return layout.UniformInset(unit.Dp(10)).Layout(gtx, e.Layout)
					})
				}),
			)
		},
		func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(
					material.CheckBox(th, &p.udpCheckbox, "TCP").Layout,
				),
				layout.Rigid(
					material.CheckBox(th, &p.tcpCheckbox, "UDP").Layout,
				),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Left: unit.Dp(16)}.Layout(gtx,
						material.Switch(th, &p.swtch, "Example Switch").Layout,
					)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Left: unit.Dp(16)}.Layout(gtx, func(gtx C) D {
						text := "enabled"
						if !p.swtch.Value {
							text = "disabled"
							gtx = gtx.Disabled()
						}
						btn := material.Button(th, &p.disableBtn, text)
						return btn.Layout(gtx)
					})
				}),
			)
		},
		func(gtx C) D {
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(material.RadioButton(th, &p.radioButtonsGroup, "r1", "RadioButton1").Layout),
			)
		},
		func(gtx C) D {
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(material.Button(th, &p.saveButton,"Save").Layout),
			)
		
		},
	}

	return material.List(th, &p.List).Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, widgets[i])
	})
}
