# goDOM

<p align="center">
    <picture>
      <img alt="goDOM logo" src="./assets/repo_logo.png" style="max-width: 100%; max-height: 100%">
    </picture>
</p>
<p align="center"><b>Easy, yeah?</b></p>
<p align="center">
    <a href="https://pkg.go.dev/github.com/bringmetheaugust/goDOM"><img src="https://pkg.go.dev/badge/github.com/stretchr/testify" alt="Doc reference"></a>
    <a href="https://lh3.googleusercontent.com/proxy/w2a-pc4X9z2kuDWoXKnSF8pY6ngZvjVuZOAXMz3ZR8NwaUj9a-KsJnpcjtUSRO9QtFV6vMb3YoHWWv6k43Cb6bHOJEka19uE54GWtVx7Lru8gi10I_968eA2thkA0dL1O-zA8WT24cI"><img src="https://img.shields.io/badge/go%20version-1.21.5-61CFDD.svg?style=flat-square" alt="Golang version"></a>
    <a href="https://cs4.pikabu.ru/post_img/big/2014/12/15/4/1418619408_1209550583.jpg"><img src="https://img.shields.io/badge/version-0.3.0-blue" alt="project version"></a>
</p>

Made by front-ender for front-enders.   
Package provide method to parse HTML and get browser-like [DOM](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model/Introduction#what_is_the_dom) and DOM API.    
It also has [jQuery](https://jquery.com/)-like API.

⚠️It's only for reading DOM, searching elements and getting their data.
Package doesn't have methods to mutate DOM.    
⚠️Before using it You should remember that sites can detect You as a bot and return unexpected HTML response.

## Installation

    go get github.com/bringmetheaugust/goDOM

## Examples

#### Using DOM-like API

```go
package motherfckrs

import "github.com/bringmetheaugust/goDOM"

func main() {
    // First, we'll get document
    bytes :=                                // HTML markup as bytes (from HTTP request, files, etc.)
    document, _, err := goDom.Create(bytes) // create document (DOM with DOM API, like in browser)
    if err != nil {return}                  // also check if markup is valid

    // Want to find some element by `id`?
    el, err := document.GetElementById("lol") // <a id="lol" class="pipi" href="http://lol.com">
    if err != nil {return}                    // check if element exists
    print(el.ClassList)                       // ["pipi"]
    print(el.Attributes)                      // {"id": "lol", class: "pipi", "href": "http://lol.com"}
    attr, _ := el.GetAttribute("href")        // "http://lol.com"

    // Or get a lot of elements by query selector?
    elements, err := document.QuerySelectorAll(".weee") // all elements in DOM which have class "weee"
    if err != nil {return}                              // check if elements are existed
    for _, el := range elements {                       // loop slice with existed elements
        // your best code here
    }
```

#### Using jQuery-like API

```go
package motherfckrs

import "github.com/bringmetheaugust/goDOM"

func main() {
    // First, we'll get document
    bytes :=                                // HTML markup as bytes (from HTTP request, files, etc.)
    _, jQ, err := goDom.Create(bytes)       // create jQ (jQuery with jQUery-like API)
    if err != nil {return}                  // also check if markup is valid

    // Want to find some element by `id` ?
    attr, err  := jQ("#lol").Attr("href")   // "http://lol.com" from <a id="lol" class="pipi" href="http://lol.com">

    // Or get `data-lol` attributes from elements with class `.wee` which have inside itself links with class `.piu`?
    jQ(".wee").Has("a.piu").Each(func(q) {
        a, _ := q.Attr("data-lol")
    })
```

#### More real examples [here](./examples).

## Docs

### Document

 * methods

    * [GetElementById](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById)
    * [GetElementsByClassName](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName)
    * [GetElementsByTagName](https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName)
    * [QuerySelector](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector) (doesn't support `>`, `+`, `~`, pseudo-elements, pseudo-classes)
    * [QuerySelectorAll](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll) (doesn't support `>`, `+`, `~`, pseudo-elements, pseudo-classes)

 * fields

	* [Body](https://developer.mozilla.org/en-US/docs/Web/API/Document/body)
    * [Doctype](https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype)
    * [Images](https://developer.mozilla.org/en-US/docs/Web/API/Document/images)
	* [Head](https://developer.mozilla.org/en-US/docs/Web/API/Document/head)
	* [Links](https://developer.mozilla.org/en-US/docs/Web/API/Document/links)
    * [Title](https://developer.mozilla.org/en-US/docs/Web/API/Document/title)

### Element

 * methods

    * [Contains](https://developer.mozilla.org/en-US/docs/Web/API/Node/contains)
    * [GetAttribute](https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute)
    * [GetElementById](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById)
    * [GetElementsByClassName](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName)
    * [GetElementsByTagName](https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName)
    * [HasAttribute](https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute)
    * [QuerySelector](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector) (doesn't support `>`, `+`, `~`, pseudo-elements, pseudo-classes)
    * [QuerySelectorAll](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll) (doesn't support `>`, `+`, `~`, pseudo-elements, pseudo-classes)

 * fields

    * [Attributes](https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes)
    * [Children](https://developer.mozilla.org/en-US/docs/Web/API/Element/children)
    * [ClassName](https://developer.mozilla.org/en-US/docs/Web/API/Element/className)
    * [ClassList](https://developer.mozilla.org/en-US/docs/Web/API/Element/classList)
    * [Id](https://developer.mozilla.org/en-US/docs/Web/API/Element/id)
    * [NextElementSibling](https://developer.mozilla.org/en-US/docs/Web/API/Element/nextElementSibling)
    * [ParentElement](https://developer.mozilla.org/en-US/docs/Web/API/Node/parentElement)
	* [PreviousElementSibling](https://developer.mozilla.org/en-US/docs/Web/API/Element/previousElementSibling)
    * [TagName](https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName)
    * [TextContent](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent) element text

### jQuery

 * [Attr](https://api.jquery.com/attr)
 * [Children](https://api.jquery.com/children)
 * [Each](https://api.jquery.com/each)
 * [Filter](https://api.jquery.com/filter)
 * [Find](https://api.jquery.com/find)
 * [First](https://api.jquery.com/first)
 * [Has](https://api.jquery.com/has)
 * [HasClass](https://api.jquery.com/hasClass)
 * [Last](https://api.jquery.com/last)
 * [Next](https://api.jquery.com/next)
 * [NextAll](https://api.jquery.com/nextAll)
 * [Not](https://api.jquery.com/not)
 * [Parent](https://api.jquery.com/parent)
 * [Parents](https://api.jquery.com/parents)
 * [Text] get current text element text content

## Something about jQuery API

This package uses jQuery API as the origin jQuery library (using JavaScript).    
For example in Golang with best practice, we should return data and/or errors almost from each function. How it should look like:

```go
    a, err := jQ(".li") // get elements with `li` classes
    if err != nil { return } // check if elements exist
    b, err := a.Has("a.my-link") // filter if they have links with class `my-link` inside itself
    if err != nil { return } // check if elements exist
    c, err := b.Find("div[data-lol=lala]") // find `div` elements with attribute `data-lol=lala` inside 
    if err != nil {return} // check if elements exist
    c.Each(func (q) { print(q) })
```

As You know, this package provides the original jQuery API, so every method (except `Attr`) always returns another jQuery element, even if elements are not found and a slice of elements is empty. So we can use jQuery as in origin jQuery library like this:

```go
    jQ(".li").Has("a.my-link").Find("div[data-lol=lala]").Each(func (q) { print(q) })
```

## Development

via Makefile

    make install

without Makefile

    sh ./scripts/install
