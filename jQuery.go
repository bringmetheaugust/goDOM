package goDom

// Function which creates JQ. This is sepatare type cause this function can has its own methods.
type JQuery func(interface{}) JQ

func createJQuery(d *Document) JQuery {
	return func(i interface{}) JQ {
		switch v := i.(type) {
		case string:
			e, err := d.QuerySelectorAll(v)

			if err != nil {
				return JQ{}
			}

			return JQ{}.new(e...)
		case *Element:
			return JQ{}.new(v)
		default:
			return JQ{}
		}
	}
}
