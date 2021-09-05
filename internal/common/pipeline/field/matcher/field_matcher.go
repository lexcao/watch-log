package matcher

type Matcher = func(key string, value interface{}) (interface{}, bool)

func OmitMatcher(field string) Matcher {
	return func(key string, value interface{}) (interface{}, bool) {
		if len(field) > 0 && field == key {
			return nil, false
		} else {
			return value, true
		}
	}
}

func PickMatcher(field string) Matcher {
	return func(key string, value interface{}) (interface{}, bool) {
		if len(field) > 0 && field != key {
			return nil, false
		} else {
			return value, true
		}
	}
}
