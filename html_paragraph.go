package h

type htmlParagraph struct {
	htmlElement
}

func Paragraph() *htmlParagraph {
	e := &htmlParagraph{}
	e.a = make(map[string]interface{})
	e.tagName = "paragraph"
	return e
}

func (e *htmlParagraph) S(style StyleMap) *htmlParagraph {
	e.htmlElement.S(style)
	return e
}

func (e *htmlParagraph) Key(key interface{}) *htmlParagraph {
	e.key = F(key)
	return e
}

func (e *htmlParagraph) ID(v string) *htmlParagraph {
	e.a["id"] = v
	return e
}

func (e *htmlParagraph) Class(v string) *htmlParagraph {
	e.a["class"] = v
	return e
}

func (e *htmlParagraph) Title(v string) *htmlParagraph {
	e.a["title"] = v
	return e
}

func (e *htmlParagraph) Lang(v string) *htmlParagraph {
	e.a["lang"] = v
	return e
}

func (e *htmlParagraph) Translate(v bool) *htmlParagraph {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlParagraph) Dir(v string) *htmlParagraph {
	e.a["dir"] = v
	return e
}

func (e *htmlParagraph) Hidden(v bool) *htmlParagraph {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlParagraph) TabIndex(v int) *htmlParagraph {
	e.a["tabindex"] = v
	return e
}

func (e *htmlParagraph) AccessKey(v string) *htmlParagraph {
	e.a["accesskey"] = v
	return e
}

func (e *htmlParagraph) Draggable(v bool) *htmlParagraph {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlParagraph) Spellcheck(v bool) *htmlParagraph {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
