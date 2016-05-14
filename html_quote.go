package h

type htmlQuote struct {
	htmlElement
}

func Quote() *htmlQuote {
	e := &htmlQuote{}
	e.a = make(map[string]interface{})
	e.tagName = "quote"
	return e
}

func (e *htmlQuote) S(style StyleMap) *htmlQuote {
	e.htmlElement.S(style)
	return e
}

func (e *htmlQuote) Key(key interface{}) *htmlQuote {
	e.key = F(key)
	return e
}

func (e *htmlQuote) Cite(v string) *htmlQuote {
	e.a["cite"] = v
	return e
}

func (e *htmlQuote) ID(v string) *htmlQuote {
	e.a["id"] = v
	return e
}

func (e *htmlQuote) Class(v string) *htmlQuote {
	e.a["class"] = v
	return e
}

func (e *htmlQuote) Title(v string) *htmlQuote {
	e.a["title"] = v
	return e
}

func (e *htmlQuote) Lang(v string) *htmlQuote {
	e.a["lang"] = v
	return e
}

func (e *htmlQuote) Translate(v bool) *htmlQuote {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlQuote) Dir(v string) *htmlQuote {
	e.a["dir"] = v
	return e
}

func (e *htmlQuote) Hidden(v bool) *htmlQuote {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlQuote) TabIndex(v int) *htmlQuote {
	e.a["tabindex"] = v
	return e
}

func (e *htmlQuote) AccessKey(v string) *htmlQuote {
	e.a["accesskey"] = v
	return e
}

func (e *htmlQuote) Draggable(v bool) *htmlQuote {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlQuote) Spellcheck(v bool) *htmlQuote {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
