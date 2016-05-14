package h

type htmlOL struct {
	htmlElement
}

func OL() *htmlOL {
	e := &htmlOL{}
	e.a = make(map[string]interface{})
	e.tagName = "ol"
	return e
}

func (e *htmlOL) S(style StyleMap) *htmlOL {
	e.htmlElement.S(style)
	return e
}

func (e *htmlOL) Key(key interface{}) *htmlOL {
	e.key = F(key)
	return e
}

func (e *htmlOL) Reversed(v bool) *htmlOL {
	if v {
		e.a["reversed"] = ""
	} else {
		delete(e.a, "reversed")
	}
	return e
}

func (e *htmlOL) Start(v int) *htmlOL {
	e.a["start"] = v
	return e
}

func (e *htmlOL) Type(v string) *htmlOL {
	e.a["type"] = v
	return e
}

func (e *htmlOL) ID(v string) *htmlOL {
	e.a["id"] = v
	return e
}

func (e *htmlOL) Class(v string) *htmlOL {
	e.a["class"] = v
	return e
}

func (e *htmlOL) Title(v string) *htmlOL {
	e.a["title"] = v
	return e
}

func (e *htmlOL) Lang(v string) *htmlOL {
	e.a["lang"] = v
	return e
}

func (e *htmlOL) Translate(v bool) *htmlOL {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlOL) Dir(v string) *htmlOL {
	e.a["dir"] = v
	return e
}

func (e *htmlOL) Hidden(v bool) *htmlOL {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlOL) TabIndex(v int) *htmlOL {
	e.a["tabindex"] = v
	return e
}

func (e *htmlOL) AccessKey(v string) *htmlOL {
	e.a["accesskey"] = v
	return e
}

func (e *htmlOL) Draggable(v bool) *htmlOL {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlOL) Spellcheck(v bool) *htmlOL {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
