package h

// HTMLEmbed represents HTML <embed> tag
type HTMLEmbed struct {
	HTMLElement
}

// Embed creates a HTML <embed> tag
func Embed() *HTMLEmbed {
	e := &HTMLEmbed{}
	e.a = make(map[string]interface{})
	e.tagName = "embed"
	return e
}

// S sets the element's CSS properties
func (e *HTMLEmbed) S(style StyleMap) *HTMLEmbed {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLEmbed) Key(key interface{}) *HTMLEmbed {
	e.key = F(key)
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLEmbed) Src(v string) *HTMLEmbed {
	e.a["src"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLEmbed) Type(v string) *HTMLEmbed {
	e.a["type"] = v
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLEmbed) Width(v string) *HTMLEmbed {
	e.a["width"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLEmbed) Height(v string) *HTMLEmbed {
	e.a["height"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLEmbed) ID(v string) *HTMLEmbed {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLEmbed) Class(v string) *HTMLEmbed {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLEmbed) Title(v string) *HTMLEmbed {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLEmbed) Lang(v string) *HTMLEmbed {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLEmbed) Translate(v bool) *HTMLEmbed {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLEmbed) Dir(v string) *HTMLEmbed {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLEmbed) Hidden(v bool) *HTMLEmbed {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLEmbed) TabIndex(v int) *HTMLEmbed {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLEmbed) AccessKey(v string) *HTMLEmbed {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLEmbed) Draggable(v bool) *HTMLEmbed {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLEmbed) Spellcheck(v bool) *HTMLEmbed {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
