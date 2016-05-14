package h

type htmlLabel struct {
	htmlElement
}

func Label() *htmlLabel {
	e := &htmlLabel{}
	e.a = make(map[string]interface{})
	e.tagName = "label"
	return e
}

func (e *htmlLabel) S(style StyleMap) *htmlLabel {
	e.htmlElement.S(style)
	return e
}

func (e *htmlLabel) Key(key interface{}) *htmlLabel {
	e.key = F(key)
	return e
}

func (e *htmlLabel) For(v string) *htmlLabel {
	e.a["htmlfor"] = v
	return e
}

func (e *htmlLabel) ID(v string) *htmlLabel {
	e.a["id"] = v
	return e
}

func (e *htmlLabel) Class(v string) *htmlLabel {
	e.a["class"] = v
	return e
}

func (e *htmlLabel) Title(v string) *htmlLabel {
	e.a["title"] = v
	return e
}

func (e *htmlLabel) Lang(v string) *htmlLabel {
	e.a["lang"] = v
	return e
}

func (e *htmlLabel) Translate(v bool) *htmlLabel {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlLabel) Dir(v string) *htmlLabel {
	e.a["dir"] = v
	return e
}

func (e *htmlLabel) Hidden(v bool) *htmlLabel {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlLabel) TabIndex(v int) *htmlLabel {
	e.a["tabindex"] = v
	return e
}

func (e *htmlLabel) AccessKey(v string) *htmlLabel {
	e.a["accesskey"] = v
	return e
}

func (e *htmlLabel) Draggable(v bool) *htmlLabel {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlLabel) Spellcheck(v bool) *htmlLabel {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
