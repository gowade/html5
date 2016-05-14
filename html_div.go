package h

type htmlDiv struct {
	htmlElement
}

func Div() *htmlDiv {
	e := &htmlDiv{}
	e.a = make(map[string]interface{})
	e.tagName = "div"
	return e
}

func (e *htmlDiv) S(style StyleMap) *htmlDiv {
	e.htmlElement.S(style)
	return e
}

func (e *htmlDiv) Key(key interface{}) *htmlDiv {
	e.key = F(key)
	return e
}

func (e *htmlDiv) ID(v string) *htmlDiv {
	e.a["id"] = v
	return e
}

func (e *htmlDiv) Class(v string) *htmlDiv {
	e.a["class"] = v
	return e
}

func (e *htmlDiv) Title(v string) *htmlDiv {
	e.a["title"] = v
	return e
}

func (e *htmlDiv) Lang(v string) *htmlDiv {
	e.a["lang"] = v
	return e
}

func (e *htmlDiv) Translate(v bool) *htmlDiv {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlDiv) Dir(v string) *htmlDiv {
	e.a["dir"] = v
	return e
}

func (e *htmlDiv) Hidden(v bool) *htmlDiv {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlDiv) TabIndex(v int) *htmlDiv {
	e.a["tabindex"] = v
	return e
}

func (e *htmlDiv) AccessKey(v string) *htmlDiv {
	e.a["accesskey"] = v
	return e
}

func (e *htmlDiv) Draggable(v bool) *htmlDiv {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlDiv) Spellcheck(v bool) *htmlDiv {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
