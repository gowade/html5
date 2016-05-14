package h

type htmlTitle struct {
	htmlElement
}

func Title() *htmlTitle {
	e := &htmlTitle{}
	e.a = make(map[string]interface{})
	e.tagName = "title"
	return e
}

func (e *htmlTitle) S(style StyleMap) *htmlTitle {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTitle) Key(key interface{}) *htmlTitle {
	e.key = F(key)
	return e
}

func (e *htmlTitle) Text(v string) *htmlTitle {
	e.a["text"] = v
	return e
}

func (e *htmlTitle) ID(v string) *htmlTitle {
	e.a["id"] = v
	return e
}

func (e *htmlTitle) Class(v string) *htmlTitle {
	e.a["class"] = v
	return e
}

func (e *htmlTitle) Title(v string) *htmlTitle {
	e.a["title"] = v
	return e
}

func (e *htmlTitle) Lang(v string) *htmlTitle {
	e.a["lang"] = v
	return e
}

func (e *htmlTitle) Translate(v bool) *htmlTitle {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTitle) Dir(v string) *htmlTitle {
	e.a["dir"] = v
	return e
}

func (e *htmlTitle) Hidden(v bool) *htmlTitle {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTitle) TabIndex(v int) *htmlTitle {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTitle) AccessKey(v string) *htmlTitle {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTitle) Draggable(v bool) *htmlTitle {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTitle) Spellcheck(v bool) *htmlTitle {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
