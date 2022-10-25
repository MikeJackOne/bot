package Messenger

var apiUrl string

type FbMessage struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Message struct {
	ID          string
	Time        int64
	SenderId    string
	RecipientId string
	Text        string
}

type Entry struct {
	ID        string `json:"id"`
	Time      int64  `json:"time"`
	Messaging []struct {
		Sender struct {
			ID string `json:"id"`
		} `json:"sender"`
		Recipient struct {
			ID string `json:"id"`
		} `json:"recipient"`
		Timestamp int64 `json:"timestamp"`
		Message   struct {
			Mid  string `json:"mid"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messaging"`
}

type Bot struct {
	Token       string
	AccessToken string
	core        []func(Message) bool
	ApiUrl      string
}

type SmallMessage struct {
	Text string `json:"text"`
}
type Recipient struct {
	ID string `json:"id"`
}

type Reply struct {
	Message   SmallMessage `json:"message"`
	Recipient Recipient    `json:"recipient"`
}
