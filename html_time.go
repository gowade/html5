package h

type htmlTime struct {
	htmlElement
}

func Time() *htmlTime {
	e := &htmlTime{}
	e.a = make(map[string]interface{})
	e.tagName = "time"
	return e
}

func (e *htmlTime) S(style StyleMap) *htmlTime {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTime) Key(key interface{}) *htmlTime {
	e.key = F(key)
	return e
}

func (e *htmlTime) DateTime(v string) *htmlTime {
	e.a["datetime"] = v
	return e
}

func (e *htmlTime) ID(v string) *htmlTime {
	e.a["id"] = v
	return e
}

func (e *htmlTime) Class(v string) *htmlTime {
	e.a["class"] = v
	return e
}

func (e *htmlTime) Title(v string) *htmlTime {
	e.a["title"] = v
	return e
}

func (e *htmlTime) Lang(v string) *htmlTime {
	e.a["lang"] = v
	return e
}

func (e *htmlTime) Translate(v bool) *htmlTime {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTime) Dir(v string) *htmlTime {
	e.a["dir"] = v
	return e
}

func (e *htmlTime) Hidden(v bool) *htmlTime {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTime) TabIndex(v int) *htmlTime {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTime) AccessKey(v string) *htmlTime {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTime) Draggable(v bool) *htmlTime {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTime) Spellcheck(v bool) *htmlTime {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
