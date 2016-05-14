package h

type htmlData struct {
	htmlElement
}

func Data() *htmlData {
	e := &htmlData{}
	e.a = make(map[string]interface{})
	e.tagName = "data"
	return e
}

func (e *htmlData) S(style StyleMap) *htmlData {
	e.htmlElement.S(style)
	return e
}

func (e *htmlData) Key(key interface{}) *htmlData {
	e.key = F(key)
	return e
}

func (e *htmlData) Value(v string) *htmlData {
	e.a["value"] = v
	return e
}

func (e *htmlData) ID(v string) *htmlData {
	e.a["id"] = v
	return e
}

func (e *htmlData) Class(v string) *htmlData {
	e.a["class"] = v
	return e
}

func (e *htmlData) Title(v string) *htmlData {
	e.a["title"] = v
	return e
}

func (e *htmlData) Lang(v string) *htmlData {
	e.a["lang"] = v
	return e
}

func (e *htmlData) Translate(v bool) *htmlData {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlData) Dir(v string) *htmlData {
	e.a["dir"] = v
	return e
}

func (e *htmlData) Hidden(v bool) *htmlData {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlData) TabIndex(v int) *htmlData {
	e.a["tabindex"] = v
	return e
}

func (e *htmlData) AccessKey(v string) *htmlData {
	e.a["accesskey"] = v
	return e
}

func (e *htmlData) Draggable(v bool) *htmlData {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlData) Spellcheck(v bool) *htmlData {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
