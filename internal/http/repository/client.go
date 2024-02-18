package repository

import (
	"crypto-rates-client/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/patrickmn/go-cache"
)

func (r *Repository) GetAllCoins() ([]model.Coin, error) {
    // проверка кеша
    cacheKey := "all_prices"
    coinsFound, found := r.Cache.Get(cacheKey)
    if found {
        return coinsFound.([]model.Coin), nil
    }

    // кеш пуст или устарел -> делает запрос к API
    url := fmt.Sprintf("%s/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1", r.BaseURL)
    response, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("ошибка при выполнении запроса: %v", err)
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ошибка при выполнении запроса: статус код %d", response.StatusCode)
    }

    var coins []model.Coin
    if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
        return nil, fmt.Errorf("ошибка при декодировании JSON: %v", err)
    }

    // обновление кеша
    r.Cache.Set(cacheKey, coins, cache.DefaultExpiration)

    return coins, nil
}

func (r *Repository) GetCoinByCurrency(currency string) (model.Coin, error) {
    // проверка кеша
    cacheKey := fmt.Sprintf("price_%s", currency)
    coin, found := r.Cache.Get(cacheKey)
    if found {
        return coin.(model.Coin), nil
    }

    // кеш пуст или устарел -> делает запрос к API
    url := fmt.Sprintf("%s/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1", r.BaseURL)
    response, err := http.Get(url)
    if err != nil {
        return model.Coin{}, fmt.Errorf("ошибка при выполнении запроса: %v", err)
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return model.Coin{}, fmt.Errorf("ошибка при выполнении запроса: статус код %d", response.StatusCode)
    }

    var coins []model.Coin
    if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
        return model.Coin{}, fmt.Errorf("ошибка при декодировании JSON: %v", err)
    }

    for _, coin := range coins {
        if strings.EqualFold(strings.ToLower(coin.Name), strings.ToLower(currency)) {
            // обновление кеша
            r.Cache.Set(cacheKey, coin, cache.DefaultExpiration)
            return coin, nil
        }
    }

    return model.Coin{}, fmt.Errorf("криптовалюта %s не найдена", currency)
}

