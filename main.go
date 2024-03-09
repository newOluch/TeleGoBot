package main

import (
	"flag"
	"log"
)

func main() {

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

}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
