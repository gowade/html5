package h

// HTMLMap represents HTML <map> tag
type HTMLMap struct {
	HTMLElement
}

// Map creates a HTML <map> tag
func Map() *HTMLMap {
	e := &HTMLMap{}
	e.a = make(map[string]interface{})
	e.tagName = "map"
	return e
}

// S sets the element's CSS properties
func (e *HTMLMap) S(style StyleMap) *HTMLMap {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLMap) Key(key interface{}) *HTMLMap {
	e.key = F(key)
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLMap) Name(v string) *HTMLMap {
	e.a["name"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLMap) ID(v string) *HTMLMap {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLMap) Class(v string) *HTMLMap {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLMap) Title(v string) *HTMLMap {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLMap) Lang(v string) *HTMLMap {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLMap) Translate(v bool) *HTMLMap {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLMap) Dir(v string) *HTMLMap {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLMap) Hidden(v bool) *HTMLMap {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLMap) TabIndex(v int) *HTMLMap {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLMap) AccessKey(v string) *HTMLMap {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLMap) Draggable(v bool) *HTMLMap {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLMap) Spellcheck(v bool) *HTMLMap {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
