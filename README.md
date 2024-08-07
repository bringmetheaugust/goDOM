# goDOM

<p align="center">
  <a href="https://tailwindcss.com" target="_blank">
    <picture>
      <img alt="goDOM logo" src="./logo.png" width="550" height="120" style="max-width: 100%;">
    </picture>
  </a>
</p>
<p align="center">
    <a href="https://pkg.go.dev/github.com/bringmetheaugust/goDOM"><img src="https://pkg.go.dev/badge/github.com/stretchr/testify" alt="Doc reference"></a>
    <a href="https://lh3.googleusercontent.com/proxy/w2a-pc4X9z2kuDWoXKnSF8pY6ngZvjVuZOAXMz3ZR8NwaUj9a-KsJnpcjtUSRO9QtFV6vMb3YoHWWv6k43Cb6bHOJEka19uE54GWtVx7Lru8gi10I_968eA2thkA0dL1O-zA8WT24cI"><img src="https://img.shields.io/badge/go%20version-1.21.5-61CFDD.svg?style=flat-square" alt="Golang version"></a>
    <a href="https://cs4.pikabu.ru/post_img/big/2014/12/15/4/1418619408_1209550583.jpg"><img src="https://img.shields.io/badge/version-1.2.3-blue" alt="project version"></a>
</p>

Made by front-ender for front-enders.   
Package provide method to parse HTML and get browser-like DOM and DOM API.  
It's only for reading DOM, searching elements and getting their data.
Doesn't have methods to mutate DOM.

## Installation

    go get github.com/bringmetheaugust/goDOM

## Examples

```go
package motherfckrs

import "github.com/bringmetheaugust/goDOM"

func main() {
    document, err := goDom.Create(bytes)       // create document (DOM with API, like in browser)
    if err != nil {return}                     // if markup is invalid
    lol, err := document.QuerySelector("#lol") // <a id="lol" class="pipi" href="http://lol.com">
    if err != nil {return}                     // if element not found
    fmt.Println(lol.ClassList)                 // ["pipi"]
    fmt.Println(lol.Attributes)                // {"id": "lol", class: "pipi", "href": "http://lol.com"}
    fmt.Println(lol.GetAttribute("href"))      // "http://lol.com"
}
```

## Docs

### DOM API methods

* `QuerySelector`
* `QuerySelectorAll`
* `GetElementById`
* `GetElementsByClassName`
* `GetElementsByTagName`

### DOM element's methods

* `GetAttribute`
* `HasAttribute`

### DOM element's fields

* `TagName`
* `TextContent`
* `Attributes`
* `Children`
* `ClassName`
* `ClassList`
* `Id`
* `ParentElement`

## Development

via Makefile

    make install

without Makefile

    sh ./scripts/install
