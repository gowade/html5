package h

// HTMLImage represents HTML <image> tag
type HTMLImage struct {
	HTMLElement
}

// Image creates a HTML <image> tag
func Image() *HTMLImage {
	e := &HTMLImage{}
	e.a = make(map[string]interface{})
	e.tagName = "image"
	return e
}

// S sets the element's CSS properties
func (e *HTMLImage) S(style StyleMap) *HTMLImage {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLImage) Key(key interface{}) *HTMLImage {
	e.key = F(key)
	return e
}

// Alt sets the element's "alt" attribute
func (e *HTMLImage) Alt(v string) *HTMLImage {
	e.a["alt"] = v
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLImage) Src(v string) *HTMLImage {
	e.a["src"] = v
	return e
}

// Srcset sets the element's "srcset" attribute
func (e *HTMLImage) Srcset(v string) *HTMLImage {
	e.a["srcset"] = v
	return e
}

// Sizes sets the element's "sizes" attribute
func (e *HTMLImage) Sizes(v string) *HTMLImage {
	e.a["sizes"] = v
	return e
}

// CrossOrigin sets the element's "crossorigin" attribute
func (e *HTMLImage) CrossOrigin(v string) *HTMLImage {
	e.a["crossorigin"] = v
	return e
}

// UseMap sets the element's "usemap" attribute
func (e *HTMLImage) UseMap(v string) *HTMLImage {
	e.a["usemap"] = v
	return e
}

// IsMap sets the element's "ismap" attribute
func (e *HTMLImage) IsMap(v bool) *HTMLImage {
	if v {
		e.a["ismap"] = ""
	} else {
		delete(e.a, "ismap")
	}
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLImage) Width(v int) *HTMLImage {
	e.a["width"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLImage) Height(v int) *HTMLImage {
	e.a["height"] = v
	return e
}

// ReferrerPolicy sets the element's "referrerpolicy" attribute
func (e *HTMLImage) ReferrerPolicy(v string) *HTMLImage {
	e.a["referrerpolicy"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLImage) ID(v string) *HTMLImage {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLImage) Class(v string) *HTMLImage {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLImage) Title(v string) *HTMLImage {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLImage) Lang(v string) *HTMLImage {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLImage) Translate(v bool) *HTMLImage {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLImage) Dir(v string) *HTMLImage {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLImage) Hidden(v bool) *HTMLImage {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLImage) TabIndex(v int) *HTMLImage {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLImage) AccessKey(v string) *HTMLImage {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLImage) Draggable(v bool) *HTMLImage {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLImage) Spellcheck(v bool) *HTMLImage {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
