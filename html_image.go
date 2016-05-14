package h

type htmlImage struct {
	htmlElement
}

func Image() *htmlImage {
	e := &htmlImage{}
	e.a = make(map[string]interface{})
	e.tagName = "image"
	return e
}

func (e *htmlImage) S(style StyleMap) *htmlImage {
	e.htmlElement.S(style)
	return e
}

func (e *htmlImage) Key(key interface{}) *htmlImage {
	e.key = F(key)
	return e
}

func (e *htmlImage) Alt(v string) *htmlImage {
	e.a["alt"] = v
	return e
}

func (e *htmlImage) Src(v string) *htmlImage {
	e.a["src"] = v
	return e
}

func (e *htmlImage) Srcset(v string) *htmlImage {
	e.a["srcset"] = v
	return e
}

func (e *htmlImage) Sizes(v string) *htmlImage {
	e.a["sizes"] = v
	return e
}

func (e *htmlImage) CrossOrigin(v string) *htmlImage {
	e.a["crossorigin"] = v
	return e
}

func (e *htmlImage) UseMap(v string) *htmlImage {
	e.a["usemap"] = v
	return e
}

func (e *htmlImage) IsMap(v bool) *htmlImage {
	if v {
		e.a["ismap"] = ""
	} else {
		delete(e.a, "ismap")
	}
	return e
}

func (e *htmlImage) Width(v int) *htmlImage {
	e.a["width"] = v
	return e
}

func (e *htmlImage) Height(v int) *htmlImage {
	e.a["height"] = v
	return e
}

func (e *htmlImage) ReferrerPolicy(v string) *htmlImage {
	e.a["referrerpolicy"] = v
	return e
}

func (e *htmlImage) ID(v string) *htmlImage {
	e.a["id"] = v
	return e
}

func (e *htmlImage) Class(v string) *htmlImage {
	e.a["class"] = v
	return e
}

func (e *htmlImage) Title(v string) *htmlImage {
	e.a["title"] = v
	return e
}

func (e *htmlImage) Lang(v string) *htmlImage {
	e.a["lang"] = v
	return e
}

func (e *htmlImage) Translate(v bool) *htmlImage {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlImage) Dir(v string) *htmlImage {
	e.a["dir"] = v
	return e
}

func (e *htmlImage) Hidden(v bool) *htmlImage {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlImage) TabIndex(v int) *htmlImage {
	e.a["tabindex"] = v
	return e
}

func (e *htmlImage) AccessKey(v string) *htmlImage {
	e.a["accesskey"] = v
	return e
}

func (e *htmlImage) Draggable(v bool) *htmlImage {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlImage) Spellcheck(v bool) *htmlImage {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
