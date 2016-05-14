package h

type htmlBR struct {
	htmlElement
}

func BR() *htmlBR {
	e := &htmlBR{}
	e.a = make(map[string]interface{})
	e.tagName = "br"
	return e
}

func (e *htmlBR) S(style StyleMap) *htmlBR {
	e.htmlElement.S(style)
	return e
}

func (e *htmlBR) Key(key interface{}) *htmlBR {
	e.key = F(key)
	return e
}

func (e *htmlBR) ID(v string) *htmlBR {
	e.a["id"] = v
	return e
}

func (e *htmlBR) Class(v string) *htmlBR {
	e.a["class"] = v
	return e
}

func (e *htmlBR) Title(v string) *htmlBR {
	e.a["title"] = v
	return e
}

func (e *htmlBR) Lang(v string) *htmlBR {
	e.a["lang"] = v
	return e
}

func (e *htmlBR) Translate(v bool) *htmlBR {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlBR) Dir(v string) *htmlBR {
	e.a["dir"] = v
	return e
}

func (e *htmlBR) Hidden(v bool) *htmlBR {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlBR) TabIndex(v int) *htmlBR {
	e.a["tabindex"] = v
	return e
}

func (e *htmlBR) AccessKey(v string) *htmlBR {
	e.a["accesskey"] = v
	return e
}

func (e *htmlBR) Draggable(v bool) *htmlBR {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlBR) Spellcheck(v bool) *htmlBR {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
