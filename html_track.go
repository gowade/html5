package h

type htmlTrack struct {
	htmlElement
}

func Track() *htmlTrack {
	e := &htmlTrack{}
	e.a = make(map[string]interface{})
	e.tagName = "track"
	return e
}

func (e *htmlTrack) S(style StyleMap) *htmlTrack {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTrack) Key(key interface{}) *htmlTrack {
	e.key = F(key)
	return e
}

func (e *htmlTrack) Kind(v string) *htmlTrack {
	e.a["kind"] = v
	return e
}

func (e *htmlTrack) Src(v string) *htmlTrack {
	e.a["src"] = v
	return e
}

func (e *htmlTrack) Srclang(v string) *htmlTrack {
	e.a["srclang"] = v
	return e
}

func (e *htmlTrack) Label(v string) *htmlTrack {
	e.a["label"] = v
	return e
}

func (e *htmlTrack) Default(v bool) *htmlTrack {
	if v {
		e.a["default"] = ""
	} else {
		delete(e.a, "default")
	}
	return e
}

func (e *htmlTrack) ID(v string) *htmlTrack {
	e.a["id"] = v
	return e
}

func (e *htmlTrack) Class(v string) *htmlTrack {
	e.a["class"] = v
	return e
}

func (e *htmlTrack) Title(v string) *htmlTrack {
	e.a["title"] = v
	return e
}

func (e *htmlTrack) Lang(v string) *htmlTrack {
	e.a["lang"] = v
	return e
}

func (e *htmlTrack) Translate(v bool) *htmlTrack {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTrack) Dir(v string) *htmlTrack {
	e.a["dir"] = v
	return e
}

func (e *htmlTrack) Hidden(v bool) *htmlTrack {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTrack) TabIndex(v int) *htmlTrack {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTrack) AccessKey(v string) *htmlTrack {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTrack) Draggable(v bool) *htmlTrack {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTrack) Spellcheck(v bool) *htmlTrack {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
