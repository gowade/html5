package h

type htmlOutput struct {
	htmlElement
}

func Output() *htmlOutput {
	e := &htmlOutput{}
	e.a = make(map[string]interface{})
	e.tagName = "output"
	return e
}

func (e *htmlOutput) S(style StyleMap) *htmlOutput {
	e.htmlElement.S(style)
	return e
}

func (e *htmlOutput) Key(key interface{}) *htmlOutput {
	e.key = F(key)
	return e
}

func (e *htmlOutput) Name(v string) *htmlOutput {
	e.a["name"] = v
	return e
}

func (e *htmlOutput) DefaultValue(v string) *htmlOutput {
	e.a["defaultvalue"] = v
	return e
}

func (e *htmlOutput) Value(v string) *htmlOutput {
	e.a["value"] = v
	return e
}

func (e *htmlOutput) CheckValidity(v bool) *htmlOutput {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlOutput) ReportValidity(v bool) *htmlOutput {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlOutput) ID(v string) *htmlOutput {
	e.a["id"] = v
	return e
}

func (e *htmlOutput) Class(v string) *htmlOutput {
	e.a["class"] = v
	return e
}

func (e *htmlOutput) Title(v string) *htmlOutput {
	e.a["title"] = v
	return e
}

func (e *htmlOutput) Lang(v string) *htmlOutput {
	e.a["lang"] = v
	return e
}

func (e *htmlOutput) Translate(v bool) *htmlOutput {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlOutput) Dir(v string) *htmlOutput {
	e.a["dir"] = v
	return e
}

func (e *htmlOutput) Hidden(v bool) *htmlOutput {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlOutput) TabIndex(v int) *htmlOutput {
	e.a["tabindex"] = v
	return e
}

func (e *htmlOutput) AccessKey(v string) *htmlOutput {
	e.a["accesskey"] = v
	return e
}

func (e *htmlOutput) Draggable(v bool) *htmlOutput {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlOutput) Spellcheck(v bool) *htmlOutput {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
