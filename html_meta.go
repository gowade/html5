package h

// HTMLMeta represents HTML <meta> tag
type HTMLMeta struct {
	HTMLElement
}

// Meta creates a HTML <meta> tag
func Meta() *HTMLMeta {
	e := &HTMLMeta{}
	e.a = make(map[string]interface{})
	e.tagName = "meta"
	return e
}

// S sets the element's CSS properties
func (e *HTMLMeta) S(style StyleMap) *HTMLMeta {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLMeta) Key(key interface{}) *HTMLMeta {
	e.key = F(key)
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLMeta) Name(v string) *HTMLMeta {
	e.a["name"] = v
	return e
}

// HttpEquiv sets the element's "httpequiv" attribute
func (e *HTMLMeta) HttpEquiv(v string) *HTMLMeta {
	e.a["httpequiv"] = v
	return e
}

// Content sets the element's "content" attribute
func (e *HTMLMeta) Content(v string) *HTMLMeta {
	e.a["content"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLMeta) ID(v string) *HTMLMeta {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLMeta) Class(v string) *HTMLMeta {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLMeta) Title(v string) *HTMLMeta {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLMeta) Lang(v string) *HTMLMeta {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLMeta) Translate(v bool) *HTMLMeta {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLMeta) Dir(v string) *HTMLMeta {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLMeta) Hidden(v bool) *HTMLMeta {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLMeta) TabIndex(v int) *HTMLMeta {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLMeta) AccessKey(v string) *HTMLMeta {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLMeta) Draggable(v bool) *HTMLMeta {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLMeta) Spellcheck(v bool) *HTMLMeta {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
