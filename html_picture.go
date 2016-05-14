package h

type htmlPicture struct {
	htmlElement
}

func Picture() *htmlPicture {
	e := &htmlPicture{}
	e.a = make(map[string]interface{})
	e.tagName = "picture"
	return e
}

func (e *htmlPicture) S(style StyleMap) *htmlPicture {
	e.htmlElement.S(style)
	return e
}

func (e *htmlPicture) Key(key interface{}) *htmlPicture {
	e.key = F(key)
	return e
}

func (e *htmlPicture) ID(v string) *htmlPicture {
	e.a["id"] = v
	return e
}

func (e *htmlPicture) Class(v string) *htmlPicture {
	e.a["class"] = v
	return e
}

func (e *htmlPicture) Title(v string) *htmlPicture {
	e.a["title"] = v
	return e
}

func (e *htmlPicture) Lang(v string) *htmlPicture {
	e.a["lang"] = v
	return e
}

func (e *htmlPicture) Translate(v bool) *htmlPicture {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlPicture) Dir(v string) *htmlPicture {
	e.a["dir"] = v
	return e
}

func (e *htmlPicture) Hidden(v bool) *htmlPicture {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlPicture) TabIndex(v int) *htmlPicture {
	e.a["tabindex"] = v
	return e
}

func (e *htmlPicture) AccessKey(v string) *htmlPicture {
	e.a["accesskey"] = v
	return e
}

func (e *htmlPicture) Draggable(v bool) *htmlPicture {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlPicture) Spellcheck(v bool) *htmlPicture {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
