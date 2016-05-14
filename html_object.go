package h

type htmlObject struct {
	htmlElement
}

func Object() *htmlObject {
	e := &htmlObject{}
	e.a = make(map[string]interface{})
	e.tagName = "object"
	return e
}

func (e *htmlObject) S(style StyleMap) *htmlObject {
	e.htmlElement.S(style)
	return e
}

func (e *htmlObject) Key(key interface{}) *htmlObject {
	e.key = F(key)
	return e
}

func (e *htmlObject) Data(v string) *htmlObject {
	e.a["data"] = v
	return e
}

func (e *htmlObject) Type(v string) *htmlObject {
	e.a["type"] = v
	return e
}

func (e *htmlObject) TypeMustMatch(v bool) *htmlObject {
	if v {
		e.a["typemustmatch"] = ""
	} else {
		delete(e.a, "typemustmatch")
	}
	return e
}

func (e *htmlObject) Name(v string) *htmlObject {
	e.a["name"] = v
	return e
}

func (e *htmlObject) UseMap(v string) *htmlObject {
	e.a["usemap"] = v
	return e
}

func (e *htmlObject) Width(v string) *htmlObject {
	e.a["width"] = v
	return e
}

func (e *htmlObject) Height(v string) *htmlObject {
	e.a["height"] = v
	return e
}

func (e *htmlObject) CheckValidity(v bool) *htmlObject {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlObject) ReportValidity(v bool) *htmlObject {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlObject) ID(v string) *htmlObject {
	e.a["id"] = v
	return e
}

func (e *htmlObject) Class(v string) *htmlObject {
	e.a["class"] = v
	return e
}

func (e *htmlObject) Title(v string) *htmlObject {
	e.a["title"] = v
	return e
}

func (e *htmlObject) Lang(v string) *htmlObject {
	e.a["lang"] = v
	return e
}

func (e *htmlObject) Translate(v bool) *htmlObject {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlObject) Dir(v string) *htmlObject {
	e.a["dir"] = v
	return e
}

func (e *htmlObject) Hidden(v bool) *htmlObject {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlObject) TabIndex(v int) *htmlObject {
	e.a["tabindex"] = v
	return e
}

func (e *htmlObject) AccessKey(v string) *htmlObject {
	e.a["accesskey"] = v
	return e
}

func (e *htmlObject) Draggable(v bool) *htmlObject {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlObject) Spellcheck(v bool) *htmlObject {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
