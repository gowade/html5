package h

type htmlTableCaption struct {
	htmlElement
}

func TableCaption() *htmlTableCaption {
	e := &htmlTableCaption{}
	e.a = make(map[string]interface{})
	e.tagName = "tablecaption"
	return e
}

func (e *htmlTableCaption) S(style StyleMap) *htmlTableCaption {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTableCaption) Key(key interface{}) *htmlTableCaption {
	e.key = F(key)
	return e
}

func (e *htmlTableCaption) ID(v string) *htmlTableCaption {
	e.a["id"] = v
	return e
}

func (e *htmlTableCaption) Class(v string) *htmlTableCaption {
	e.a["class"] = v
	return e
}

func (e *htmlTableCaption) Title(v string) *htmlTableCaption {
	e.a["title"] = v
	return e
}

func (e *htmlTableCaption) Lang(v string) *htmlTableCaption {
	e.a["lang"] = v
	return e
}

func (e *htmlTableCaption) Translate(v bool) *htmlTableCaption {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTableCaption) Dir(v string) *htmlTableCaption {
	e.a["dir"] = v
	return e
}

func (e *htmlTableCaption) Hidden(v bool) *htmlTableCaption {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTableCaption) TabIndex(v int) *htmlTableCaption {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTableCaption) AccessKey(v string) *htmlTableCaption {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTableCaption) Draggable(v bool) *htmlTableCaption {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTableCaption) Spellcheck(v bool) *htmlTableCaption {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
