package h

// HTMLA represents HTML <a> tag
type HTMLA struct {
	HTMLElement
}

// A creates a HTML <a> tag
func A() *HTMLA {
	e := &HTMLA{}
	e.a = make(map[string]interface{})
	e.tagName = "a"
	return e
}

// S sets the element's CSS properties
func (e *HTMLA) S(style StyleMap) *HTMLA {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLA) Key(key interface{}) *HTMLA {
	e.key = F(key)
	return e
}

// Target sets the element's "target" attribute
func (e *HTMLA) Target(v string) *HTMLA {
	e.a["target"] = v
	return e
}

// Download sets the element's "download" attribute
func (e *HTMLA) Download(v string) *HTMLA {
	e.a["download"] = v
	return e
}

// Rel sets the element's "rel" attribute
func (e *HTMLA) Rel(v string) *HTMLA {
	e.a["rel"] = v
	return e
}

// Hreflang sets the element's "hreflang" attribute
func (e *HTMLA) Hreflang(v string) *HTMLA {
	e.a["hreflang"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLA) Type(v string) *HTMLA {
	e.a["type"] = v
	return e
}

// Text sets the element's "text" attribute
func (e *HTMLA) Text(v string) *HTMLA {
	e.a["text"] = v
	return e
}

// ReferrerPolicy sets the element's "referrerpolicy" attribute
func (e *HTMLA) ReferrerPolicy(v string) *HTMLA {
	e.a["referrerpolicy"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLA) ID(v string) *HTMLA {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLA) Class(v string) *HTMLA {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLA) Title(v string) *HTMLA {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLA) Lang(v string) *HTMLA {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLA) Translate(v bool) *HTMLA {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLA) Dir(v string) *HTMLA {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLA) Hidden(v bool) *HTMLA {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLA) TabIndex(v int) *HTMLA {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLA) AccessKey(v string) *HTMLA {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLA) Draggable(v bool) *HTMLA {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLA) Spellcheck(v bool) *HTMLA {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
