package h_test

import (
	"testing"

	"github.com/gowade/wade/h"
)

type ToggleAll struct{}
type ToggleOne struct{ i int }
type EditOne struct{ i int }
type RemoveOne struct{ i int }

type Todo struct {
	Active  bool
	Title   string
	Content string
}

type App struct {
	Todos []Todo
}

func (t App) Render() h.Node {
	return h.Div().C(
		h.Div().ID("main").Hidden(len(t.Todos) == 0).C(
			h.Input().ID("toggle-all").Type(h.InputCheckBox).OnClick(ToggleAll{}),
			h.Label().For("toggle-all").T("Mark all as complete"),
			h.UL().ID("todo-list").L(func(l *h.L) {
				for i, todo := range t.Todos {
					if todo.Active {
						l.Add(h.LI().C(
							h.Div().Class("view").C(
								h.Input().
									Class("toggle").
									Type(h.InputCheckBox).
									OnClick(ToggleOne{i}),
								h.Label().OnDoubleClick(EditOne{i}).T(todo.Title),
								h.Button().Class("destroy").OnClick(RemoveOne{i}),
							),
							h.Form().C(
								h.Input().Class("edit").
									Value(todo.Content).
									StringRef(&todo.Content),
							),
						))
					}
				}
			}),
		),
	)
}

func TestHInterface(t *testing.T) {
}
