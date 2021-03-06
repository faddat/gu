package trees_test

import (
	"testing"

	"github.com/gu-io/gu/tests"
	"github.com/gu-io/gu/trees"
)

func TestParseSelector(t *testing.T) {
	sels := trees.Query.ParseSelector("div:nth-child(n*2).shower.ball(before: all)")
	if sels == nil {
		tests.Failed(t, "Should have returned lists of selectors")
	}
	tests.Passed(t, "Should have returned lists of selectors")

	elem := sels[0]
	if elem.Tag != "div" {
		tests.Failed(t, "Should have div as selector tag")
	}
	tests.Passed(t, "Should have div as selector tag")

	if elem.Psuedo != ":nth-child(n*2)" {
		tests.Failed(t, "Should have psuedo selecto as selector tag")
	}
	tests.Passed(t, "Should have psuedo selecto as selector tag")

	if elem.AttrName != "" {
		tests.Failed(t, "Should not have a attribute match required: %q", elem.AttrName)
	}
	tests.Passed(t, "Should not have a attribute match required")

	if len(elem.Classes) != 2 {
		tests.Failed(t, "Should have 2 classes as part of element selector")
	}
	tests.Passed(t, "Should have 2 classes as part of element selector")

	if elem.Classes[0] != "shower" {
		tests.Failed(t, "Should contain %q as part of its classes", elem.Classes[0])
	}
	tests.Passed(t, "Should contain %q as part of its class", elem.Classes[0])

	if elem.Classes[1] != "ball" {
		tests.Failed(t, "Should contain %q as part of its classes", elem.Classes[1])
	}
	tests.Passed(t, "Should contain %q as part of its class", elem.Classes[1])

}

func TestParseAttr(t *testing.T) {
	sels := trees.Query.ParseSelector("div[rel|='bull'](before: all)")
	if sels == nil {
		tests.Failed(t, "Should have returned lists of selectors")
	}
	tests.Passed(t, "Should have returned lists of selectors")

	sel := sels[0]

	if sel.AttrName != "rel" {
		tests.Failed(t, "Should have selector with attribute name 'rel'")
	}
	tests.Passed(t, "Should have selector with attribute name 'rel'")

	if sel.AttrOp != "|=" {
		tests.Failed(t, "Should have selector with attribute op '|='")
	}
	tests.Passed(t, "Should have selector with attribute op '|='")

	if sel.AttrValue != "bull" {
		tests.Failed(t, "Should have selector with attribute value 'bull'")
	}
	tests.Passed(t, "Should have selector with attribute value 'bull'")
}

func TestParseMultiSelectors(t *testing.T) {
	sels := trees.Query.ParseSelector("div.shower.ball a.embeded, h1#header.main-header.color, div[alt*='blocker'], div::nth-child(2n+1)")
	if sels == nil {
		tests.Failed(t, "Should have returned lists of selectors")
	}
	tests.Passed(t, "Should have returned lists of selectors")

	if len(sels) != 4 {
		tests.Failed(t, "Should have returned only 4 elements")
	}
	tests.Passed(t, "Should have returned only 4 elements")

	if len(sels[0].Children) != 1 {
		tests.Failed(t, "Should have atleast first element with only 1 elements")
	}
	tests.Passed(t, "Should have atleast first element with only 1 elements")
}

func TestQueries(t *testing.T) {
	tree := trees.ParseAsRoot("section.tree-house#house", `
    <div class="wrapper" aria="wrapper-div">
      <section id="header" class="section"></section>
      <section id="menu" class="section"></section>
      <section id="content" class="section"></section>
    </div>

    <div class="links">
      <a rel="delay" href="#delay">Delay</a>
    </div>
  `)

	class, err := trees.GetAttr(tree, "class")
	if err != nil {
		tests.Failed(t, "Should have root with provided class property")
	}
	tests.Passed(t, "Should have root with provided class property")

	if _, val := class.Render(); val != "tree-house" {
		tests.Failed(t, "Should have class value matching 'tree-house': %q", val)
	}
	tests.Passed(t, "Should have class value matching 'tree-house'")

	id, err := trees.GetAttr(tree, "id")
	if err != nil {
		tests.Failed(t, "Should have root with provided id property")
	}
	tests.Passed(t, "Should have root with provided id property")

	if _, val := id.Render(); val != "house" {
		tests.Failed(t, "Should have class value matching 'house': %q", val)
	}
	tests.Passed(t, "Should have class value matching 'house'")

	if div := trees.Query.Query(tree, "div.wrapper"); div == nil {
		tests.Failed(t, "Should have returned a div with provided class 'wrapper'")
	}
	tests.Passed(t, "Should have returned a div with provided class 'wrapper'")

	if item := trees.Query.Query(tree, "section#menu"); item == nil {
		tests.Failed(t, "Should have returned a section with provided id 'menu'")
	}
	tests.Passed(t, "Should have returned a section with provided id 'menu'")

	if div := trees.Query.Query(tree, "div[aria*=wrapper-div]"); div == nil {
		tests.Failed(t, "Should have returned a div with provided attr 'div[aria*=wrapper-div]'")
	}
	tests.Passed(t, "Should have returned a div with provided attr 'div[aria*=wrapper-div]'")

	if div := trees.Query.Query(tree, "div[aria=wrapper-div]"); div == nil {
		tests.Failed(t, "Should have returned a div with provided attr 'div[aria*=wrapper-div]'")
	}
	tests.Passed(t, "Should have returned a div with provided attr 'div[aria*=wrapper-div]'")

	if div := trees.Query.QueryAll(tree, "div[aria*=wrapper-div]"); len(div) != 1 {
		tests.Failed(t, "Should have returned a div with provided attr 'div[aria*=wrapper-div]': %d", len(div))
	}
	tests.Passed(t, "Should have returned a div with provided attr 'div[aria*=wrapper-div]'")

	items := trees.Query.QueryAll(tree, "section.section")
	if len(items) != 3 {
		tests.Failed(t, "Should have returned 3 elements for selector 'section.section': %d", len(items))
	}
	tests.Passed(t, "Should have returned 3 elements for selector 'section.section'")

}
