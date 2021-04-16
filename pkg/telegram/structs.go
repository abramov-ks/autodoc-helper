package telegram

type TelegramConfig struct {
	Token  string `yaml:"token"`
	ChatId int    `yaml:"chat_id"`
}
