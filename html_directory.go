package h

type htmlDirectory struct {
	htmlElement
}

func Directory() *htmlDirectory {
	e := &htmlDirectory{}
	e.a = make(map[string]interface{})
	e.tagName = "directory"
	return e
}

func (e *htmlDirectory) S(style StyleMap) *htmlDirectory {
	e.htmlElement.S(style)
	return e
}

func (e *htmlDirectory) Key(key interface{}) *htmlDirectory {
	e.key = F(key)
	return e
}

func (e *htmlDirectory) Compact(v bool) *htmlDirectory {
	if v {
		e.a["compact"] = ""
	} else {
		delete(e.a, "compact")
	}
	return e
}

func (e *htmlDirectory) ID(v string) *htmlDirectory {
	e.a["id"] = v
	return e
}

func (e *htmlDirectory) Class(v string) *htmlDirectory {
	e.a["class"] = v
	return e
}

func (e *htmlDirectory) Title(v string) *htmlDirectory {
	e.a["title"] = v
	return e
}

func (e *htmlDirectory) Lang(v string) *htmlDirectory {
	e.a["lang"] = v
	return e
}

func (e *htmlDirectory) Translate(v bool) *htmlDirectory {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlDirectory) Dir(v string) *htmlDirectory {
	e.a["dir"] = v
	return e
}

func (e *htmlDirectory) Hidden(v bool) *htmlDirectory {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlDirectory) TabIndex(v int) *htmlDirectory {
	e.a["tabindex"] = v
	return e
}

func (e *htmlDirectory) AccessKey(v string) *htmlDirectory {
	e.a["accesskey"] = v
	return e
}

func (e *htmlDirectory) Draggable(v bool) *htmlDirectory {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlDirectory) Spellcheck(v bool) *htmlDirectory {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
