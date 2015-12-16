package pdclient

//NewClient - constructor for a new dispenser client
func NewClient(apiKey string, client clientDoer) *PDClient {
	return &PDClient{
		APIKey: apiKey,
		client: client,
	}
}