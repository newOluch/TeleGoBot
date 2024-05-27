package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	e "github.com/p.kuznetsov/TeleGoBot/lib"
)

// создаём клиент, который будет общаться с телеграмом.
// выполняет функции: получения апдейтов (новых сообщений) и отправка новых сообщений пользователям
type Client struct {
	host     string // host API сервиса Telegram
	basePath string // Префикс, с которого начинаются все запросы, следом, после указания хоста.
	//выглядит он, примерно, следующим образом: tg-bot.com/bot<token>. bot - это то с чего будет начинаться наш базовый путь.
	client http.Client // клиент который будет использоваться для выполнения HTTP-запросов
}

const (
	getUpdatesMethod  = "getUpdates"
	SendMessageMethod = "sendMessage"
)

// Когда создаётся новый экземпляр Client с помощью функции-конструктора New, он инициализируется с хостом API и токеном бота
func New(host string, token string) Client { // функция-конструктор, которая будет создавать Client
	return Client{
		host:     host,
		basePath: newBasePath(token), // Вызывает функцию newBasePath с параметром token для создания базового пути. Используется для аутентификации запросов к API
		client:   http.Client{},      // инициализируется новое значение http.Client
	}
}

func newBasePath(token string) string { // плюс данного подхода заключается в том, что если придётся создвавть токен в разных местах программы,
	// то мы сможем это делать с помощью одной и той же функции.
	// И если Тегерам решить изменить формарование префикса, то не придётся рефакторить код в десятках мест
	return "bot" + token
}

// Из https://core.telegram.org/bots/api#getupdates
func (c *Client) Updates(offset int, limit int) (updates []Update, err error) { // метод для получения обновлений от сервера Telegram
	defer func() { err = e.WrapIfErr("can't get updates", err) }()
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	// теперь нам нужно отправить запрос. Т.к. код для отправки запроса будет выглядеть одинакого для всех методов нашего клиента,
	// то мы вынесем его в отдельную функцию

	data, err := c.doRequest(getUpdatesMethod, q)

	var res UpdateResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Result, nil
}

// метод SendMessage используется для отправки сообщения в указанный чат
func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(SendMessageMethod, q)

	if err != nil {
		e.Wrap("can't send message", err)
	}

	return nil
}

// Метод, в котором формируется полный URL, включая host, basePath и метод
func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {

	defer func() { err = e.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method), // c.basePath + method
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil) // cоздаётся новый HTTP-запрос

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {

}
