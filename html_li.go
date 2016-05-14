package h

type htmlLI struct {
	htmlElement
}

func LI() *htmlLI {
	e := &htmlLI{}
	e.a = make(map[string]interface{})
	e.tagName = "li"
	return e
}

func (e *htmlLI) S(style StyleMap) *htmlLI {
	e.htmlElement.S(style)
	return e
}

func (e *htmlLI) Key(key interface{}) *htmlLI {
	e.key = F(key)
	return e
}

func (e *htmlLI) Value(v int) *htmlLI {
	e.a["value"] = v
	return e
}

func (e *htmlLI) ID(v string) *htmlLI {
	e.a["id"] = v
	return e
}

func (e *htmlLI) Class(v string) *htmlLI {
	e.a["class"] = v
	return e
}

func (e *htmlLI) Title(v string) *htmlLI {
	e.a["title"] = v
	return e
}

func (e *htmlLI) Lang(v string) *htmlLI {
	e.a["lang"] = v
	return e
}

func (e *htmlLI) Translate(v bool) *htmlLI {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlLI) Dir(v string) *htmlLI {
	e.a["dir"] = v
	return e
}

func (e *htmlLI) Hidden(v bool) *htmlLI {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlLI) TabIndex(v int) *htmlLI {
	e.a["tabindex"] = v
	return e
}

func (e *htmlLI) AccessKey(v string) *htmlLI {
	e.a["accesskey"] = v
	return e
}

func (e *htmlLI) Draggable(v bool) *htmlLI {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlLI) Spellcheck(v bool) *htmlLI {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
