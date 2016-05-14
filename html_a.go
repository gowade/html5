package h

type htmlA struct {
	htmlElement
}

func A() *htmlA {
	e := &htmlA{}
	e.a = make(map[string]interface{})
	e.tagName = "a"
	return e
}

func (e *htmlA) S(style StyleMap) *htmlA {
	e.htmlElement.S(style)
	return e
}

func (e *htmlA) Key(key interface{}) *htmlA {
	e.key = F(key)
	return e
}

func (e *htmlA) Target(v string) *htmlA {
	e.a["target"] = v
	return e
}

func (e *htmlA) Download(v string) *htmlA {
	e.a["download"] = v
	return e
}

func (e *htmlA) Rel(v string) *htmlA {
	e.a["rel"] = v
	return e
}

func (e *htmlA) Hreflang(v string) *htmlA {
	e.a["hreflang"] = v
	return e
}

func (e *htmlA) Type(v string) *htmlA {
	e.a["type"] = v
	return e
}

func (e *htmlA) Text(v string) *htmlA {
	e.a["text"] = v
	return e
}

func (e *htmlA) ReferrerPolicy(v string) *htmlA {
	e.a["referrerpolicy"] = v
	return e
}

func (e *htmlA) ID(v string) *htmlA {
	e.a["id"] = v
	return e
}

func (e *htmlA) Class(v string) *htmlA {
	e.a["class"] = v
	return e
}

func (e *htmlA) Title(v string) *htmlA {
	e.a["title"] = v
	return e
}

func (e *htmlA) Lang(v string) *htmlA {
	e.a["lang"] = v
	return e
}

func (e *htmlA) Translate(v bool) *htmlA {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlA) Dir(v string) *htmlA {
	e.a["dir"] = v
	return e
}

func (e *htmlA) Hidden(v bool) *htmlA {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlA) TabIndex(v int) *htmlA {
	e.a["tabindex"] = v
	return e
}

func (e *htmlA) AccessKey(v string) *htmlA {
	e.a["accesskey"] = v
	return e
}

func (e *htmlA) Draggable(v bool) *htmlA {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlA) Spellcheck(v bool) *htmlA {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
