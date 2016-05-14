package h

type htmlFieldSet struct {
	htmlElement
}

func FieldSet() *htmlFieldSet {
	e := &htmlFieldSet{}
	e.a = make(map[string]interface{})
	e.tagName = "fieldset"
	return e
}

func (e *htmlFieldSet) S(style StyleMap) *htmlFieldSet {
	e.htmlElement.S(style)
	return e
}

func (e *htmlFieldSet) Key(key interface{}) *htmlFieldSet {
	e.key = F(key)
	return e
}

func (e *htmlFieldSet) Disabled(v bool) *htmlFieldSet {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlFieldSet) Name(v string) *htmlFieldSet {
	e.a["name"] = v
	return e
}

func (e *htmlFieldSet) CheckValidity(v bool) *htmlFieldSet {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlFieldSet) ReportValidity(v bool) *htmlFieldSet {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlFieldSet) ID(v string) *htmlFieldSet {
	e.a["id"] = v
	return e
}

func (e *htmlFieldSet) Class(v string) *htmlFieldSet {
	e.a["class"] = v
	return e
}

func (e *htmlFieldSet) Title(v string) *htmlFieldSet {
	e.a["title"] = v
	return e
}

func (e *htmlFieldSet) Lang(v string) *htmlFieldSet {
	e.a["lang"] = v
	return e
}

func (e *htmlFieldSet) Translate(v bool) *htmlFieldSet {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlFieldSet) Dir(v string) *htmlFieldSet {
	e.a["dir"] = v
	return e
}

func (e *htmlFieldSet) Hidden(v bool) *htmlFieldSet {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlFieldSet) TabIndex(v int) *htmlFieldSet {
	e.a["tabindex"] = v
	return e
}

func (e *htmlFieldSet) AccessKey(v string) *htmlFieldSet {
	e.a["accesskey"] = v
	return e
}

func (e *htmlFieldSet) Draggable(v bool) *htmlFieldSet {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlFieldSet) Spellcheck(v bool) *htmlFieldSet {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
