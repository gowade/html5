package h

// HTMLIFrame represents HTML <iframe> tag
type HTMLIFrame struct {
	HTMLElement
}

// IFrame creates a HTML <iframe> tag
func IFrame() *HTMLIFrame {
	e := &HTMLIFrame{}
	e.a = make(map[string]interface{})
	e.tagName = "iframe"
	return e
}

// S sets the element's CSS properties
func (e *HTMLIFrame) S(style StyleMap) *HTMLIFrame {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLIFrame) Key(key interface{}) *HTMLIFrame {
	e.key = F(key)
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLIFrame) Src(v string) *HTMLIFrame {
	e.a["src"] = v
	return e
}

// Srcdoc sets the element's "srcdoc" attribute
func (e *HTMLIFrame) Srcdoc(v string) *HTMLIFrame {
	e.a["srcdoc"] = v
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLIFrame) Name(v string) *HTMLIFrame {
	e.a["name"] = v
	return e
}

// AllowFullscreen sets the element's "allowfullscreen" attribute
func (e *HTMLIFrame) AllowFullscreen(v bool) *HTMLIFrame {
	if v {
		e.a["allowfullscreen"] = ""
	} else {
		delete(e.a, "allowfullscreen")
	}
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLIFrame) Width(v string) *HTMLIFrame {
	e.a["width"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLIFrame) Height(v string) *HTMLIFrame {
	e.a["height"] = v
	return e
}

// ReferrerPolicy sets the element's "referrerpolicy" attribute
func (e *HTMLIFrame) ReferrerPolicy(v string) *HTMLIFrame {
	e.a["referrerpolicy"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLIFrame) ID(v string) *HTMLIFrame {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLIFrame) Class(v string) *HTMLIFrame {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLIFrame) Title(v string) *HTMLIFrame {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLIFrame) Lang(v string) *HTMLIFrame {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLIFrame) Translate(v bool) *HTMLIFrame {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLIFrame) Dir(v string) *HTMLIFrame {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLIFrame) Hidden(v bool) *HTMLIFrame {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLIFrame) TabIndex(v int) *HTMLIFrame {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLIFrame) AccessKey(v string) *HTMLIFrame {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLIFrame) Draggable(v bool) *HTMLIFrame {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLIFrame) Spellcheck(v bool) *HTMLIFrame {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
