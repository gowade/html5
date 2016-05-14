package h

type htmlEmbed struct {
	htmlElement
}

func Embed() *htmlEmbed {
	e := &htmlEmbed{}
	e.a = make(map[string]interface{})
	e.tagName = "embed"
	return e
}

func (e *htmlEmbed) S(style StyleMap) *htmlEmbed {
	e.htmlElement.S(style)
	return e
}

func (e *htmlEmbed) Key(key interface{}) *htmlEmbed {
	e.key = F(key)
	return e
}

func (e *htmlEmbed) Src(v string) *htmlEmbed {
	e.a["src"] = v
	return e
}

func (e *htmlEmbed) Type(v string) *htmlEmbed {
	e.a["type"] = v
	return e
}

func (e *htmlEmbed) Width(v string) *htmlEmbed {
	e.a["width"] = v
	return e
}

func (e *htmlEmbed) Height(v string) *htmlEmbed {
	e.a["height"] = v
	return e
}

func (e *htmlEmbed) ID(v string) *htmlEmbed {
	e.a["id"] = v
	return e
}

func (e *htmlEmbed) Class(v string) *htmlEmbed {
	e.a["class"] = v
	return e
}

func (e *htmlEmbed) Title(v string) *htmlEmbed {
	e.a["title"] = v
	return e
}

func (e *htmlEmbed) Lang(v string) *htmlEmbed {
	e.a["lang"] = v
	return e
}

func (e *htmlEmbed) Translate(v bool) *htmlEmbed {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlEmbed) Dir(v string) *htmlEmbed {
	e.a["dir"] = v
	return e
}

func (e *htmlEmbed) Hidden(v bool) *htmlEmbed {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlEmbed) TabIndex(v int) *htmlEmbed {
	e.a["tabindex"] = v
	return e
}

func (e *htmlEmbed) AccessKey(v string) *htmlEmbed {
	e.a["accesskey"] = v
	return e
}

func (e *htmlEmbed) Draggable(v bool) *htmlEmbed {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlEmbed) Spellcheck(v bool) *htmlEmbed {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
