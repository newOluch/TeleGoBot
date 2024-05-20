package telegram

// здесь определены все типы, с которым будет работать наш клиент из telegram.go

/*
https://core.telegram.org/bots/api#making-requests
Ответ содержит объект JSON, который всегда имеет логическое поле «ok» и может иметь необязательное строковое поле «description»
с удобочитаемым описанием результата. Если «ok» равно True, запрос был успешным, и результат запроса можно найти в поле «result».
В случае неудачного запроса «ок» принимает значение «ложь», а ошибка объясняется в «описании».
Также возвращается целочисленное поле error_code, но его содержимое может быть изменено в будущем.
Некоторые ошибки также могут иметь необязательное поле «Параметры» типа ResponseParameters, которое может помочь автоматически обработать ошибку.
*/

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// https://core.telegram.org/bots/api#getting-updates
type Update struct {
	ID      int    `json:"update_id"` // update_id из Field
	Message string `json:"message"`   // message из Field
}
