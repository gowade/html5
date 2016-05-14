package h

type htmlForm struct {
	htmlElement
}

func Form() *htmlForm {
	e := &htmlForm{}
	e.a = make(map[string]interface{})
	e.tagName = "form"
	return e
}

func (e *htmlForm) S(style StyleMap) *htmlForm {
	e.htmlElement.S(style)
	return e
}

func (e *htmlForm) Key(key interface{}) *htmlForm {
	e.key = F(key)
	return e
}

func (e *htmlForm) AcceptCharset(v string) *htmlForm {
	e.a["acceptcharset"] = v
	return e
}

func (e *htmlForm) Action(v string) *htmlForm {
	e.a["action"] = v
	return e
}

func (e *htmlForm) Autocomplete(v string) *htmlForm {
	e.a["autocomplete"] = v
	return e
}

func (e *htmlForm) Enctype(v string) *htmlForm {
	e.a["enctype"] = v
	return e
}

func (e *htmlForm) Encoding(v string) *htmlForm {
	e.a["encoding"] = v
	return e
}

func (e *htmlForm) Method(v string) *htmlForm {
	e.a["method"] = v
	return e
}

func (e *htmlForm) Name(v string) *htmlForm {
	e.a["name"] = v
	return e
}

func (e *htmlForm) NoValidate(v bool) *htmlForm {
	if v {
		e.a["novalidate"] = ""
	} else {
		delete(e.a, "novalidate")
	}
	return e
}

func (e *htmlForm) Target(v string) *htmlForm {
	e.a["target"] = v
	return e
}

func (e *htmlForm) CheckValidity(v bool) *htmlForm {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlForm) ReportValidity(v bool) *htmlForm {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlForm) ID(v string) *htmlForm {
	e.a["id"] = v
	return e
}

func (e *htmlForm) Class(v string) *htmlForm {
	e.a["class"] = v
	return e
}

func (e *htmlForm) Title(v string) *htmlForm {
	e.a["title"] = v
	return e
}

func (e *htmlForm) Lang(v string) *htmlForm {
	e.a["lang"] = v
	return e
}

func (e *htmlForm) Translate(v bool) *htmlForm {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlForm) Dir(v string) *htmlForm {
	e.a["dir"] = v
	return e
}

func (e *htmlForm) Hidden(v bool) *htmlForm {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlForm) TabIndex(v int) *htmlForm {
	e.a["tabindex"] = v
	return e
}

func (e *htmlForm) AccessKey(v string) *htmlForm {
	e.a["accesskey"] = v
	return e
}

func (e *htmlForm) Draggable(v bool) *htmlForm {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlForm) Spellcheck(v bool) *htmlForm {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
