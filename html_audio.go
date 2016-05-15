package h

// HTMLAudio represents HTML <audio> tag
type HTMLAudio struct {
	HTMLMedia
}

// Audio creates a HTML <audio> tag
func Audio() *HTMLAudio {
	e := &HTMLAudio{}
	e.a = make(map[string]interface{})
	e.tagName = "audio"
	return e
}

// S sets the element's CSS properties
func (e *HTMLAudio) S(style StyleMap) *HTMLAudio {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLAudio) Key(key interface{}) *HTMLAudio {
	e.key = F(key)
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLAudio) Src(v string) *HTMLAudio {
	e.a["src"] = v
	return e
}

// CrossOrigin sets the element's "crossorigin" attribute
func (e *HTMLAudio) CrossOrigin(v string) *HTMLAudio {
	e.a["crossorigin"] = v
	return e
}

// Preload sets the element's "preload" attribute
func (e *HTMLAudio) Preload(v string) *HTMLAudio {
	e.a["preload"] = v
	return e
}

// Autoplay sets the element's "autoplay" attribute
func (e *HTMLAudio) Autoplay(v bool) *HTMLAudio {
	if v {
		e.a["autoplay"] = ""
	} else {
		delete(e.a, "autoplay")
	}
	return e
}

// Loop sets the element's "loop" attribute
func (e *HTMLAudio) Loop(v bool) *HTMLAudio {
	if v {
		e.a["loop"] = ""
	} else {
		delete(e.a, "loop")
	}
	return e
}

// Controls sets the element's "controls" attribute
func (e *HTMLAudio) Controls(v bool) *HTMLAudio {
	if v {
		e.a["controls"] = ""
	} else {
		delete(e.a, "controls")
	}
	return e
}

// Muted sets the element's "muted" attribute
func (e *HTMLAudio) Muted(v bool) *HTMLAudio {
	if v {
		e.a["muted"] = ""
	} else {
		delete(e.a, "muted")
	}
	return e
}

// DefaultMuted sets the element's "defaultmuted" attribute
func (e *HTMLAudio) DefaultMuted(v bool) *HTMLAudio {
	if v {
		e.a["defaultmuted"] = ""
	} else {
		delete(e.a, "defaultmuted")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLAudio) ID(v string) *HTMLAudio {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLAudio) Class(v string) *HTMLAudio {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLAudio) Title(v string) *HTMLAudio {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLAudio) Lang(v string) *HTMLAudio {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLAudio) Translate(v bool) *HTMLAudio {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLAudio) Dir(v string) *HTMLAudio {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLAudio) Hidden(v bool) *HTMLAudio {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLAudio) TabIndex(v int) *HTMLAudio {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLAudio) AccessKey(v string) *HTMLAudio {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLAudio) Draggable(v bool) *HTMLAudio {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLAudio) Spellcheck(v bool) *HTMLAudio {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
