package h

// HTMLTime represents HTML <time> tag
type HTMLTime struct {
	HTMLElement
}

// Time creates a HTML <time> tag
func Time() *HTMLTime {
	e := &HTMLTime{}
	e.a = make(map[string]interface{})
	e.tagName = "time"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTime) S(style StyleMap) *HTMLTime {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTime) Key(key interface{}) *HTMLTime {
	e.key = F(key)
	return e
}

// DateTime sets the element's "datetime" attribute
func (e *HTMLTime) DateTime(v string) *HTMLTime {
	e.a["datetime"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTime) ID(v string) *HTMLTime {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTime) Class(v string) *HTMLTime {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTime) Title(v string) *HTMLTime {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTime) Lang(v string) *HTMLTime {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTime) Translate(v bool) *HTMLTime {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTime) Dir(v string) *HTMLTime {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTime) Hidden(v bool) *HTMLTime {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTime) TabIndex(v int) *HTMLTime {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTime) AccessKey(v string) *HTMLTime {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTime) Draggable(v bool) *HTMLTime {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTime) Spellcheck(v bool) *HTMLTime {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
