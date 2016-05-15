package h

// HTMLStyle represents HTML <style> tag
type HTMLStyle struct {
	HTMLElement
}

// Style creates a HTML <style> tag
func Style() *HTMLStyle {
	e := &HTMLStyle{}
	e.a = make(map[string]interface{})
	e.tagName = "style"
	return e
}

// S sets the element's CSS properties
func (e *HTMLStyle) S(style StyleMap) *HTMLStyle {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLStyle) Key(key interface{}) *HTMLStyle {
	e.key = F(key)
	return e
}

// Media sets the element's "media" attribute
func (e *HTMLStyle) Media(v string) *HTMLStyle {
	e.a["media"] = v
	return e
}

// Nonce sets the element's "nonce" attribute
func (e *HTMLStyle) Nonce(v string) *HTMLStyle {
	e.a["nonce"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLStyle) Type(v string) *HTMLStyle {
	e.a["type"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLStyle) ID(v string) *HTMLStyle {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLStyle) Class(v string) *HTMLStyle {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLStyle) Title(v string) *HTMLStyle {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLStyle) Lang(v string) *HTMLStyle {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLStyle) Translate(v bool) *HTMLStyle {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLStyle) Dir(v string) *HTMLStyle {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLStyle) Hidden(v bool) *HTMLStyle {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLStyle) TabIndex(v int) *HTMLStyle {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLStyle) AccessKey(v string) *HTMLStyle {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLStyle) Draggable(v bool) *HTMLStyle {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLStyle) Spellcheck(v bool) *HTMLStyle {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
