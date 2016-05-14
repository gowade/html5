package h

type htmlDetails struct {
	htmlElement
}

func Details() *htmlDetails {
	e := &htmlDetails{}
	e.a = make(map[string]interface{})
	e.tagName = "details"
	return e
}

func (e *htmlDetails) S(style StyleMap) *htmlDetails {
	e.htmlElement.S(style)
	return e
}

func (e *htmlDetails) Key(key interface{}) *htmlDetails {
	e.key = F(key)
	return e
}

func (e *htmlDetails) Open(v bool) *htmlDetails {
	if v {
		e.a["open"] = ""
	} else {
		delete(e.a, "open")
	}
	return e
}

func (e *htmlDetails) ID(v string) *htmlDetails {
	e.a["id"] = v
	return e
}

func (e *htmlDetails) Class(v string) *htmlDetails {
	e.a["class"] = v
	return e
}

func (e *htmlDetails) Title(v string) *htmlDetails {
	e.a["title"] = v
	return e
}

func (e *htmlDetails) Lang(v string) *htmlDetails {
	e.a["lang"] = v
	return e
}

func (e *htmlDetails) Translate(v bool) *htmlDetails {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlDetails) Dir(v string) *htmlDetails {
	e.a["dir"] = v
	return e
}

func (e *htmlDetails) Hidden(v bool) *htmlDetails {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlDetails) TabIndex(v int) *htmlDetails {
	e.a["tabindex"] = v
	return e
}

func (e *htmlDetails) AccessKey(v string) *htmlDetails {
	e.a["accesskey"] = v
	return e
}

func (e *htmlDetails) Draggable(v bool) *htmlDetails {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlDetails) Spellcheck(v bool) *htmlDetails {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
