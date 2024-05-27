package telegram

import "github.com/p.kuznetsov/TeleGoBot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}
