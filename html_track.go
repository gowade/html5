package h

// HTMLTrack represents HTML <track> tag
type HTMLTrack struct {
	HTMLElement
}

// Track creates a HTML <track> tag
func Track() *HTMLTrack {
	e := &HTMLTrack{}
	e.a = make(map[string]interface{})
	e.tagName = "track"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTrack) S(style StyleMap) *HTMLTrack {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTrack) Key(key interface{}) *HTMLTrack {
	e.key = F(key)
	return e
}

// Kind sets the element's "kind" attribute
func (e *HTMLTrack) Kind(v string) *HTMLTrack {
	e.a["kind"] = v
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLTrack) Src(v string) *HTMLTrack {
	e.a["src"] = v
	return e
}

// Srclang sets the element's "srclang" attribute
func (e *HTMLTrack) Srclang(v string) *HTMLTrack {
	e.a["srclang"] = v
	return e
}

// Label sets the element's "label" attribute
func (e *HTMLTrack) Label(v string) *HTMLTrack {
	e.a["label"] = v
	return e
}

// Default sets the element's "default" attribute
func (e *HTMLTrack) Default(v bool) *HTMLTrack {
	if v {
		e.a["default"] = ""
	} else {
		delete(e.a, "default")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTrack) ID(v string) *HTMLTrack {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTrack) Class(v string) *HTMLTrack {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTrack) Title(v string) *HTMLTrack {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTrack) Lang(v string) *HTMLTrack {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTrack) Translate(v bool) *HTMLTrack {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTrack) Dir(v string) *HTMLTrack {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTrack) Hidden(v bool) *HTMLTrack {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTrack) TabIndex(v int) *HTMLTrack {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTrack) AccessKey(v string) *HTMLTrack {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTrack) Draggable(v bool) *HTMLTrack {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTrack) Spellcheck(v bool) *HTMLTrack {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
