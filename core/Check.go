package core

type checkResult struct {
	key   string
	value words
}

func worker(content string, class string, response chan checkResult) {
	var resp checkResult
	resp = checkResult{
		key:   "违禁品",
		value: words{"abc", "gun"},
	}
	response <- resp
	wg.Done()
}
