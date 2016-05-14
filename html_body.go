package h

type htmlBody struct {
	htmlElement
}

func Body() *htmlBody {
	e := &htmlBody{}
	e.a = make(map[string]interface{})
	e.tagName = "body"
	return e
}

func (e *htmlBody) S(style StyleMap) *htmlBody {
	e.htmlElement.S(style)
	return e
}

func (e *htmlBody) Key(key interface{}) *htmlBody {
	e.key = F(key)
	return e
}

func (e *htmlBody) ID(v string) *htmlBody {
	e.a["id"] = v
	return e
}

func (e *htmlBody) Class(v string) *htmlBody {
	e.a["class"] = v
	return e
}

func (e *htmlBody) Title(v string) *htmlBody {
	e.a["title"] = v
	return e
}

func (e *htmlBody) Lang(v string) *htmlBody {
	e.a["lang"] = v
	return e
}

func (e *htmlBody) Translate(v bool) *htmlBody {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlBody) Dir(v string) *htmlBody {
	e.a["dir"] = v
	return e
}

func (e *htmlBody) Hidden(v bool) *htmlBody {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlBody) TabIndex(v int) *htmlBody {
	e.a["tabindex"] = v
	return e
}

func (e *htmlBody) AccessKey(v string) *htmlBody {
	e.a["accesskey"] = v
	return e
}

func (e *htmlBody) Draggable(v bool) *htmlBody {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlBody) Spellcheck(v bool) *htmlBody {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
