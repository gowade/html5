package h

type htmlMenuItem struct {
	htmlElement
}

func MenuItem() *htmlMenuItem {
	e := &htmlMenuItem{}
	e.a = make(map[string]interface{})
	e.tagName = "menuitem"
	return e
}

func (e *htmlMenuItem) S(style StyleMap) *htmlMenuItem {
	e.htmlElement.S(style)
	return e
}

func (e *htmlMenuItem) Key(key interface{}) *htmlMenuItem {
	e.key = F(key)
	return e
}

func (e *htmlMenuItem) Type(v string) *htmlMenuItem {
	e.a["type"] = v
	return e
}

func (e *htmlMenuItem) Label(v string) *htmlMenuItem {
	e.a["label"] = v
	return e
}

func (e *htmlMenuItem) Icon(v string) *htmlMenuItem {
	e.a["icon"] = v
	return e
}

func (e *htmlMenuItem) Disabled(v bool) *htmlMenuItem {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlMenuItem) Checked(v bool) *htmlMenuItem {
	if v {
		e.a["checked"] = ""
	} else {
		delete(e.a, "checked")
	}
	return e
}

func (e *htmlMenuItem) Radiogroup(v string) *htmlMenuItem {
	e.a["radiogroup"] = v
	return e
}

func (e *htmlMenuItem) Default(v bool) *htmlMenuItem {
	if v {
		e.a["default"] = ""
	} else {
		delete(e.a, "default")
	}
	return e
}

func (e *htmlMenuItem) ID(v string) *htmlMenuItem {
	e.a["id"] = v
	return e
}

func (e *htmlMenuItem) Class(v string) *htmlMenuItem {
	e.a["class"] = v
	return e
}

func (e *htmlMenuItem) Title(v string) *htmlMenuItem {
	e.a["title"] = v
	return e
}

func (e *htmlMenuItem) Lang(v string) *htmlMenuItem {
	e.a["lang"] = v
	return e
}

func (e *htmlMenuItem) Translate(v bool) *htmlMenuItem {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlMenuItem) Dir(v string) *htmlMenuItem {
	e.a["dir"] = v
	return e
}

func (e *htmlMenuItem) Hidden(v bool) *htmlMenuItem {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlMenuItem) TabIndex(v int) *htmlMenuItem {
	e.a["tabindex"] = v
	return e
}

func (e *htmlMenuItem) AccessKey(v string) *htmlMenuItem {
	e.a["accesskey"] = v
	return e
}

func (e *htmlMenuItem) Draggable(v bool) *htmlMenuItem {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlMenuItem) Spellcheck(v bool) *htmlMenuItem {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
