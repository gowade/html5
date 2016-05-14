package h

type htmlPre struct {
	htmlElement
}

func Pre() *htmlPre {
	e := &htmlPre{}
	e.a = make(map[string]interface{})
	e.tagName = "pre"
	return e
}

func (e *htmlPre) S(style StyleMap) *htmlPre {
	e.htmlElement.S(style)
	return e
}

func (e *htmlPre) Key(key interface{}) *htmlPre {
	e.key = F(key)
	return e
}

func (e *htmlPre) ID(v string) *htmlPre {
	e.a["id"] = v
	return e
}

func (e *htmlPre) Class(v string) *htmlPre {
	e.a["class"] = v
	return e
}

func (e *htmlPre) Title(v string) *htmlPre {
	e.a["title"] = v
	return e
}

func (e *htmlPre) Lang(v string) *htmlPre {
	e.a["lang"] = v
	return e
}

func (e *htmlPre) Translate(v bool) *htmlPre {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlPre) Dir(v string) *htmlPre {
	e.a["dir"] = v
	return e
}

func (e *htmlPre) Hidden(v bool) *htmlPre {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlPre) TabIndex(v int) *htmlPre {
	e.a["tabindex"] = v
	return e
}

func (e *htmlPre) AccessKey(v string) *htmlPre {
	e.a["accesskey"] = v
	return e
}

func (e *htmlPre) Draggable(v bool) *htmlPre {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlPre) Spellcheck(v bool) *htmlPre {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
