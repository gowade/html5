package html5

// HTMLFont represents HTML <font> tag
type HTMLFont struct {
	HTMLElement
}

// Font creates an HTML <font> tag element
func Font() *HTMLFont {
	e := &HTMLFont{}
	e.a = make(map[string]interface{})
	e.tagName = "font"
	return e
}

// S sets the element's CSS properties
func (e *HTMLFont) S(style StyleMap) *HTMLFont {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLFont) Key(key interface{}) *HTMLFont {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLFont) Ref(dest *DOMElement) *HTMLFont {
	e.ref = dest
	return e
}

// Color sets the element's "color" attribute
func (e *HTMLFont) Color(v string) *HTMLFont {
	e.a["color"] = v
	return e
}

// Face sets the element's "face" attribute
func (e *HTMLFont) Face(v string) *HTMLFont {
	e.a["face"] = v
	return e
}

// Size sets the element's "size" attribute
func (e *HTMLFont) Size(v string) *HTMLFont {
	e.a["size"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLFont) ID(v string) *HTMLFont {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLFont) Class(v string) *HTMLFont {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLFont) Title(v string) *HTMLFont {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLFont) Lang(v string) *HTMLFont {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLFont) Translate(v bool) *HTMLFont {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLFont) Dir(v string) *HTMLFont {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLFont) Hidden(v bool) *HTMLFont {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLFont) TabIndex(v int) *HTMLFont {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLFont) AccessKey(v string) *HTMLFont {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLFont) Draggable(v bool) *HTMLFont {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLFont) Spellcheck(v bool) *HTMLFont {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
