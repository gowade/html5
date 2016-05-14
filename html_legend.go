package h

type htmlLegend struct {
	htmlElement
}

func Legend() *htmlLegend {
	e := &htmlLegend{}
	e.a = make(map[string]interface{})
	e.tagName = "legend"
	return e
}

func (e *htmlLegend) S(style StyleMap) *htmlLegend {
	e.htmlElement.S(style)
	return e
}

func (e *htmlLegend) Key(key interface{}) *htmlLegend {
	e.key = F(key)
	return e
}

func (e *htmlLegend) ID(v string) *htmlLegend {
	e.a["id"] = v
	return e
}

func (e *htmlLegend) Class(v string) *htmlLegend {
	e.a["class"] = v
	return e
}

func (e *htmlLegend) Title(v string) *htmlLegend {
	e.a["title"] = v
	return e
}

func (e *htmlLegend) Lang(v string) *htmlLegend {
	e.a["lang"] = v
	return e
}

func (e *htmlLegend) Translate(v bool) *htmlLegend {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlLegend) Dir(v string) *htmlLegend {
	e.a["dir"] = v
	return e
}

func (e *htmlLegend) Hidden(v bool) *htmlLegend {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlLegend) TabIndex(v int) *htmlLegend {
	e.a["tabindex"] = v
	return e
}

func (e *htmlLegend) AccessKey(v string) *htmlLegend {
	e.a["accesskey"] = v
	return e
}

func (e *htmlLegend) Draggable(v bool) *htmlLegend {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlLegend) Spellcheck(v bool) *htmlLegend {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
