package app

type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	From From   `json:"from"`
	Text string `json:"text"`
}

type Chat struct {
	ID int `json:"id"`
}

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type From struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
}
