package h

type htmlHeading struct {
	htmlElement
}

func Heading() *htmlHeading {
	e := &htmlHeading{}
	e.a = make(map[string]interface{})
	e.tagName = "heading"
	return e
}

func (e *htmlHeading) S(style StyleMap) *htmlHeading {
	e.htmlElement.S(style)
	return e
}

func (e *htmlHeading) Key(key interface{}) *htmlHeading {
	e.key = F(key)
	return e
}

func (e *htmlHeading) ID(v string) *htmlHeading {
	e.a["id"] = v
	return e
}

func (e *htmlHeading) Class(v string) *htmlHeading {
	e.a["class"] = v
	return e
}

func (e *htmlHeading) Title(v string) *htmlHeading {
	e.a["title"] = v
	return e
}

func (e *htmlHeading) Lang(v string) *htmlHeading {
	e.a["lang"] = v
	return e
}

func (e *htmlHeading) Translate(v bool) *htmlHeading {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlHeading) Dir(v string) *htmlHeading {
	e.a["dir"] = v
	return e
}

func (e *htmlHeading) Hidden(v bool) *htmlHeading {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlHeading) TabIndex(v int) *htmlHeading {
	e.a["tabindex"] = v
	return e
}

func (e *htmlHeading) AccessKey(v string) *htmlHeading {
	e.a["accesskey"] = v
	return e
}

func (e *htmlHeading) Draggable(v bool) *htmlHeading {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlHeading) Spellcheck(v bool) *htmlHeading {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
