package h

type htmlOption struct {
	htmlElement
}

func Option() *htmlOption {
	e := &htmlOption{}
	e.a = make(map[string]interface{})
	e.tagName = "option"
	return e
}

func (e *htmlOption) S(style StyleMap) *htmlOption {
	e.htmlElement.S(style)
	return e
}

func (e *htmlOption) Key(key interface{}) *htmlOption {
	e.key = F(key)
	return e
}

func (e *htmlOption) Disabled(v bool) *htmlOption {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlOption) Label(v string) *htmlOption {
	e.a["label"] = v
	return e
}

func (e *htmlOption) DefaultSelected(v bool) *htmlOption {
	if v {
		e.a["defaultselected"] = ""
	} else {
		delete(e.a, "defaultselected")
	}
	return e
}

func (e *htmlOption) Selected(v bool) *htmlOption {
	if v {
		e.a["selected"] = ""
	} else {
		delete(e.a, "selected")
	}
	return e
}

func (e *htmlOption) Value(v string) *htmlOption {
	e.a["value"] = v
	return e
}

func (e *htmlOption) Text(v string) *htmlOption {
	e.a["text"] = v
	return e
}

func (e *htmlOption) ID(v string) *htmlOption {
	e.a["id"] = v
	return e
}

func (e *htmlOption) Class(v string) *htmlOption {
	e.a["class"] = v
	return e
}

func (e *htmlOption) Title(v string) *htmlOption {
	e.a["title"] = v
	return e
}

func (e *htmlOption) Lang(v string) *htmlOption {
	e.a["lang"] = v
	return e
}

func (e *htmlOption) Translate(v bool) *htmlOption {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlOption) Dir(v string) *htmlOption {
	e.a["dir"] = v
	return e
}

func (e *htmlOption) Hidden(v bool) *htmlOption {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlOption) TabIndex(v int) *htmlOption {
	e.a["tabindex"] = v
	return e
}

func (e *htmlOption) AccessKey(v string) *htmlOption {
	e.a["accesskey"] = v
	return e
}

func (e *htmlOption) Draggable(v bool) *htmlOption {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlOption) Spellcheck(v bool) *htmlOption {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
