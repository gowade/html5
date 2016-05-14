package h

type htmlTableSection struct {
	htmlElement
}

func TableSection() *htmlTableSection {
	e := &htmlTableSection{}
	e.a = make(map[string]interface{})
	e.tagName = "tablesection"
	return e
}

func (e *htmlTableSection) S(style StyleMap) *htmlTableSection {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTableSection) Key(key interface{}) *htmlTableSection {
	e.key = F(key)
	return e
}

func (e *htmlTableSection) ID(v string) *htmlTableSection {
	e.a["id"] = v
	return e
}

func (e *htmlTableSection) Class(v string) *htmlTableSection {
	e.a["class"] = v
	return e
}

func (e *htmlTableSection) Title(v string) *htmlTableSection {
	e.a["title"] = v
	return e
}

func (e *htmlTableSection) Lang(v string) *htmlTableSection {
	e.a["lang"] = v
	return e
}

func (e *htmlTableSection) Translate(v bool) *htmlTableSection {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTableSection) Dir(v string) *htmlTableSection {
	e.a["dir"] = v
	return e
}

func (e *htmlTableSection) Hidden(v bool) *htmlTableSection {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTableSection) TabIndex(v int) *htmlTableSection {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTableSection) AccessKey(v string) *htmlTableSection {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTableSection) Draggable(v bool) *htmlTableSection {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTableSection) Spellcheck(v bool) *htmlTableSection {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
