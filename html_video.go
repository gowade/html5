package html5

// HTMLVideo represents HTML <video> tag
type HTMLVideo struct {
	HTMLMedia
}

// Video creates an HTML <video> tag element
func Video() *HTMLVideo {
	e := &HTMLVideo{}
	e.a = make(map[string]interface{})
	e.tagName = "video"
	return e
}

// S sets the element's CSS properties
func (e *HTMLVideo) S(style StyleMap) *HTMLVideo {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLVideo) Key(key interface{}) *HTMLVideo {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLVideo) Ref(dest *DOMElement) *HTMLVideo {
	e.ref = dest
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLVideo) Width(v int) *HTMLVideo {
	e.a["width"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLVideo) Height(v int) *HTMLVideo {
	e.a["height"] = v
	return e
}

// Poster sets the element's "poster" attribute
func (e *HTMLVideo) Poster(v string) *HTMLVideo {
	e.a["poster"] = v
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLVideo) Src(v string) *HTMLVideo {
	e.a["src"] = v
	return e
}

// CrossOrigin sets the element's "crossorigin" attribute
func (e *HTMLVideo) CrossOrigin(v string) *HTMLVideo {
	e.a["crossorigin"] = v
	return e
}

// Preload sets the element's "preload" attribute
func (e *HTMLVideo) Preload(v string) *HTMLVideo {
	e.a["preload"] = v
	return e
}

// Autoplay sets the element's "autoplay" attribute
func (e *HTMLVideo) Autoplay(v bool) *HTMLVideo {
	if v {
		e.a["autoplay"] = ""
	} else {
		delete(e.a, "autoplay")
	}
	return e
}

// Loop sets the element's "loop" attribute
func (e *HTMLVideo) Loop(v bool) *HTMLVideo {
	if v {
		e.a["loop"] = ""
	} else {
		delete(e.a, "loop")
	}
	return e
}

// Controls sets the element's "controls" attribute
func (e *HTMLVideo) Controls(v bool) *HTMLVideo {
	if v {
		e.a["controls"] = ""
	} else {
		delete(e.a, "controls")
	}
	return e
}

// Muted sets the element's "muted" attribute
func (e *HTMLVideo) Muted(v bool) *HTMLVideo {
	if v {
		e.a["muted"] = ""
	} else {
		delete(e.a, "muted")
	}
	return e
}

// DefaultMuted sets the element's "defaultmuted" attribute
func (e *HTMLVideo) DefaultMuted(v bool) *HTMLVideo {
	if v {
		e.a["defaultmuted"] = ""
	} else {
		delete(e.a, "defaultmuted")
	}
	return e
}
