package h

type htmlHR struct {
	htmlElement
}

func HR() *htmlHR {
	e := &htmlHR{}
	e.a = make(map[string]interface{})
	e.tagName = "hr"
	return e
}

func (e *htmlHR) S(style StyleMap) *htmlHR {
	e.htmlElement.S(style)
	return e
}

func (e *htmlHR) Key(key interface{}) *htmlHR {
	e.key = F(key)
	return e
}

func (e *htmlHR) ID(v string) *htmlHR {
	e.a["id"] = v
	return e
}

func (e *htmlHR) Class(v string) *htmlHR {
	e.a["class"] = v
	return e
}

func (e *htmlHR) Title(v string) *htmlHR {
	e.a["title"] = v
	return e
}

func (e *htmlHR) Lang(v string) *htmlHR {
	e.a["lang"] = v
	return e
}

func (e *htmlHR) Translate(v bool) *htmlHR {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlHR) Dir(v string) *htmlHR {
	e.a["dir"] = v
	return e
}

func (e *htmlHR) Hidden(v bool) *htmlHR {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlHR) TabIndex(v int) *htmlHR {
	e.a["tabindex"] = v
	return e
}

func (e *htmlHR) AccessKey(v string) *htmlHR {
	e.a["accesskey"] = v
	return e
}

func (e *htmlHR) Draggable(v bool) *htmlHR {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlHR) Spellcheck(v bool) *htmlHR {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
