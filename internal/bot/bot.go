package botapi

func NewBot() *Bot {
	return &Bot{}
}

func (b *Bot) GetUserID(username string) (int64, error) {
	return 11111, nil
}
