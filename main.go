package main

import (
	"flag"
	"log"

	"github.com/p.kuznetsov/TeleGoBot/clients/telegram"
)

// token = flags.Get(token) флаг для получения конфиденциального токена

// tgClient = telegram.New(token)

// tgClient = telegram.New(token) клиент для общения с Telegram

/*
   fetcher и processor будут общаться с API Telegram
   fetcher - будет отправлять туда запросы для получения новых событий
   processor - после обработки будет отправлять нам новые сообщения
*/
// fetcher = fetcher.New(tgClient) сборщик
// processor = processor.New(tgClient) обработчик
// consumer.Start(fetcher, processor). Потребитель = сборщик + обработчик

const (
	tgBothost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBothost, mustToken())
}

func mustToken() string { // функции, которые вместо того, чтобы возвращать ошибку - аварийно завершают программу имеют приставку "must"
	// такая практика предполагается для функций, в которых предмет возвращаемой ошибки имеет критическую роль в её работе
	// и без этого параметра функция бессмыслена. Основые области применения: запуск программы / парсинг конфигов
	token := flag.String("token-bot-token", "", "token for access to telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified") // записывает сообщение в лог через стандартный логгер и завершает программу с кодом выхода 1.
		// Когда вызывается log.Fatal, сообщение будет записано в стандартный вывод ошибок (stderr), после чего программа завершит свою работу.
		// Это используется для обозначения критических ошибок, после которых продолжение работы программы может быть невозможно или нежелательно.
	}

	return *token
}
