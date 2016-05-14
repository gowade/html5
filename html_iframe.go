package h

type htmlIFrame struct {
	htmlElement
}

func IFrame() *htmlIFrame {
	e := &htmlIFrame{}
	e.a = make(map[string]interface{})
	e.tagName = "iframe"
	return e
}

func (e *htmlIFrame) S(style StyleMap) *htmlIFrame {
	e.htmlElement.S(style)
	return e
}

func (e *htmlIFrame) Key(key interface{}) *htmlIFrame {
	e.key = F(key)
	return e
}

func (e *htmlIFrame) Src(v string) *htmlIFrame {
	e.a["src"] = v
	return e
}

func (e *htmlIFrame) Srcdoc(v string) *htmlIFrame {
	e.a["srcdoc"] = v
	return e
}

func (e *htmlIFrame) Name(v string) *htmlIFrame {
	e.a["name"] = v
	return e
}

func (e *htmlIFrame) AllowFullscreen(v bool) *htmlIFrame {
	if v {
		e.a["allowfullscreen"] = ""
	} else {
		delete(e.a, "allowfullscreen")
	}
	return e
}

func (e *htmlIFrame) Width(v string) *htmlIFrame {
	e.a["width"] = v
	return e
}

func (e *htmlIFrame) Height(v string) *htmlIFrame {
	e.a["height"] = v
	return e
}

func (e *htmlIFrame) ReferrerPolicy(v string) *htmlIFrame {
	e.a["referrerpolicy"] = v
	return e
}

func (e *htmlIFrame) ID(v string) *htmlIFrame {
	e.a["id"] = v
	return e
}

func (e *htmlIFrame) Class(v string) *htmlIFrame {
	e.a["class"] = v
	return e
}

func (e *htmlIFrame) Title(v string) *htmlIFrame {
	e.a["title"] = v
	return e
}

func (e *htmlIFrame) Lang(v string) *htmlIFrame {
	e.a["lang"] = v
	return e
}

func (e *htmlIFrame) Translate(v bool) *htmlIFrame {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlIFrame) Dir(v string) *htmlIFrame {
	e.a["dir"] = v
	return e
}

func (e *htmlIFrame) Hidden(v bool) *htmlIFrame {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlIFrame) TabIndex(v int) *htmlIFrame {
	e.a["tabindex"] = v
	return e
}

func (e *htmlIFrame) AccessKey(v string) *htmlIFrame {
	e.a["accesskey"] = v
	return e
}

func (e *htmlIFrame) Draggable(v bool) *htmlIFrame {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlIFrame) Spellcheck(v bool) *htmlIFrame {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
