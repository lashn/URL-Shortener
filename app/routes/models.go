package routes

//request format
type request struct {
	URL string `json:"url"`
}

//map for saving the data
var url_map = make(map[string]string)


