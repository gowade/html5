package h

type htmlCanvas struct {
	htmlElement
}

func Canvas() *htmlCanvas {
	e := &htmlCanvas{}
	e.a = make(map[string]interface{})
	e.tagName = "canvas"
	return e
}

func (e *htmlCanvas) S(style StyleMap) *htmlCanvas {
	e.htmlElement.S(style)
	return e
}

func (e *htmlCanvas) Key(key interface{}) *htmlCanvas {
	e.key = F(key)
	return e
}

func (e *htmlCanvas) Width(v int) *htmlCanvas {
	e.a["width"] = v
	return e
}

func (e *htmlCanvas) Height(v int) *htmlCanvas {
	e.a["height"] = v
	return e
}

func (e *htmlCanvas) ProbablySupportsContext(v bool) *htmlCanvas {
	if v {
		e.a["probablysupportscontext"] = ""
	} else {
		delete(e.a, "probablysupportscontext")
	}
	return e
}

func (e *htmlCanvas) ToDataURL(v string) *htmlCanvas {
	e.a["todataurl"] = v
	return e
}

func (e *htmlCanvas) ID(v string) *htmlCanvas {
	e.a["id"] = v
	return e
}

func (e *htmlCanvas) Class(v string) *htmlCanvas {
	e.a["class"] = v
	return e
}

func (e *htmlCanvas) Title(v string) *htmlCanvas {
	e.a["title"] = v
	return e
}

func (e *htmlCanvas) Lang(v string) *htmlCanvas {
	e.a["lang"] = v
	return e
}

func (e *htmlCanvas) Translate(v bool) *htmlCanvas {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlCanvas) Dir(v string) *htmlCanvas {
	e.a["dir"] = v
	return e
}

func (e *htmlCanvas) Hidden(v bool) *htmlCanvas {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlCanvas) TabIndex(v int) *htmlCanvas {
	e.a["tabindex"] = v
	return e
}

func (e *htmlCanvas) AccessKey(v string) *htmlCanvas {
	e.a["accesskey"] = v
	return e
}

func (e *htmlCanvas) Draggable(v bool) *htmlCanvas {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlCanvas) Spellcheck(v bool) *htmlCanvas {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
