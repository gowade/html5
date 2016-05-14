package h

type htmlStyle struct {
	htmlElement
}

func Style() *htmlStyle {
	e := &htmlStyle{}
	e.a = make(map[string]interface{})
	e.tagName = "style"
	return e
}

func (e *htmlStyle) S(style StyleMap) *htmlStyle {
	e.htmlElement.S(style)
	return e
}

func (e *htmlStyle) Key(key interface{}) *htmlStyle {
	e.key = F(key)
	return e
}

func (e *htmlStyle) Media(v string) *htmlStyle {
	e.a["media"] = v
	return e
}

func (e *htmlStyle) Nonce(v string) *htmlStyle {
	e.a["nonce"] = v
	return e
}

func (e *htmlStyle) Type(v string) *htmlStyle {
	e.a["type"] = v
	return e
}

func (e *htmlStyle) ID(v string) *htmlStyle {
	e.a["id"] = v
	return e
}

func (e *htmlStyle) Class(v string) *htmlStyle {
	e.a["class"] = v
	return e
}

func (e *htmlStyle) Title(v string) *htmlStyle {
	e.a["title"] = v
	return e
}

func (e *htmlStyle) Lang(v string) *htmlStyle {
	e.a["lang"] = v
	return e
}

func (e *htmlStyle) Translate(v bool) *htmlStyle {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlStyle) Dir(v string) *htmlStyle {
	e.a["dir"] = v
	return e
}

func (e *htmlStyle) Hidden(v bool) *htmlStyle {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlStyle) TabIndex(v int) *htmlStyle {
	e.a["tabindex"] = v
	return e
}

func (e *htmlStyle) AccessKey(v string) *htmlStyle {
	e.a["accesskey"] = v
	return e
}

func (e *htmlStyle) Draggable(v bool) *htmlStyle {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlStyle) Spellcheck(v bool) *htmlStyle {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
