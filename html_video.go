package h

// HTMLVideo represents HTML <video> tag
type HTMLVideo struct {
	HTMLMedia
}

// Video creates an HTML <video> tag element
func Video() *HTMLVideo {
	e := &HTMLVideo{}
	e.a = make(map[string]interface{})
	e.tagName = "video"
	return e
}

// S sets the element's CSS properties
func (e *HTMLVideo) S(style StyleMap) *HTMLVideo {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLVideo) Key(key interface{}) *HTMLVideo {
	e.key = F(key)
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLVideo) Width(v int) *HTMLVideo {
	e.a["width"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLVideo) Height(v int) *HTMLVideo {
	e.a["height"] = v
	return e
}

// Poster sets the element's "poster" attribute
func (e *HTMLVideo) Poster(v string) *HTMLVideo {
	e.a["poster"] = v
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLVideo) Src(v string) *HTMLVideo {
	e.a["src"] = v
	return e
}

// CrossOrigin sets the element's "crossorigin" attribute
func (e *HTMLVideo) CrossOrigin(v string) *HTMLVideo {
	e.a["crossorigin"] = v
	return e
}

// Preload sets the element's "preload" attribute
func (e *HTMLVideo) Preload(v string) *HTMLVideo {
	e.a["preload"] = v
	return e
}

// Autoplay sets the element's "autoplay" attribute
func (e *HTMLVideo) Autoplay(v bool) *HTMLVideo {
	if v {
		e.a["autoplay"] = ""
	} else {
		delete(e.a, "autoplay")
	}
	return e
}

// Loop sets the element's "loop" attribute
func (e *HTMLVideo) Loop(v bool) *HTMLVideo {
	if v {
		e.a["loop"] = ""
	} else {
		delete(e.a, "loop")
	}
	return e
}

// Controls sets the element's "controls" attribute
func (e *HTMLVideo) Controls(v bool) *HTMLVideo {
	if v {
		e.a["controls"] = ""
	} else {
		delete(e.a, "controls")
	}
	return e
}

// Muted sets the element's "muted" attribute
func (e *HTMLVideo) Muted(v bool) *HTMLVideo {
	if v {
		e.a["muted"] = ""
	} else {
		delete(e.a, "muted")
	}
	return e
}

// DefaultMuted sets the element's "defaultmuted" attribute
func (e *HTMLVideo) DefaultMuted(v bool) *HTMLVideo {
	if v {
		e.a["defaultmuted"] = ""
	} else {
		delete(e.a, "defaultmuted")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLVideo) ID(v string) *HTMLVideo {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLVideo) Class(v string) *HTMLVideo {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLVideo) Title(v string) *HTMLVideo {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLVideo) Lang(v string) *HTMLVideo {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLVideo) Translate(v bool) *HTMLVideo {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLVideo) Dir(v string) *HTMLVideo {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLVideo) Hidden(v bool) *HTMLVideo {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLVideo) TabIndex(v int) *HTMLVideo {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLVideo) AccessKey(v string) *HTMLVideo {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLVideo) Draggable(v bool) *HTMLVideo {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLVideo) Spellcheck(v bool) *HTMLVideo {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
