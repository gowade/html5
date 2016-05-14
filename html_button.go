package h

type htmlButton struct {
	htmlElement
}

func Button() *htmlButton {
	e := &htmlButton{}
	e.a = make(map[string]interface{})
	e.tagName = "button"
	return e
}

func (e *htmlButton) S(style StyleMap) *htmlButton {
	e.htmlElement.S(style)
	return e
}

func (e *htmlButton) Key(key interface{}) *htmlButton {
	e.key = F(key)
	return e
}

func (e *htmlButton) Autofocus(v bool) *htmlButton {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

func (e *htmlButton) Disabled(v bool) *htmlButton {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlButton) FormAction(v string) *htmlButton {
	e.a["formaction"] = v
	return e
}

func (e *htmlButton) FormEnctype(v string) *htmlButton {
	e.a["formenctype"] = v
	return e
}

func (e *htmlButton) FormMethod(v string) *htmlButton {
	e.a["formmethod"] = v
	return e
}

func (e *htmlButton) FormNoValidate(v bool) *htmlButton {
	if v {
		e.a["formnovalidate"] = ""
	} else {
		delete(e.a, "formnovalidate")
	}
	return e
}

func (e *htmlButton) FormTarget(v string) *htmlButton {
	e.a["formtarget"] = v
	return e
}

func (e *htmlButton) Name(v string) *htmlButton {
	e.a["name"] = v
	return e
}

func (e *htmlButton) Type(v string) *htmlButton {
	e.a["type"] = v
	return e
}

func (e *htmlButton) Value(v string) *htmlButton {
	e.a["value"] = v
	return e
}

func (e *htmlButton) CheckValidity(v bool) *htmlButton {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlButton) ReportValidity(v bool) *htmlButton {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlButton) ID(v string) *htmlButton {
	e.a["id"] = v
	return e
}

func (e *htmlButton) Class(v string) *htmlButton {
	e.a["class"] = v
	return e
}

func (e *htmlButton) Title(v string) *htmlButton {
	e.a["title"] = v
	return e
}

func (e *htmlButton) Lang(v string) *htmlButton {
	e.a["lang"] = v
	return e
}

func (e *htmlButton) Translate(v bool) *htmlButton {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlButton) Dir(v string) *htmlButton {
	e.a["dir"] = v
	return e
}

func (e *htmlButton) Hidden(v bool) *htmlButton {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlButton) TabIndex(v int) *htmlButton {
	e.a["tabindex"] = v
	return e
}

func (e *htmlButton) AccessKey(v string) *htmlButton {
	e.a["accesskey"] = v
	return e
}

func (e *htmlButton) Draggable(v bool) *htmlButton {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlButton) Spellcheck(v bool) *htmlButton {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
