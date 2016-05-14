package h

type htmlApplet struct {
	htmlElement
}

func Applet() *htmlApplet {
	e := &htmlApplet{}
	e.a = make(map[string]interface{})
	e.tagName = "applet"
	return e
}

func (e *htmlApplet) S(style StyleMap) *htmlApplet {
	e.htmlElement.S(style)
	return e
}

func (e *htmlApplet) Key(key interface{}) *htmlApplet {
	e.key = F(key)
	return e
}

func (e *htmlApplet) Align(v string) *htmlApplet {
	e.a["align"] = v
	return e
}

func (e *htmlApplet) Alt(v string) *htmlApplet {
	e.a["alt"] = v
	return e
}

func (e *htmlApplet) Archive(v string) *htmlApplet {
	e.a["archive"] = v
	return e
}

func (e *htmlApplet) Code(v string) *htmlApplet {
	e.a["code"] = v
	return e
}

func (e *htmlApplet) CodeBase(v string) *htmlApplet {
	e.a["codebase"] = v
	return e
}

func (e *htmlApplet) Height(v string) *htmlApplet {
	e.a["height"] = v
	return e
}

func (e *htmlApplet) Hspace(v int) *htmlApplet {
	e.a["hspace"] = v
	return e
}

func (e *htmlApplet) Name(v string) *htmlApplet {
	e.a["name"] = v
	return e
}

func (e *htmlApplet) Object(v string) *htmlApplet {
	e.a["object"] = v
	return e
}

func (e *htmlApplet) Vspace(v int) *htmlApplet {
	e.a["vspace"] = v
	return e
}

func (e *htmlApplet) Width(v string) *htmlApplet {
	e.a["width"] = v
	return e
}

func (e *htmlApplet) ID(v string) *htmlApplet {
	e.a["id"] = v
	return e
}

func (e *htmlApplet) Class(v string) *htmlApplet {
	e.a["class"] = v
	return e
}

func (e *htmlApplet) Title(v string) *htmlApplet {
	e.a["title"] = v
	return e
}

func (e *htmlApplet) Lang(v string) *htmlApplet {
	e.a["lang"] = v
	return e
}

func (e *htmlApplet) Translate(v bool) *htmlApplet {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlApplet) Dir(v string) *htmlApplet {
	e.a["dir"] = v
	return e
}

func (e *htmlApplet) Hidden(v bool) *htmlApplet {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlApplet) TabIndex(v int) *htmlApplet {
	e.a["tabindex"] = v
	return e
}

func (e *htmlApplet) AccessKey(v string) *htmlApplet {
	e.a["accesskey"] = v
	return e
}

func (e *htmlApplet) Draggable(v bool) *htmlApplet {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlApplet) Spellcheck(v bool) *htmlApplet {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
