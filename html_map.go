package h

type htmlMap struct {
	htmlElement
}

func Map() *htmlMap {
	e := &htmlMap{}
	e.a = make(map[string]interface{})
	e.tagName = "map"
	return e
}

func (e *htmlMap) S(style StyleMap) *htmlMap {
	e.htmlElement.S(style)
	return e
}

func (e *htmlMap) Key(key interface{}) *htmlMap {
	e.key = F(key)
	return e
}

func (e *htmlMap) Name(v string) *htmlMap {
	e.a["name"] = v
	return e
}

func (e *htmlMap) ID(v string) *htmlMap {
	e.a["id"] = v
	return e
}

func (e *htmlMap) Class(v string) *htmlMap {
	e.a["class"] = v
	return e
}

func (e *htmlMap) Title(v string) *htmlMap {
	e.a["title"] = v
	return e
}

func (e *htmlMap) Lang(v string) *htmlMap {
	e.a["lang"] = v
	return e
}

func (e *htmlMap) Translate(v bool) *htmlMap {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlMap) Dir(v string) *htmlMap {
	e.a["dir"] = v
	return e
}

func (e *htmlMap) Hidden(v bool) *htmlMap {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlMap) TabIndex(v int) *htmlMap {
	e.a["tabindex"] = v
	return e
}

func (e *htmlMap) AccessKey(v string) *htmlMap {
	e.a["accesskey"] = v
	return e
}

func (e *htmlMap) Draggable(v bool) *htmlMap {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlMap) Spellcheck(v bool) *htmlMap {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
