package validate

// must: if ture len(input) > 1, false mean that the input is a optional value
func Poster(uid, device string, must bool) error {
	return nil
}

// must: if ture len(input) > 1, false mean that the input is a optional value
func Pictures(in []string, must bool) error {
	// 1.是否是合法的URL数组
	// 2.URL格式是否合法
	// 3.URL是否合法
	// 4......
	return nil
}

// must: if ture len(input) > 1, false mean that the input is a optional value
func Content(in string, must bool) error {
	// 1.是否是合法的json格式
	// 2.内容是否合法
	// 3......
	return nil
}

// must: if ture len(input) > 1, false mean that the input is a optional value
func Location(in string, must bool) error {
	// 1.是否是合法的tag数组
	// 2.tag是否合法
	// 3......
	return nil
}

// must: if ture len(input) > 1, false mean that the input is a optional value
func Tags(in []string, must bool) error {
	// 1.是否是合法的tag数组
	// 2.tag是否合法
	// 3......
	return nil
}
