package h

// HTMLSource represents HTML <source> tag
type HTMLSource struct {
	HTMLElement
}

// Source creates an HTML <source> tag element
func Source() *HTMLSource {
	e := &HTMLSource{}
	e.a = make(map[string]interface{})
	e.tagName = "source"
	return e
}

// S sets the element's CSS properties
func (e *HTMLSource) S(style StyleMap) *HTMLSource {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLSource) Key(key interface{}) *HTMLSource {
	e.key = F(key)
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLSource) Src(v string) *HTMLSource {
	e.a["src"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLSource) Type(v string) *HTMLSource {
	e.a["type"] = v
	return e
}

// Srcset sets the element's "srcset" attribute
func (e *HTMLSource) Srcset(v string) *HTMLSource {
	e.a["srcset"] = v
	return e
}

// Sizes sets the element's "sizes" attribute
func (e *HTMLSource) Sizes(v string) *HTMLSource {
	e.a["sizes"] = v
	return e
}

// Media sets the element's "media" attribute
func (e *HTMLSource) Media(v string) *HTMLSource {
	e.a["media"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLSource) ID(v string) *HTMLSource {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLSource) Class(v string) *HTMLSource {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLSource) Title(v string) *HTMLSource {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLSource) Lang(v string) *HTMLSource {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLSource) Translate(v bool) *HTMLSource {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLSource) Dir(v string) *HTMLSource {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLSource) Hidden(v bool) *HTMLSource {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLSource) TabIndex(v int) *HTMLSource {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLSource) AccessKey(v string) *HTMLSource {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLSource) Draggable(v bool) *HTMLSource {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLSource) Spellcheck(v bool) *HTMLSource {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
