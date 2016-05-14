package h

type htmlSource struct {
	htmlElement
}

func Source() *htmlSource {
	e := &htmlSource{}
	e.a = make(map[string]interface{})
	e.tagName = "source"
	return e
}

func (e *htmlSource) S(style StyleMap) *htmlSource {
	e.htmlElement.S(style)
	return e
}

func (e *htmlSource) Key(key interface{}) *htmlSource {
	e.key = F(key)
	return e
}

func (e *htmlSource) Src(v string) *htmlSource {
	e.a["src"] = v
	return e
}

func (e *htmlSource) Type(v string) *htmlSource {
	e.a["type"] = v
	return e
}

func (e *htmlSource) Srcset(v string) *htmlSource {
	e.a["srcset"] = v
	return e
}

func (e *htmlSource) Sizes(v string) *htmlSource {
	e.a["sizes"] = v
	return e
}

func (e *htmlSource) Media(v string) *htmlSource {
	e.a["media"] = v
	return e
}

func (e *htmlSource) ID(v string) *htmlSource {
	e.a["id"] = v
	return e
}

func (e *htmlSource) Class(v string) *htmlSource {
	e.a["class"] = v
	return e
}

func (e *htmlSource) Title(v string) *htmlSource {
	e.a["title"] = v
	return e
}

func (e *htmlSource) Lang(v string) *htmlSource {
	e.a["lang"] = v
	return e
}

func (e *htmlSource) Translate(v bool) *htmlSource {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlSource) Dir(v string) *htmlSource {
	e.a["dir"] = v
	return e
}

func (e *htmlSource) Hidden(v bool) *htmlSource {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlSource) TabIndex(v int) *htmlSource {
	e.a["tabindex"] = v
	return e
}

func (e *htmlSource) AccessKey(v string) *htmlSource {
	e.a["accesskey"] = v
	return e
}

func (e *htmlSource) Draggable(v bool) *htmlSource {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlSource) Spellcheck(v bool) *htmlSource {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
