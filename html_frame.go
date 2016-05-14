package h

type htmlFrame struct {
	htmlElement
}

func Frame() *htmlFrame {
	e := &htmlFrame{}
	e.a = make(map[string]interface{})
	e.tagName = "frame"
	return e
}

func (e *htmlFrame) S(style StyleMap) *htmlFrame {
	e.htmlElement.S(style)
	return e
}

func (e *htmlFrame) Key(key interface{}) *htmlFrame {
	e.key = F(key)
	return e
}

func (e *htmlFrame) Name(v string) *htmlFrame {
	e.a["name"] = v
	return e
}

func (e *htmlFrame) Scrolling(v string) *htmlFrame {
	e.a["scrolling"] = v
	return e
}

func (e *htmlFrame) Src(v string) *htmlFrame {
	e.a["src"] = v
	return e
}

func (e *htmlFrame) FrameBorder(v string) *htmlFrame {
	e.a["frameborder"] = v
	return e
}

func (e *htmlFrame) LongDesc(v string) *htmlFrame {
	e.a["longdesc"] = v
	return e
}

func (e *htmlFrame) NoResize(v bool) *htmlFrame {
	if v {
		e.a["noresize"] = ""
	} else {
		delete(e.a, "noresize")
	}
	return e
}

func (e *htmlFrame) MarginHeight(v string) *htmlFrame {
	e.a["marginheight"] = v
	return e
}

func (e *htmlFrame) MarginWidth(v string) *htmlFrame {
	e.a["marginwidth"] = v
	return e
}

func (e *htmlFrame) ID(v string) *htmlFrame {
	e.a["id"] = v
	return e
}

func (e *htmlFrame) Class(v string) *htmlFrame {
	e.a["class"] = v
	return e
}

func (e *htmlFrame) Title(v string) *htmlFrame {
	e.a["title"] = v
	return e
}

func (e *htmlFrame) Lang(v string) *htmlFrame {
	e.a["lang"] = v
	return e
}

func (e *htmlFrame) Translate(v bool) *htmlFrame {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlFrame) Dir(v string) *htmlFrame {
	e.a["dir"] = v
	return e
}

func (e *htmlFrame) Hidden(v bool) *htmlFrame {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlFrame) TabIndex(v int) *htmlFrame {
	e.a["tabindex"] = v
	return e
}

func (e *htmlFrame) AccessKey(v string) *htmlFrame {
	e.a["accesskey"] = v
	return e
}

func (e *htmlFrame) Draggable(v bool) *htmlFrame {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlFrame) Spellcheck(v bool) *htmlFrame {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
