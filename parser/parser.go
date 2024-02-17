package parser

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// парсит информации с веб-страницы и записывает в inst_data.csv файл
func ParseInfluencers() {
	// искомый url
    url := "https://hypeauditor.com/top-instagram-all-russia/"
	// GET запрос по искомому url
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
	//чтение html документа из поля response.Body
    doc, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	// возвращает спарщенную таблицу, которая представляет массив структур Influencer
    influencers := extractInfluencers(doc)
	// записывает результат в виде таблицу в файл inst_data.csv в пакете data
    writeCSV(influencers)
}

// извлекает информацию об инфлюенсерах из html документа.
func extractInfluencers(doc *goquery.Document) []Influencer {
    influencers := []Influencer{}
	// находит строки таблицы (элемент с классом row)
    doc.Find(".row").Each(func(i int, s *goquery.Selection) {
        // пропускает row с заголовками
        if i==0{
            return
        }
		// извлекает текст из каждой строки таблицы, удаляет лишние пробелы
        rank := strings.TrimSpace(s.Find(".row-cell.rank > span").Text())
        name := strings.TrimSpace(s.Find(".contributor__name-content").Text())
        nick := strings.TrimSpace(s.Find(".contributor__title").Text())
        subscribers := strings.TrimSpace(s.Find(".subscribers").Text())
        audience := strings.TrimSpace(s.Find(".audience").Text())
        authentic := strings.TrimSpace(s.Find(".authentic").Text())
        engagement := strings.TrimSpace(s.Find(".engagement").Text())
		
        influencer := Influencer{
            Rank:        rank,
            Name:        name,
            Nick:        nick,
            Subscribers: subscribers,
            Audience:    audience,
            Authentic:   authentic,
            Engagement:  engagement,
        }

        influencers = append(influencers, influencer)
    })

    return influencers
}

// записывает данные об инфлюенсерах в .csv файл
func writeCSV(influencers []Influencer) {
    file, err := os.Create("./data/inst_data.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	// устанавливает кодировку utf-8 для записи файла, так как в спрашенной таблице есть символы кириллицы
    file.WriteString("\xEF\xBB\xBF")

    writer := csv.NewWriter(file)
    writer.Comma = ';'
    defer writer.Flush()
	//делаем построчную запись
    for _, influencer := range influencers {
        if err := writer.Write([]string{
            influencer.Rank, influencer.Name, influencer.Nick,
            influencer.Subscribers, influencer.Audience,
            influencer.Authentic, influencer.Engagement,
        }); err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Данные успешно записаны в файл inst_data.csv")
}
