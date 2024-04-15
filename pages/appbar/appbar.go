package appbar

import (
	"fmt"
	"image"
	"image/color"

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

// var (
// 	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
// 	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
// 	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
// 	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
// )

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	heartBtn, plusBtn, startButton          widget.Clickable
	exampleOverflowState widget.Clickable
	favorited, boiling                              bool
	widget.List
	*page.Router
}

// New constructs a Page with the provided router.
func New(router *page.Router) *Page {
	return &Page{
		Router: router,
	}
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		// {
		// 	OverflowAction: component.OverflowAction{
		// 		Name: "Favorite",
		// 		Tag:  &p.heartBtn,
		// 	},
		// 	Layout: func(gtx layout.Context, bg, fg color.NRGBA) layout.Dimensions {
		// 		if p.heartBtn.Clicked(gtx) {
		// 			p.favorited = !p.favorited
		// 		}
		// 		btn := component.SimpleIconButton(bg, fg, &p.heartBtn, icon.HeartIcon)
		// 		btn.Background = bg
		// 		if p.favorited {
		// 			btn.Color = color.NRGBA{R: 200, A: 255}
		// 		} else {
		// 			btn.Color = fg
		// 		}
		// 		return btn.Layout(gtx)
		// 	},
		// },
		component.SimpleIconAction(&p.plusBtn, icon.PlusIcon,
			component.OverflowAction{
				Name: "Create",
				Tag:  &p.plusBtn,
			},
		),
	}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{
		{
			Name: "Example 1",
			Tag:  &p.exampleOverflowState,
		},
		{
			Name: "Example 2",
			Tag:  &p.exampleOverflowState,
		},
	}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Home",
		Icon: icon.HomeIcon,
	}
}

// const (
// 	settingNameColumnWidth    = .3
// 	settingDetailsColumnWidth = 1 - settingNameColumnWidth
// )

func listItem(gtx layout.Context, th *material.Theme, item string) layout.Dimensions {
    return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
        layout.Rigid(func(gtx layout.Context) layout.Dimensions {
            return layout.Inset{Bottom: unit.Dp(1)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
                label := material.Body1(th, item)
                return label.Layout(gtx)
            })
        }),
        // layout.Rigid(func(gtx C) D {
		// 	gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(1))
		// 	gtx.Constraints.Max.Y = gtx.Constraints.Min.Y

		// 	dr := image.Rectangle{Max: gtx.Constraints.Min}
		// 	paint.LinearGradientOp{
		// 		Stop1:  layout.FPt(dr.Min),
		// 		Stop2:  layout.FPt(dr.Max),
		// 		Color1: color.NRGBA{R: 0x10, G: 0x10, B: 0xff, A: 0xFF},
		// 		Color2: color.NRGBA{R: 0x10, G: 0x10, B: 0xff, A: 0xFF},
		// 	}.Add(gtx.Ops)
		// 	defer clip.Rect(dr).Push(gtx.Ops).Pop()
		// 	paint.PaintOp{}.Add(gtx.Ops)
		// 	return layout.Dimensions{
		// 		Size: gtx.Constraints.Max,
		// 	}
		// }),
    )
}

func proxyList(gtx C,th *material.Theme, items []string) D {
	list := layout.List{Axis: layout.Vertical}
	return layout.Flex{
		// Vertical alignment, from top to bottom
		Axis: layout.Vertical,
		// Empty space is left at the start, i.e. at the top
		Spacing: layout.SpaceEnd,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(2), Width: unit.Dp(1)}
			return list.Layout(gtx, len(items), func(gtx C, i int) D {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx C) D {
					return border.Layout(gtx, func(gtx C) D {
						return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
							return listItem(gtx, th, items[i])
						})
					})
					// return material.Clickable(gtx, &widget.Clickable{}, func(gtx C) D {
					// 	return material.Body1(th, items[i]).Layout(gtx)
					// })
				})
			})
		}),
	)
}



func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	//var list = layout.List{Axis: layout.Vertical}
	items := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"}
	if p.startButton.Clicked(gtx) {
		p.boiling = !p.boiling
		fmt.Println("button is clicked")
	}
	return layout.Flex{
		// Vertical alignment, from top to bottom
		Axis: layout.Vertical,
		// Empty space is left at the start, i.e. at the top
		//Spacing: layout.SpaceStart,
	}.Layout(gtx,
		// insert a list of items
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return proxyList(gtx, th, items)
		}),
		layout.Rigid(
			func(gtx C) D {
				bar := material.ProgressBar(th, 0.1)
				return bar.Layout(gtx)
			},
		),
		layout.Rigid(
			func(gtx C) D {
				// We start by defining a set of margins
				margins := layout.Inset{
					Top:    unit.Dp(25),
					Bottom: unit.Dp(25),
					Right:  unit.Dp(35),
					Left:   unit.Dp(35),
				}
				// Then we lay out within those margins ...
				return margins.Layout(gtx,
					// ...the same function we earlier used to create a button
					func(gtx C) D {
						var text string
						if !p.boiling {
							text = "Connect"
						} else {
							text = "Disconnect"
						}
						btn := material.Button(th, &p.startButton, text)
						return btn.Layout(gtx)
					},
				)
			},
		),
	)
}
