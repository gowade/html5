package h

type htmlKeygen struct {
	htmlElement
}

func Keygen() *htmlKeygen {
	e := &htmlKeygen{}
	e.a = make(map[string]interface{})
	e.tagName = "keygen"
	return e
}

func (e *htmlKeygen) S(style StyleMap) *htmlKeygen {
	e.htmlElement.S(style)
	return e
}

func (e *htmlKeygen) Key(key interface{}) *htmlKeygen {
	e.key = F(key)
	return e
}

func (e *htmlKeygen) Autofocus(v bool) *htmlKeygen {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

func (e *htmlKeygen) Challenge(v string) *htmlKeygen {
	e.a["challenge"] = v
	return e
}

func (e *htmlKeygen) Disabled(v bool) *htmlKeygen {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlKeygen) Keytype(v string) *htmlKeygen {
	e.a["keytype"] = v
	return e
}

func (e *htmlKeygen) Name(v string) *htmlKeygen {
	e.a["name"] = v
	return e
}

func (e *htmlKeygen) CheckValidity(v bool) *htmlKeygen {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlKeygen) ReportValidity(v bool) *htmlKeygen {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlKeygen) ID(v string) *htmlKeygen {
	e.a["id"] = v
	return e
}

func (e *htmlKeygen) Class(v string) *htmlKeygen {
	e.a["class"] = v
	return e
}

func (e *htmlKeygen) Title(v string) *htmlKeygen {
	e.a["title"] = v
	return e
}

func (e *htmlKeygen) Lang(v string) *htmlKeygen {
	e.a["lang"] = v
	return e
}

func (e *htmlKeygen) Translate(v bool) *htmlKeygen {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlKeygen) Dir(v string) *htmlKeygen {
	e.a["dir"] = v
	return e
}

func (e *htmlKeygen) Hidden(v bool) *htmlKeygen {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlKeygen) TabIndex(v int) *htmlKeygen {
	e.a["tabindex"] = v
	return e
}

func (e *htmlKeygen) AccessKey(v string) *htmlKeygen {
	e.a["accesskey"] = v
	return e
}

func (e *htmlKeygen) Draggable(v bool) *htmlKeygen {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlKeygen) Spellcheck(v bool) *htmlKeygen {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
