package h

type htmlTableCol struct {
	htmlElement
}

func TableCol() *htmlTableCol {
	e := &htmlTableCol{}
	e.a = make(map[string]interface{})
	e.tagName = "tablecol"
	return e
}

func (e *htmlTableCol) S(style StyleMap) *htmlTableCol {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTableCol) Key(key interface{}) *htmlTableCol {
	e.key = F(key)
	return e
}

func (e *htmlTableCol) Span(v int) *htmlTableCol {
	e.a["span"] = v
	return e
}

func (e *htmlTableCol) ID(v string) *htmlTableCol {
	e.a["id"] = v
	return e
}

func (e *htmlTableCol) Class(v string) *htmlTableCol {
	e.a["class"] = v
	return e
}

func (e *htmlTableCol) Title(v string) *htmlTableCol {
	e.a["title"] = v
	return e
}

func (e *htmlTableCol) Lang(v string) *htmlTableCol {
	e.a["lang"] = v
	return e
}

func (e *htmlTableCol) Translate(v bool) *htmlTableCol {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTableCol) Dir(v string) *htmlTableCol {
	e.a["dir"] = v
	return e
}

func (e *htmlTableCol) Hidden(v bool) *htmlTableCol {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTableCol) TabIndex(v int) *htmlTableCol {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTableCol) AccessKey(v string) *htmlTableCol {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTableCol) Draggable(v bool) *htmlTableCol {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTableCol) Spellcheck(v bool) *htmlTableCol {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
