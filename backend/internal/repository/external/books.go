package external

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/interfaces"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type externalBookRepository struct{}

func NewExternalBookRepository() interfaces.ExternalBookRepository {
	return &externalBookRepository{}
}

func (repo *externalBookRepository) SearchBooks(ctx context.Context, q string, pagination int) ([]entity.ExternalBookResponse, error) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")

	encodedQuery := url.QueryEscape(q)

	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&key=%s", encodedQuery, apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var jsonResponse map[string]interface{}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, err
	}

	items, ok := jsonResponse["items"].([]interface{})

	if !ok {
		return nil, fmt.Errorf("items not found in jsonResponse")
	}

	var books []entity.ExternalBookResponse

	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		volumeInfo, ok := itemMap["volumeInfo"].(map[string]interface{})
		if !ok {
			continue
		}

		var book entity.ExternalBookResponse
		if err = json.Unmarshal(jsonify(volumeInfo), &book); err != nil {
			return nil, err
		}
		if id, ok := itemMap["id"].(string); ok {
			book.GoogleBookId = id
		}
		books = append(books, book)
	}

	return books, nil
}

func jsonify(m map[string]interface{}) []byte {
	data, _ := json.Marshal(m)
	return data
}
