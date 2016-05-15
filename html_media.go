package html5

// HTMLMedia represents HTML <media> tag
type HTMLMedia struct {
	HTMLElement
}

// Media creates an HTML <media> tag element
func Media() *HTMLMedia {
	e := &HTMLMedia{}
	e.a = make(map[string]interface{})
	e.tagName = "media"
	return e
}

// S sets the element's CSS properties
func (e *HTMLMedia) S(style StyleMap) *HTMLMedia {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLMedia) Key(key interface{}) *HTMLMedia {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLMedia) Ref(dest *DOMElement) *HTMLMedia {
	e.ref = dest
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLMedia) Src(v string) *HTMLMedia {
	e.a["src"] = v
	return e
}

// CrossOrigin sets the element's "crossorigin" attribute
func (e *HTMLMedia) CrossOrigin(v string) *HTMLMedia {
	e.a["crossorigin"] = v
	return e
}

// Preload sets the element's "preload" attribute
func (e *HTMLMedia) Preload(v string) *HTMLMedia {
	e.a["preload"] = v
	return e
}

// Autoplay sets the element's "autoplay" attribute
func (e *HTMLMedia) Autoplay(v bool) *HTMLMedia {
	if v {
		e.a["autoplay"] = ""
	} else {
		delete(e.a, "autoplay")
	}
	return e
}

// Loop sets the element's "loop" attribute
func (e *HTMLMedia) Loop(v bool) *HTMLMedia {
	if v {
		e.a["loop"] = ""
	} else {
		delete(e.a, "loop")
	}
	return e
}

// Controls sets the element's "controls" attribute
func (e *HTMLMedia) Controls(v bool) *HTMLMedia {
	if v {
		e.a["controls"] = ""
	} else {
		delete(e.a, "controls")
	}
	return e
}

// Muted sets the element's "muted" attribute
func (e *HTMLMedia) Muted(v bool) *HTMLMedia {
	if v {
		e.a["muted"] = ""
	} else {
		delete(e.a, "muted")
	}
	return e
}

// DefaultMuted sets the element's "defaultmuted" attribute
func (e *HTMLMedia) DefaultMuted(v bool) *HTMLMedia {
	if v {
		e.a["defaultmuted"] = ""
	} else {
		delete(e.a, "defaultmuted")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLMedia) ID(v string) *HTMLMedia {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLMedia) Class(v string) *HTMLMedia {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLMedia) Title(v string) *HTMLMedia {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLMedia) Lang(v string) *HTMLMedia {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLMedia) Translate(v bool) *HTMLMedia {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLMedia) Dir(v string) *HTMLMedia {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLMedia) Hidden(v bool) *HTMLMedia {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLMedia) TabIndex(v int) *HTMLMedia {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLMedia) AccessKey(v string) *HTMLMedia {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLMedia) Draggable(v bool) *HTMLMedia {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLMedia) Spellcheck(v bool) *HTMLMedia {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
