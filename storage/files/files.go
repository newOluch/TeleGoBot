package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	e "github.com/p.kuznetsov/TeleGoBot/lib"
	"github.com/p.kuznetsov/TeleGoBot/storage"
)

type Storage struct {
	basePath string // хранит информацию о корневом расположении каталогов со статьями
}

const defaultPerm = 0774

var ErrNoSavedPages = errors.New("no saved pages")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s *Storage) Save(page *storage.Page) (err error) {

	defer func() { err = e.WrapIfErr("can't save page", err) }() // вначале функции определяем способ обработки ошибок.

	fPath := filepath.Join(s.basePath, page.UserName) // формируем путь до директории, куда будет сохраняться файл для каждого пользователя

	if err := os.MkdirAll(fPath, defaultPerm); err != nil { // создаём все нужные директории
		return err
	}

	fName, err := fileName(page) // формируем имя файла
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName) // дописываем имя файла к пути

	file, err := os.Create(fPath) // создаём файл

	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil { // записываем в него страницу в нужном формате. В результате этой процедуры
		// страница будет преобразована в формат gob и записана в указанный файл
		return nil
	}

	return nil
}

func (s *Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = e.Wrap("can't pick random page", err) }()

	path := filepath.Join(s.basePath, userName) // аналогично получаем путь до директории с файлами

	files, err := os.ReadDir(path) // получаем список файлов в директории

	if err != nil {
		return nil, err
	}

	if len(files) == 0 { // обработка ошибки в случае того, когда сохранённых страниц нет
		return nil, ErrNoSavedPages
	}

	rand.Seed(time.Now().UnixNano()) // генерируем псевдно-случайное число для получения статьи
	n := rand.Intn(len(files))       // указываем верхнюю границу, которая будет совпадать с числом файлов

	file := files[n] // получаем случайных файл, с тем номером, который мы сгенероировали

	return s.decodepage(filepath.Join(path, file.Name())) // декодируем файл и получаем его содержимое
}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)

	if err != nil {
		return e.Wrap("can't remove file", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can't remove file: %s", path)
		return e.Wrap(msg, err)
	}
	return nil
}

func (s Storage) IsExists(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)

	if err != nil {
		return false, e.Wrap("can't check if file exists", err)
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exists", path)
		return false, e.Wrap(msg, err)
	}
	return true, nil
}

func (s *Storage) decodepage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath) // открываем файл

	if err != nil { // обрабатываем ошибку
		return nil, e.Wrap("can't decode page", err)
	}
	defer func() { _ = f.Close() }() // закрываем файл

	var p storage.Page // создаём переменную, в которую файл будет декодирован

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("can't decode page", err)
	}
	return &p, nil
}

func fileName(p *storage.Page) (string, error) { // создаём функцию для определения имени файла
	return p.Hash()
}
