package h

type htmlVideo struct {
	htmlMedia
}

func Video() *htmlVideo {
	e := &htmlVideo{}
	e.a = make(map[string]interface{})
	e.tagName = "video"
	return e
}

func (e *htmlVideo) S(style StyleMap) *htmlVideo {
	e.htmlElement.S(style)
	return e
}

func (e *htmlVideo) Key(key interface{}) *htmlVideo {
	e.key = F(key)
	return e
}

func (e *htmlVideo) Width(v int) *htmlVideo {
	e.a["width"] = v
	return e
}

func (e *htmlVideo) Height(v int) *htmlVideo {
	e.a["height"] = v
	return e
}

func (e *htmlVideo) Poster(v string) *htmlVideo {
	e.a["poster"] = v
	return e
}

func (e *htmlVideo) Src(v string) *htmlVideo {
	e.a["src"] = v
	return e
}

func (e *htmlVideo) CrossOrigin(v string) *htmlVideo {
	e.a["crossorigin"] = v
	return e
}

func (e *htmlVideo) Preload(v string) *htmlVideo {
	e.a["preload"] = v
	return e
}

func (e *htmlVideo) Autoplay(v bool) *htmlVideo {
	if v {
		e.a["autoplay"] = ""
	} else {
		delete(e.a, "autoplay")
	}
	return e
}

func (e *htmlVideo) Loop(v bool) *htmlVideo {
	if v {
		e.a["loop"] = ""
	} else {
		delete(e.a, "loop")
	}
	return e
}

func (e *htmlVideo) Controls(v bool) *htmlVideo {
	if v {
		e.a["controls"] = ""
	} else {
		delete(e.a, "controls")
	}
	return e
}

func (e *htmlVideo) Muted(v bool) *htmlVideo {
	if v {
		e.a["muted"] = ""
	} else {
		delete(e.a, "muted")
	}
	return e
}

func (e *htmlVideo) DefaultMuted(v bool) *htmlVideo {
	if v {
		e.a["defaultmuted"] = ""
	} else {
		delete(e.a, "defaultmuted")
	}
	return e
}

func (e *htmlVideo) ID(v string) *htmlVideo {
	e.a["id"] = v
	return e
}

func (e *htmlVideo) Class(v string) *htmlVideo {
	e.a["class"] = v
	return e
}

func (e *htmlVideo) Title(v string) *htmlVideo {
	e.a["title"] = v
	return e
}

func (e *htmlVideo) Lang(v string) *htmlVideo {
	e.a["lang"] = v
	return e
}

func (e *htmlVideo) Translate(v bool) *htmlVideo {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlVideo) Dir(v string) *htmlVideo {
	e.a["dir"] = v
	return e
}

func (e *htmlVideo) Hidden(v bool) *htmlVideo {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlVideo) TabIndex(v int) *htmlVideo {
	e.a["tabindex"] = v
	return e
}

func (e *htmlVideo) AccessKey(v string) *htmlVideo {
	e.a["accesskey"] = v
	return e
}

func (e *htmlVideo) Draggable(v bool) *htmlVideo {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlVideo) Spellcheck(v bool) *htmlVideo {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
