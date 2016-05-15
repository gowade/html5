package html5

// HTMLScript represents HTML <script> tag
type HTMLScript struct {
	HTMLElement
}

// Script creates an HTML <script> tag element
func Script() *HTMLScript {
	e := &HTMLScript{}
	e.a = make(map[string]interface{})
	e.tagName = "script"
	return e
}

// S sets the element's CSS properties
func (e *HTMLScript) S(style StyleMap) *HTMLScript {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLScript) Key(key interface{}) *HTMLScript {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLScript) Ref(dest *DOMElement) *HTMLScript {
	e.ref = dest
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLScript) Src(v string) *HTMLScript {
	e.a["src"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLScript) Type(v string) *HTMLScript {
	e.a["type"] = v
	return e
}

// Charset sets the element's "charset" attribute
func (e *HTMLScript) Charset(v string) *HTMLScript {
	e.a["charset"] = v
	return e
}

// Async sets the element's "async" attribute
func (e *HTMLScript) Async(v bool) *HTMLScript {
	if v {
		e.a["async"] = ""
	} else {
		delete(e.a, "async")
	}
	return e
}

// Defer sets the element's "defer" attribute
func (e *HTMLScript) Defer(v bool) *HTMLScript {
	if v {
		e.a["defer"] = ""
	} else {
		delete(e.a, "defer")
	}
	return e
}

// CrossOrigin sets the element's "crossorigin" attribute
func (e *HTMLScript) CrossOrigin(v string) *HTMLScript {
	e.a["crossorigin"] = v
	return e
}

// Text sets the element's "text" attribute
func (e *HTMLScript) Text(v string) *HTMLScript {
	e.a["text"] = v
	return e
}

// Nonce sets the element's "nonce" attribute
func (e *HTMLScript) Nonce(v string) *HTMLScript {
	e.a["nonce"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLScript) ID(v string) *HTMLScript {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLScript) Class(v string) *HTMLScript {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLScript) Title(v string) *HTMLScript {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLScript) Lang(v string) *HTMLScript {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLScript) Translate(v bool) *HTMLScript {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLScript) Dir(v string) *HTMLScript {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLScript) Hidden(v bool) *HTMLScript {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLScript) TabIndex(v int) *HTMLScript {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLScript) AccessKey(v string) *HTMLScript {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLScript) Draggable(v bool) *HTMLScript {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLScript) Spellcheck(v bool) *HTMLScript {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
