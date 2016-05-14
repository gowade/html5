package h

type htmlMeta struct {
	htmlElement
}

func Meta() *htmlMeta {
	e := &htmlMeta{}
	e.a = make(map[string]interface{})
	e.tagName = "meta"
	return e
}

func (e *htmlMeta) S(style StyleMap) *htmlMeta {
	e.htmlElement.S(style)
	return e
}

func (e *htmlMeta) Key(key interface{}) *htmlMeta {
	e.key = F(key)
	return e
}

func (e *htmlMeta) Name(v string) *htmlMeta {
	e.a["name"] = v
	return e
}

func (e *htmlMeta) HttpEquiv(v string) *htmlMeta {
	e.a["httpequiv"] = v
	return e
}

func (e *htmlMeta) Content(v string) *htmlMeta {
	e.a["content"] = v
	return e
}

func (e *htmlMeta) ID(v string) *htmlMeta {
	e.a["id"] = v
	return e
}

func (e *htmlMeta) Class(v string) *htmlMeta {
	e.a["class"] = v
	return e
}

func (e *htmlMeta) Title(v string) *htmlMeta {
	e.a["title"] = v
	return e
}

func (e *htmlMeta) Lang(v string) *htmlMeta {
	e.a["lang"] = v
	return e
}

func (e *htmlMeta) Translate(v bool) *htmlMeta {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlMeta) Dir(v string) *htmlMeta {
	e.a["dir"] = v
	return e
}

func (e *htmlMeta) Hidden(v bool) *htmlMeta {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlMeta) TabIndex(v int) *htmlMeta {
	e.a["tabindex"] = v
	return e
}

func (e *htmlMeta) AccessKey(v string) *htmlMeta {
	e.a["accesskey"] = v
	return e
}

func (e *htmlMeta) Draggable(v bool) *htmlMeta {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlMeta) Spellcheck(v bool) *htmlMeta {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
