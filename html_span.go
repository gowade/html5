package h

type htmlSpan struct {
	htmlElement
}

func Span() *htmlSpan {
	e := &htmlSpan{}
	e.a = make(map[string]interface{})
	e.tagName = "span"
	return e
}

func (e *htmlSpan) S(style StyleMap) *htmlSpan {
	e.htmlElement.S(style)
	return e
}

func (e *htmlSpan) Key(key interface{}) *htmlSpan {
	e.key = F(key)
	return e
}

func (e *htmlSpan) ID(v string) *htmlSpan {
	e.a["id"] = v
	return e
}

func (e *htmlSpan) Class(v string) *htmlSpan {
	e.a["class"] = v
	return e
}

func (e *htmlSpan) Title(v string) *htmlSpan {
	e.a["title"] = v
	return e
}

func (e *htmlSpan) Lang(v string) *htmlSpan {
	e.a["lang"] = v
	return e
}

func (e *htmlSpan) Translate(v bool) *htmlSpan {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlSpan) Dir(v string) *htmlSpan {
	e.a["dir"] = v
	return e
}

func (e *htmlSpan) Hidden(v bool) *htmlSpan {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlSpan) TabIndex(v int) *htmlSpan {
	e.a["tabindex"] = v
	return e
}

func (e *htmlSpan) AccessKey(v string) *htmlSpan {
	e.a["accesskey"] = v
	return e
}

func (e *htmlSpan) Draggable(v bool) *htmlSpan {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlSpan) Spellcheck(v bool) *htmlSpan {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
