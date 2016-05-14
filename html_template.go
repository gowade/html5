package h

type htmlTemplate struct {
	htmlElement
}

func Template() *htmlTemplate {
	e := &htmlTemplate{}
	e.a = make(map[string]interface{})
	e.tagName = "template"
	return e
}

func (e *htmlTemplate) S(style StyleMap) *htmlTemplate {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTemplate) Key(key interface{}) *htmlTemplate {
	e.key = F(key)
	return e
}

func (e *htmlTemplate) ID(v string) *htmlTemplate {
	e.a["id"] = v
	return e
}

func (e *htmlTemplate) Class(v string) *htmlTemplate {
	e.a["class"] = v
	return e
}

func (e *htmlTemplate) Title(v string) *htmlTemplate {
	e.a["title"] = v
	return e
}

func (e *htmlTemplate) Lang(v string) *htmlTemplate {
	e.a["lang"] = v
	return e
}

func (e *htmlTemplate) Translate(v bool) *htmlTemplate {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTemplate) Dir(v string) *htmlTemplate {
	e.a["dir"] = v
	return e
}

func (e *htmlTemplate) Hidden(v bool) *htmlTemplate {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTemplate) TabIndex(v int) *htmlTemplate {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTemplate) AccessKey(v string) *htmlTemplate {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTemplate) Draggable(v bool) *htmlTemplate {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTemplate) Spellcheck(v bool) *htmlTemplate {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
