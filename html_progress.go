package h

type htmlProgress struct {
	htmlElement
}

func Progress() *htmlProgress {
	e := &htmlProgress{}
	e.a = make(map[string]interface{})
	e.tagName = "progress"
	return e
}

func (e *htmlProgress) S(style StyleMap) *htmlProgress {
	e.htmlElement.S(style)
	return e
}

func (e *htmlProgress) Key(key interface{}) *htmlProgress {
	e.key = F(key)
	return e
}

func (e *htmlProgress) ID(v string) *htmlProgress {
	e.a["id"] = v
	return e
}

func (e *htmlProgress) Class(v string) *htmlProgress {
	e.a["class"] = v
	return e
}

func (e *htmlProgress) Title(v string) *htmlProgress {
	e.a["title"] = v
	return e
}

func (e *htmlProgress) Lang(v string) *htmlProgress {
	e.a["lang"] = v
	return e
}

func (e *htmlProgress) Translate(v bool) *htmlProgress {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlProgress) Dir(v string) *htmlProgress {
	e.a["dir"] = v
	return e
}

func (e *htmlProgress) Hidden(v bool) *htmlProgress {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlProgress) TabIndex(v int) *htmlProgress {
	e.a["tabindex"] = v
	return e
}

func (e *htmlProgress) AccessKey(v string) *htmlProgress {
	e.a["accesskey"] = v
	return e
}

func (e *htmlProgress) Draggable(v bool) *htmlProgress {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlProgress) Spellcheck(v bool) *htmlProgress {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
