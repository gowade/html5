package html5

// HTMLAudio represents HTML <audio> tag
type HTMLAudio struct {
	HTMLMedia
}

// Audio creates an HTML <audio> tag element
func Audio() *HTMLAudio {
	e := &HTMLAudio{}
	e.a = make(map[string]interface{})
	e.tagName = "audio"
	return e
}

// S sets the element's CSS properties
func (e *HTMLAudio) S(style StyleMap) *HTMLAudio {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLAudio) Key(key interface{}) *HTMLAudio {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLAudio) Ref(dest *DOMElement) *HTMLAudio {
	e.ref = dest
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLAudio) Src(v string) *HTMLAudio {
	e.a["src"] = v
	return e
}

// CrossOrigin sets the element's "crossorigin" attribute
func (e *HTMLAudio) CrossOrigin(v string) *HTMLAudio {
	e.a["crossorigin"] = v
	return e
}

// Preload sets the element's "preload" attribute
func (e *HTMLAudio) Preload(v string) *HTMLAudio {
	e.a["preload"] = v
	return e
}

// Autoplay sets the element's "autoplay" attribute
func (e *HTMLAudio) Autoplay(v bool) *HTMLAudio {
	if v {
		e.a["autoplay"] = ""
	} else {
		delete(e.a, "autoplay")
	}
	return e
}

// Loop sets the element's "loop" attribute
func (e *HTMLAudio) Loop(v bool) *HTMLAudio {
	if v {
		e.a["loop"] = ""
	} else {
		delete(e.a, "loop")
	}
	return e
}

// Controls sets the element's "controls" attribute
func (e *HTMLAudio) Controls(v bool) *HTMLAudio {
	if v {
		e.a["controls"] = ""
	} else {
		delete(e.a, "controls")
	}
	return e
}

// Muted sets the element's "muted" attribute
func (e *HTMLAudio) Muted(v bool) *HTMLAudio {
	if v {
		e.a["muted"] = ""
	} else {
		delete(e.a, "muted")
	}
	return e
}

// DefaultMuted sets the element's "defaultmuted" attribute
func (e *HTMLAudio) DefaultMuted(v bool) *HTMLAudio {
	if v {
		e.a["defaultmuted"] = ""
	} else {
		delete(e.a, "defaultmuted")
	}
	return e
}
