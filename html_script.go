package h

type htmlScript struct {
	htmlElement
}

func Script() *htmlScript {
	e := &htmlScript{}
	e.a = make(map[string]interface{})
	e.tagName = "script"
	return e
}

func (e *htmlScript) S(style StyleMap) *htmlScript {
	e.htmlElement.S(style)
	return e
}

func (e *htmlScript) Key(key interface{}) *htmlScript {
	e.key = F(key)
	return e
}

func (e *htmlScript) Src(v string) *htmlScript {
	e.a["src"] = v
	return e
}

func (e *htmlScript) Type(v string) *htmlScript {
	e.a["type"] = v
	return e
}

func (e *htmlScript) Charset(v string) *htmlScript {
	e.a["charset"] = v
	return e
}

func (e *htmlScript) Async(v bool) *htmlScript {
	if v {
		e.a["async"] = ""
	} else {
		delete(e.a, "async")
	}
	return e
}

func (e *htmlScript) Defer(v bool) *htmlScript {
	if v {
		e.a["defer"] = ""
	} else {
		delete(e.a, "defer")
	}
	return e
}

func (e *htmlScript) CrossOrigin(v string) *htmlScript {
	e.a["crossorigin"] = v
	return e
}

func (e *htmlScript) Text(v string) *htmlScript {
	e.a["text"] = v
	return e
}

func (e *htmlScript) Nonce(v string) *htmlScript {
	e.a["nonce"] = v
	return e
}

func (e *htmlScript) ID(v string) *htmlScript {
	e.a["id"] = v
	return e
}

func (e *htmlScript) Class(v string) *htmlScript {
	e.a["class"] = v
	return e
}

func (e *htmlScript) Title(v string) *htmlScript {
	e.a["title"] = v
	return e
}

func (e *htmlScript) Lang(v string) *htmlScript {
	e.a["lang"] = v
	return e
}

func (e *htmlScript) Translate(v bool) *htmlScript {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlScript) Dir(v string) *htmlScript {
	e.a["dir"] = v
	return e
}

func (e *htmlScript) Hidden(v bool) *htmlScript {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlScript) TabIndex(v int) *htmlScript {
	e.a["tabindex"] = v
	return e
}

func (e *htmlScript) AccessKey(v string) *htmlScript {
	e.a["accesskey"] = v
	return e
}

func (e *htmlScript) Draggable(v bool) *htmlScript {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlScript) Spellcheck(v bool) *htmlScript {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
