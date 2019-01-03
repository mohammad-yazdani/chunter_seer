package listen

type RequestHandler func(string) (string, error)

type Request map[string][]string
type Response map[string]map[string]string

var handlers = make(map[string]RequestHandler, 0)

func AddHandler(tag string, handler RequestHandler)  {
	handlers[tag] = handler
}

func handleRequest(request Request) Response {
	output := make(map[string]map[string]string, 0)

	for key, val := range request {
		if output[key] == nil {
			output[key] = make(map[string]string, 0)
		}

		handler := handlers[key]
		for _, arg := range val {
			out, err := handler(arg)
			if err != nil {
				output[key][arg] = err.Error()
			} else {
				output[key][arg] = out
			}
		}
	}

	return output
}
