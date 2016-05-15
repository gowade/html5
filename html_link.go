package h

// HTMLLink represents HTML <link> tag
type HTMLLink struct {
	HTMLElement
}

// Link creates a HTML <link> tag
func Link() *HTMLLink {
	e := &HTMLLink{}
	e.a = make(map[string]interface{})
	e.tagName = "link"
	return e
}

// S sets the element's CSS properties
func (e *HTMLLink) S(style StyleMap) *HTMLLink {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLLink) Key(key interface{}) *HTMLLink {
	e.key = F(key)
	return e
}

// Href sets the element's "href" attribute
func (e *HTMLLink) Href(v string) *HTMLLink {
	e.a["href"] = v
	return e
}

// CrossOrigin sets the element's "crossorigin" attribute
func (e *HTMLLink) CrossOrigin(v string) *HTMLLink {
	e.a["crossorigin"] = v
	return e
}

// Rel sets the element's "rel" attribute
func (e *HTMLLink) Rel(v string) *HTMLLink {
	e.a["rel"] = v
	return e
}

// Media sets the element's "media" attribute
func (e *HTMLLink) Media(v string) *HTMLLink {
	e.a["media"] = v
	return e
}

// Hreflang sets the element's "hreflang" attribute
func (e *HTMLLink) Hreflang(v string) *HTMLLink {
	e.a["hreflang"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLLink) Type(v string) *HTMLLink {
	e.a["type"] = v
	return e
}

// ReferrerPolicy sets the element's "referrerpolicy" attribute
func (e *HTMLLink) ReferrerPolicy(v string) *HTMLLink {
	e.a["referrerpolicy"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLLink) ID(v string) *HTMLLink {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLLink) Class(v string) *HTMLLink {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLLink) Title(v string) *HTMLLink {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLLink) Lang(v string) *HTMLLink {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLLink) Translate(v bool) *HTMLLink {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLLink) Dir(v string) *HTMLLink {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLLink) Hidden(v bool) *HTMLLink {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLLink) TabIndex(v int) *HTMLLink {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLLink) AccessKey(v string) *HTMLLink {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLLink) Draggable(v bool) *HTMLLink {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLLink) Spellcheck(v bool) *HTMLLink {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
