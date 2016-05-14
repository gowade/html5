package h

type htmlMenu struct {
	htmlElement
}

func Menu() *htmlMenu {
	e := &htmlMenu{}
	e.a = make(map[string]interface{})
	e.tagName = "menu"
	return e
}

func (e *htmlMenu) S(style StyleMap) *htmlMenu {
	e.htmlElement.S(style)
	return e
}

func (e *htmlMenu) Key(key interface{}) *htmlMenu {
	e.key = F(key)
	return e
}

func (e *htmlMenu) Type(v string) *htmlMenu {
	e.a["type"] = v
	return e
}

func (e *htmlMenu) Label(v string) *htmlMenu {
	e.a["label"] = v
	return e
}

func (e *htmlMenu) ID(v string) *htmlMenu {
	e.a["id"] = v
	return e
}

func (e *htmlMenu) Class(v string) *htmlMenu {
	e.a["class"] = v
	return e
}

func (e *htmlMenu) Title(v string) *htmlMenu {
	e.a["title"] = v
	return e
}

func (e *htmlMenu) Lang(v string) *htmlMenu {
	e.a["lang"] = v
	return e
}

func (e *htmlMenu) Translate(v bool) *htmlMenu {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlMenu) Dir(v string) *htmlMenu {
	e.a["dir"] = v
	return e
}

func (e *htmlMenu) Hidden(v bool) *htmlMenu {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlMenu) TabIndex(v int) *htmlMenu {
	e.a["tabindex"] = v
	return e
}

func (e *htmlMenu) AccessKey(v string) *htmlMenu {
	e.a["accesskey"] = v
	return e
}

func (e *htmlMenu) Draggable(v bool) *htmlMenu {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlMenu) Spellcheck(v bool) *htmlMenu {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
