package h

type htmlOptGroup struct {
	htmlElement
}

func OptGroup() *htmlOptGroup {
	e := &htmlOptGroup{}
	e.a = make(map[string]interface{})
	e.tagName = "optgroup"
	return e
}

func (e *htmlOptGroup) S(style StyleMap) *htmlOptGroup {
	e.htmlElement.S(style)
	return e
}

func (e *htmlOptGroup) Key(key interface{}) *htmlOptGroup {
	e.key = F(key)
	return e
}

func (e *htmlOptGroup) Disabled(v bool) *htmlOptGroup {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlOptGroup) Label(v string) *htmlOptGroup {
	e.a["label"] = v
	return e
}

func (e *htmlOptGroup) ID(v string) *htmlOptGroup {
	e.a["id"] = v
	return e
}

func (e *htmlOptGroup) Class(v string) *htmlOptGroup {
	e.a["class"] = v
	return e
}

func (e *htmlOptGroup) Title(v string) *htmlOptGroup {
	e.a["title"] = v
	return e
}

func (e *htmlOptGroup) Lang(v string) *htmlOptGroup {
	e.a["lang"] = v
	return e
}

func (e *htmlOptGroup) Translate(v bool) *htmlOptGroup {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlOptGroup) Dir(v string) *htmlOptGroup {
	e.a["dir"] = v
	return e
}

func (e *htmlOptGroup) Hidden(v bool) *htmlOptGroup {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlOptGroup) TabIndex(v int) *htmlOptGroup {
	e.a["tabindex"] = v
	return e
}

func (e *htmlOptGroup) AccessKey(v string) *htmlOptGroup {
	e.a["accesskey"] = v
	return e
}

func (e *htmlOptGroup) Draggable(v bool) *htmlOptGroup {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlOptGroup) Spellcheck(v bool) *htmlOptGroup {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
