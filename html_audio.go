package h

type htmlAudio struct {
	htmlMedia
}

func Audio() *htmlAudio {
	e := &htmlAudio{}
	e.a = make(map[string]interface{})
	e.tagName = "audio"
	return e
}

func (e *htmlAudio) S(style StyleMap) *htmlAudio {
	e.htmlElement.S(style)
	return e
}

func (e *htmlAudio) Key(key interface{}) *htmlAudio {
	e.key = F(key)
	return e
}

func (e *htmlAudio) Src(v string) *htmlAudio {
	e.a["src"] = v
	return e
}

func (e *htmlAudio) CrossOrigin(v string) *htmlAudio {
	e.a["crossorigin"] = v
	return e
}

func (e *htmlAudio) Preload(v string) *htmlAudio {
	e.a["preload"] = v
	return e
}

func (e *htmlAudio) Autoplay(v bool) *htmlAudio {
	if v {
		e.a["autoplay"] = ""
	} else {
		delete(e.a, "autoplay")
	}
	return e
}

func (e *htmlAudio) Loop(v bool) *htmlAudio {
	if v {
		e.a["loop"] = ""
	} else {
		delete(e.a, "loop")
	}
	return e
}

func (e *htmlAudio) Controls(v bool) *htmlAudio {
	if v {
		e.a["controls"] = ""
	} else {
		delete(e.a, "controls")
	}
	return e
}

func (e *htmlAudio) Muted(v bool) *htmlAudio {
	if v {
		e.a["muted"] = ""
	} else {
		delete(e.a, "muted")
	}
	return e
}

func (e *htmlAudio) DefaultMuted(v bool) *htmlAudio {
	if v {
		e.a["defaultmuted"] = ""
	} else {
		delete(e.a, "defaultmuted")
	}
	return e
}

func (e *htmlAudio) ID(v string) *htmlAudio {
	e.a["id"] = v
	return e
}

func (e *htmlAudio) Class(v string) *htmlAudio {
	e.a["class"] = v
	return e
}

func (e *htmlAudio) Title(v string) *htmlAudio {
	e.a["title"] = v
	return e
}

func (e *htmlAudio) Lang(v string) *htmlAudio {
	e.a["lang"] = v
	return e
}

func (e *htmlAudio) Translate(v bool) *htmlAudio {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlAudio) Dir(v string) *htmlAudio {
	e.a["dir"] = v
	return e
}

func (e *htmlAudio) Hidden(v bool) *htmlAudio {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlAudio) TabIndex(v int) *htmlAudio {
	e.a["tabindex"] = v
	return e
}

func (e *htmlAudio) AccessKey(v string) *htmlAudio {
	e.a["accesskey"] = v
	return e
}

func (e *htmlAudio) Draggable(v bool) *htmlAudio {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlAudio) Spellcheck(v bool) *htmlAudio {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
