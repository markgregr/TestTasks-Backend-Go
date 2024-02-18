# Парсер данных Instagram Influencers

Этот проект предоставляет функциональность для парсинга данных о популярных Instagram Influencers с веб-страницы [hypeauditor.com](https://hypeauditor.com/top-instagram-all-russia/) и записи полученной информации в файл формата CSV.

## Установка

Для установки клиента выполните следующие шаги:

1. Склонируйте репозиторий на ваш компьютер:

   ```bash
   git clone https://github.com/markgregr/TestTasks-Backend-Go.git
   ```

2. Перейдите в каталог проекта:

   ```bash
   cd crypto-rates-client
   ```

3. Перейдите в ветку feature/crypto-rates-client:

   ```bash
   git checkout feature/instagram-parser
   ```

4. Установите зависимости:

   ```bash
   go mod tidy
   ```

## Использование

Чтобы использовать парсер данных Instagram Influencers, выполните следующие шаги:

1. Запустите приложение:

   ```bash
   go run main.goы
   ```

## Результат

После выполнения парсинга, данные об Instagram Influencers будут сохранены в файле inst_data.csv в каталоге data.
