package storage

import (
	"crypto/sha1"
	"fmt"
	"io"

	e "github.com/p.kuznetsov/TeleGoBot/lib"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) { // создаётся метод, для генерации хэша страницы, которым будет именоваться каждая новый файл в директории.
	//Это необходим  для соблюдения уникальности наименований в каталоге
	h := sha1.New() // создаём генератор

	if _, err := io.WriteString(h, p.URL); err != nil { // генерируем хэш
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil // возвращаем итоговый hash с помощью метода Sum()
}
