package h

type htmlUL struct {
	htmlElement
}

func UL() *htmlUL {
	e := &htmlUL{}
	e.a = make(map[string]interface{})
	e.tagName = "ul"
	return e
}

func (e *htmlUL) S(style StyleMap) *htmlUL {
	e.htmlElement.S(style)
	return e
}

func (e *htmlUL) Key(key interface{}) *htmlUL {
	e.key = F(key)
	return e
}

func (e *htmlUL) ID(v string) *htmlUL {
	e.a["id"] = v
	return e
}

func (e *htmlUL) Class(v string) *htmlUL {
	e.a["class"] = v
	return e
}

func (e *htmlUL) Title(v string) *htmlUL {
	e.a["title"] = v
	return e
}

func (e *htmlUL) Lang(v string) *htmlUL {
	e.a["lang"] = v
	return e
}

func (e *htmlUL) Translate(v bool) *htmlUL {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlUL) Dir(v string) *htmlUL {
	e.a["dir"] = v
	return e
}

func (e *htmlUL) Hidden(v bool) *htmlUL {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlUL) TabIndex(v int) *htmlUL {
	e.a["tabindex"] = v
	return e
}

func (e *htmlUL) AccessKey(v string) *htmlUL {
	e.a["accesskey"] = v
	return e
}

func (e *htmlUL) Draggable(v bool) *htmlUL {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlUL) Spellcheck(v bool) *htmlUL {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
