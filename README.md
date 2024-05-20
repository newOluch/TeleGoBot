# TeleGoBot

 Архитектура программы

token = flag.Get(token)

tgClient = telegram.New(token) --> клиент для работы с Телеграмом

fetcher = fetcher.New(tgClient) --> отправляет запросы к API Telegram, чтобы получить оттуда новые события

processor = processor.New(tgClient) --> обработчик событий

consumer.Start(fetcher, processor) --> получает и обрабатывает события