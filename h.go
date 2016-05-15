package h

import (
	"fmt"

	"github.com/gowade/vdom"
)

type StyleMap map[string]interface{}

type Node interface {
	ToVNode() vdom.VNode
}

type TextNode string

func (t TextNode) ToVNode() vdom.VNode {
	return vdom.VText(t)
}

func (e *HTMLElement) ToVNode() vdom.VNode {
	ve := vdom.VElement{
		TagName:   e.tagName,
		Namespace: e.namespace,
		Key:       e.key,
		Props:     e.a,
		Children:  make([]vdom.VNode, len(e.c)),
	}

	for i := range e.c {
		ve.Children[i] = e.c[i].ToVNode()
	}

	return ve
}

// F is a helper method to make string,
// it works like fmt.Sprintf if the first argument is string,
// otherwise it works like fmt.Sprint
func F(formatOrValue interface{}, args ...interface{}) string {
	if s, ok := formatOrValue.(string); ok {
		if len(args) == 0 {
			return s
		}
		return fmt.Sprintf(s, args...)
	}

	l := make([]interface{}, 0, len(args)+1)
	l = append(l, formatOrValue)
	l = append(l, args...)
	return fmt.Sprint(l...)
}

// T is a helper method to make text node,
// it works like F() but returns TextNode instead of just string
func T(formatOrValue interface{}, args ...interface{}) TextNode {
	return TextNode(F(formatOrValue, args...))
}

type L []Node

// Add adds new nodes into the list
func (l *L) Add(nodes ...Node) {
	*l = append(*l, node...)
}

type HTMLElement struct {
	a         map[string]interface{}
	c         L
	style     map[string]interface{}
	tagName   string
	namespace string
	key       string
}

// C is a helper method to declare the element's children
func (e *HTMLElement) C(children ...Node) *HTMLElement {
	e.c = children
	return e
}

// L takes a function and call it so that it can add children to the element
//
// Typically used to declare children using "for" loops:
//
//  parent.L(func(l *L) {
//  	for _, item := range someCollection {
//  		l.Add(h.Span().T(item))
//  	}
//  })
func (e *HTMLElement) L(iterateFn func(childrenP *L)) *HTMLElement {
	iterateFn(&e.c)
	return e
}

// T is shorthand for C(h.T()),
// it makes the element have a single child that is a text node
func (e *HTMLElement) T(formatOrValue interface{}, args ...interface{}) *HTMLElement {
	return e.C(T(formatOrValue, args...))
}

// A sets value for a custom attribute,
// should only be used for custom, nonstandard attribute or element
func (e *HTMLElement) A(name string, value interface{}) *HTMLElement {
	if valueBool, ok := value.(bool); ok {
		if valueBool {
			e.a[name] = ""
		} else {
			delete(e.a, name)
		}
		return e
	}

	e.a[name] = value
	return e
}

// S sets the element's CSS properties
func (e *HTMLElement) S(style StyleMap) *HTMLElement {
	e.style = style
	return e
}

// Element creates a custom element,
// should only be used for custom, nonstandard elements
func Element(tagName string) *HTMLElement {
	return &HTMLElement{tagName: tagName, a: make(map[string]interface{})}
}

// Element creates a custom element with a given namespace,
// should only be used for custom, nonstandard elements
func ElementNS(tagName string, namespace string) *HTMLElement {
	e := Element(tagName)
	e.namespace = namespace
	return e
}
