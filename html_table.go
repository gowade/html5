package h

type htmlTable struct {
	htmlElement
}

func Table() *htmlTable {
	e := &htmlTable{}
	e.a = make(map[string]interface{})
	e.tagName = "table"
	return e
}

func (e *htmlTable) S(style StyleMap) *htmlTable {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTable) Key(key interface{}) *htmlTable {
	e.key = F(key)
	return e
}

func (e *htmlTable) ID(v string) *htmlTable {
	e.a["id"] = v
	return e
}

func (e *htmlTable) Class(v string) *htmlTable {
	e.a["class"] = v
	return e
}

func (e *htmlTable) Title(v string) *htmlTable {
	e.a["title"] = v
	return e
}

func (e *htmlTable) Lang(v string) *htmlTable {
	e.a["lang"] = v
	return e
}

func (e *htmlTable) Translate(v bool) *htmlTable {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTable) Dir(v string) *htmlTable {
	e.a["dir"] = v
	return e
}

func (e *htmlTable) Hidden(v bool) *htmlTable {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTable) TabIndex(v int) *htmlTable {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTable) AccessKey(v string) *htmlTable {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTable) Draggable(v bool) *htmlTable {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTable) Spellcheck(v bool) *htmlTable {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
