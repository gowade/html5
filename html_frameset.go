package h

// HTMLFrameSet represents HTML <frameset> tag
type HTMLFrameSet struct {
	HTMLElement
}

// FrameSet creates a HTML <frameset> tag
func FrameSet() *HTMLFrameSet {
	e := &HTMLFrameSet{}
	e.a = make(map[string]interface{})
	e.tagName = "frameset"
	return e
}

// S sets the element's CSS properties
func (e *HTMLFrameSet) S(style StyleMap) *HTMLFrameSet {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLFrameSet) Key(key interface{}) *HTMLFrameSet {
	e.key = F(key)
	return e
}

// Cols sets the element's "cols" attribute
func (e *HTMLFrameSet) Cols(v string) *HTMLFrameSet {
	e.a["cols"] = v
	return e
}

// Rows sets the element's "rows" attribute
func (e *HTMLFrameSet) Rows(v string) *HTMLFrameSet {
	e.a["rows"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLFrameSet) ID(v string) *HTMLFrameSet {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLFrameSet) Class(v string) *HTMLFrameSet {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLFrameSet) Title(v string) *HTMLFrameSet {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLFrameSet) Lang(v string) *HTMLFrameSet {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLFrameSet) Translate(v bool) *HTMLFrameSet {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLFrameSet) Dir(v string) *HTMLFrameSet {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLFrameSet) Hidden(v bool) *HTMLFrameSet {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLFrameSet) TabIndex(v int) *HTMLFrameSet {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLFrameSet) AccessKey(v string) *HTMLFrameSet {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLFrameSet) Draggable(v bool) *HTMLFrameSet {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLFrameSet) Spellcheck(v bool) *HTMLFrameSet {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
