package h

type htmlHead struct {
	htmlElement
}

func Head() *htmlHead {
	e := &htmlHead{}
	e.a = make(map[string]interface{})
	e.tagName = "head"
	return e
}

func (e *htmlHead) S(style StyleMap) *htmlHead {
	e.htmlElement.S(style)
	return e
}

func (e *htmlHead) Key(key interface{}) *htmlHead {
	e.key = F(key)
	return e
}

func (e *htmlHead) ID(v string) *htmlHead {
	e.a["id"] = v
	return e
}

func (e *htmlHead) Class(v string) *htmlHead {
	e.a["class"] = v
	return e
}

func (e *htmlHead) Title(v string) *htmlHead {
	e.a["title"] = v
	return e
}

func (e *htmlHead) Lang(v string) *htmlHead {
	e.a["lang"] = v
	return e
}

func (e *htmlHead) Translate(v bool) *htmlHead {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlHead) Dir(v string) *htmlHead {
	e.a["dir"] = v
	return e
}

func (e *htmlHead) Hidden(v bool) *htmlHead {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlHead) TabIndex(v int) *htmlHead {
	e.a["tabindex"] = v
	return e
}

func (e *htmlHead) AccessKey(v string) *htmlHead {
	e.a["accesskey"] = v
	return e
}

func (e *htmlHead) Draggable(v bool) *htmlHead {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlHead) Spellcheck(v bool) *htmlHead {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
