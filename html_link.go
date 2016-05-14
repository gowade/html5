package h

type htmlLink struct {
	htmlElement
}

func Link() *htmlLink {
	e := &htmlLink{}
	e.a = make(map[string]interface{})
	e.tagName = "link"
	return e
}

func (e *htmlLink) S(style StyleMap) *htmlLink {
	e.htmlElement.S(style)
	return e
}

func (e *htmlLink) Key(key interface{}) *htmlLink {
	e.key = F(key)
	return e
}

func (e *htmlLink) Href(v string) *htmlLink {
	e.a["href"] = v
	return e
}

func (e *htmlLink) CrossOrigin(v string) *htmlLink {
	e.a["crossorigin"] = v
	return e
}

func (e *htmlLink) Rel(v string) *htmlLink {
	e.a["rel"] = v
	return e
}

func (e *htmlLink) Media(v string) *htmlLink {
	e.a["media"] = v
	return e
}

func (e *htmlLink) Hreflang(v string) *htmlLink {
	e.a["hreflang"] = v
	return e
}

func (e *htmlLink) Type(v string) *htmlLink {
	e.a["type"] = v
	return e
}

func (e *htmlLink) ReferrerPolicy(v string) *htmlLink {
	e.a["referrerpolicy"] = v
	return e
}

func (e *htmlLink) ID(v string) *htmlLink {
	e.a["id"] = v
	return e
}

func (e *htmlLink) Class(v string) *htmlLink {
	e.a["class"] = v
	return e
}

func (e *htmlLink) Title(v string) *htmlLink {
	e.a["title"] = v
	return e
}

func (e *htmlLink) Lang(v string) *htmlLink {
	e.a["lang"] = v
	return e
}

func (e *htmlLink) Translate(v bool) *htmlLink {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlLink) Dir(v string) *htmlLink {
	e.a["dir"] = v
	return e
}

func (e *htmlLink) Hidden(v bool) *htmlLink {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlLink) TabIndex(v int) *htmlLink {
	e.a["tabindex"] = v
	return e
}

func (e *htmlLink) AccessKey(v string) *htmlLink {
	e.a["accesskey"] = v
	return e
}

func (e *htmlLink) Draggable(v bool) *htmlLink {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlLink) Spellcheck(v bool) *htmlLink {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
