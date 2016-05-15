package h

// HTMLBase represents HTML <base> tag
type HTMLBase struct {
	HTMLElement
}

// Base creates an HTML <base> tag element
func Base() *HTMLBase {
	e := &HTMLBase{}
	e.a = make(map[string]interface{})
	e.tagName = "base"
	return e
}

// S sets the element's CSS properties
func (e *HTMLBase) S(style StyleMap) *HTMLBase {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLBase) Key(key interface{}) *HTMLBase {
	e.key = F(key)
	return e
}

// Href sets the element's "href" attribute
func (e *HTMLBase) Href(v string) *HTMLBase {
	e.a["href"] = v
	return e
}

// Target sets the element's "target" attribute
func (e *HTMLBase) Target(v string) *HTMLBase {
	e.a["target"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLBase) ID(v string) *HTMLBase {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLBase) Class(v string) *HTMLBase {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLBase) Title(v string) *HTMLBase {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLBase) Lang(v string) *HTMLBase {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLBase) Translate(v bool) *HTMLBase {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLBase) Dir(v string) *HTMLBase {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLBase) Hidden(v bool) *HTMLBase {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLBase) TabIndex(v int) *HTMLBase {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLBase) AccessKey(v string) *HTMLBase {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLBase) Draggable(v bool) *HTMLBase {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLBase) Spellcheck(v bool) *HTMLBase {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
