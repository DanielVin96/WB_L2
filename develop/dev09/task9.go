package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func DownloadHtml(link string) (int, error) {
	// Cоздание коллектора colly, c максимальной глубиной поиска в 1
	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.URLFilters(
			regexp.MustCompile(link),
		),
	)

	// Cоздание map, для хранения и отслеживания уже загруженных ссылок в виде ключа
	downloadedLinks := make(map[string]bool)

	// OnHTML handler
	// Аргумент "a[href]" указывает на атрибут href
	c.OnHTML("a[href]", func(e *colly.HTMLElement) { // В кач-ве агрумента OnHTML передаем Аргумент "a[href]"
		href := e.Attr("href")
		if !downloadedLinks[href] { // если ключа нет (значит мы не загружали сайт)
			downloadedLinks[href] = true // то добавляем в map
			c.Visit(href)                // и посещаем сайт
		}
	})

	// OnResponse handler
	c.OnResponse(func(r *colly.Response) {
		//Получение полного пути к сайту для сохранения
		linkPath := strings.TrimPrefix(r.Request.URL.String(), link)
		if linkPath == "" { // если путь пустой
			linkPath = "/" // то устанавливаем его как "/"
		}
		filePath := filepath.Join(link, linkPath) // объед-ние нач-ого Url и относительного пути

		// сохранение файла
		err := r.Save(filePath)
		if err != nil {
			fmt.Println("Ошибка сохранения файла:", err)
			return
		}
	})

	//Посещение ссылки перед загрузкой
	err := c.Visit(link)
	if err != nil {
		return 0, err
	}

	// Возвращаем кол-во скачанных страниц
	return len(downloadedLinks), nil
}
func mkdir(folderName string) {
	_, err := os.Stat(folderName)
	if os.IsNotExist(err) && os.MkdirAll(folderName, os.ModePerm) != nil {
		log.Panic(err)
	}
}

func main() {
	mkdir("downloaded_pages") // Создание главной директории для скачанных стр-ц

	//если аргументов из командной строки меньше 2 то выводим ошибку
	if len(os.Args) < 2 {
		fmt.Println("Пожалуйста, дайте ссылку для скачивания")
		return
	}
	link := os.Args[1] // Сохранение url в переменную

	//запускаем метод скачивания сайта и возвращаем кол-во страниц
	numPages, err := DownloadHtml(link)
	if err != nil {
		fmt.Println("Ошибка скачивания сайта:", err)
		return
	}
	fmt.Printf("Загружено %d страниц\n", numPages) // печатаем кол-во скачанных страниц
}
