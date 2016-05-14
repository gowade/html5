package h

type htmlMeter struct {
	htmlElement
}

func Meter() *htmlMeter {
	e := &htmlMeter{}
	e.a = make(map[string]interface{})
	e.tagName = "meter"
	return e
}

func (e *htmlMeter) S(style StyleMap) *htmlMeter {
	e.htmlElement.S(style)
	return e
}

func (e *htmlMeter) Key(key interface{}) *htmlMeter {
	e.key = F(key)
	return e
}

func (e *htmlMeter) ID(v string) *htmlMeter {
	e.a["id"] = v
	return e
}

func (e *htmlMeter) Class(v string) *htmlMeter {
	e.a["class"] = v
	return e
}

func (e *htmlMeter) Title(v string) *htmlMeter {
	e.a["title"] = v
	return e
}

func (e *htmlMeter) Lang(v string) *htmlMeter {
	e.a["lang"] = v
	return e
}

func (e *htmlMeter) Translate(v bool) *htmlMeter {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlMeter) Dir(v string) *htmlMeter {
	e.a["dir"] = v
	return e
}

func (e *htmlMeter) Hidden(v bool) *htmlMeter {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlMeter) TabIndex(v int) *htmlMeter {
	e.a["tabindex"] = v
	return e
}

func (e *htmlMeter) AccessKey(v string) *htmlMeter {
	e.a["accesskey"] = v
	return e
}

func (e *htmlMeter) Draggable(v bool) *htmlMeter {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlMeter) Spellcheck(v bool) *htmlMeter {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
