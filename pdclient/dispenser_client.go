package pdclient

func NewClient(apiKey string, client clientDoer) *PDClient {
	return &PDClient{
		APIKey: apiKey,
		client: client,
	}
}
