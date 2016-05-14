package h

type htmlMedia struct {
	htmlElement
}

func Media() *htmlMedia {
	e := &htmlMedia{}
	e.a = make(map[string]interface{})
	e.tagName = "media"
	return e
}

func (e *htmlMedia) S(style StyleMap) *htmlMedia {
	e.htmlElement.S(style)
	return e
}

func (e *htmlMedia) Key(key interface{}) *htmlMedia {
	e.key = F(key)
	return e
}

func (e *htmlMedia) Src(v string) *htmlMedia {
	e.a["src"] = v
	return e
}

func (e *htmlMedia) CrossOrigin(v string) *htmlMedia {
	e.a["crossorigin"] = v
	return e
}

func (e *htmlMedia) Preload(v string) *htmlMedia {
	e.a["preload"] = v
	return e
}

func (e *htmlMedia) Autoplay(v bool) *htmlMedia {
	if v {
		e.a["autoplay"] = ""
	} else {
		delete(e.a, "autoplay")
	}
	return e
}

func (e *htmlMedia) Loop(v bool) *htmlMedia {
	if v {
		e.a["loop"] = ""
	} else {
		delete(e.a, "loop")
	}
	return e
}

func (e *htmlMedia) Controls(v bool) *htmlMedia {
	if v {
		e.a["controls"] = ""
	} else {
		delete(e.a, "controls")
	}
	return e
}

func (e *htmlMedia) Muted(v bool) *htmlMedia {
	if v {
		e.a["muted"] = ""
	} else {
		delete(e.a, "muted")
	}
	return e
}

func (e *htmlMedia) DefaultMuted(v bool) *htmlMedia {
	if v {
		e.a["defaultmuted"] = ""
	} else {
		delete(e.a, "defaultmuted")
	}
	return e
}

func (e *htmlMedia) ID(v string) *htmlMedia {
	e.a["id"] = v
	return e
}

func (e *htmlMedia) Class(v string) *htmlMedia {
	e.a["class"] = v
	return e
}

func (e *htmlMedia) Title(v string) *htmlMedia {
	e.a["title"] = v
	return e
}

func (e *htmlMedia) Lang(v string) *htmlMedia {
	e.a["lang"] = v
	return e
}

func (e *htmlMedia) Translate(v bool) *htmlMedia {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlMedia) Dir(v string) *htmlMedia {
	e.a["dir"] = v
	return e
}

func (e *htmlMedia) Hidden(v bool) *htmlMedia {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlMedia) TabIndex(v int) *htmlMedia {
	e.a["tabindex"] = v
	return e
}

func (e *htmlMedia) AccessKey(v string) *htmlMedia {
	e.a["accesskey"] = v
	return e
}

func (e *htmlMedia) Draggable(v bool) *htmlMedia {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlMedia) Spellcheck(v bool) *htmlMedia {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
