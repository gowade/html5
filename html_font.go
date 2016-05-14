package h

type htmlFont struct {
	htmlElement
}

func Font() *htmlFont {
	e := &htmlFont{}
	e.a = make(map[string]interface{})
	e.tagName = "font"
	return e
}

func (e *htmlFont) S(style StyleMap) *htmlFont {
	e.htmlElement.S(style)
	return e
}

func (e *htmlFont) Key(key interface{}) *htmlFont {
	e.key = F(key)
	return e
}

func (e *htmlFont) Color(v string) *htmlFont {
	e.a["color"] = v
	return e
}

func (e *htmlFont) Face(v string) *htmlFont {
	e.a["face"] = v
	return e
}

func (e *htmlFont) Size(v string) *htmlFont {
	e.a["size"] = v
	return e
}

func (e *htmlFont) ID(v string) *htmlFont {
	e.a["id"] = v
	return e
}

func (e *htmlFont) Class(v string) *htmlFont {
	e.a["class"] = v
	return e
}

func (e *htmlFont) Title(v string) *htmlFont {
	e.a["title"] = v
	return e
}

func (e *htmlFont) Lang(v string) *htmlFont {
	e.a["lang"] = v
	return e
}

func (e *htmlFont) Translate(v bool) *htmlFont {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlFont) Dir(v string) *htmlFont {
	e.a["dir"] = v
	return e
}

func (e *htmlFont) Hidden(v bool) *htmlFont {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlFont) TabIndex(v int) *htmlFont {
	e.a["tabindex"] = v
	return e
}

func (e *htmlFont) AccessKey(v string) *htmlFont {
	e.a["accesskey"] = v
	return e
}

func (e *htmlFont) Draggable(v bool) *htmlFont {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlFont) Spellcheck(v bool) *htmlFont {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
