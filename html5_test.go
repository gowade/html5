package html5_test

import (
	"testing"

	h "github.com/gowade/html5"
)

type ToggleAll struct{}
type ToggleOne struct{ i int }
type EditOne struct{ i int }
type RemoveOne struct{ i int }

type Todo struct {
	Active  bool
	Title   string
	Content string
	El      h.DOMElement
}

type App struct {
	Todos []Todo
}

func (t App) Render() h.Node {
	return h.Div().C(
		h.Div().ID("main").Hidden(len(t.Todos) == 0).C(
			h.Input().ID("toggle-all").Type("checkbox").OnClick(&ToggleAll{}),
			h.Label().For("toggle-all").T("Mark all as complete"),
			h.UL().ID("todo-list").L(func(l *h.L) {
				for i, todo := range t.Todos {
					if !todo.Active {
						continue
					}

					l.LI().C(
						h.Div().Class("view").C(
							h.Input().
								Class("toggle").
								Type("checkbox").
								OnClick(&ToggleOne{i}),
							h.Label().OnDblClick(&EditOne{i}).T(todo.Title),
							h.Button().Class("destroy").OnClick(&RemoveOne{i}),
						),
						h.Form().C(
							h.Input().Class("edit").
								Value(todo.Content).
								Ref(&todo.El),
						),
					)
				}
			}),
		),
	)
}

func TestHInterface(t *testing.T) {
}
