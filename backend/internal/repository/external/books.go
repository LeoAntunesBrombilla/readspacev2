package external

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
	"time"
)

type externalBookRepository struct{}

func NewExternalBookRepository() repository.ExternalBookRepository {
	return &externalBookRepository{}
}

func (repo *externalBookRepository) SearchBooks(ctx context.Context, queryParam string, pagination int) ([]entity.ExternalBook, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")

	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&key=%s", queryParam, apiKey)

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

	var books []entity.ExternalBook

	for _, item := range items {
		volumeInfo := item.(map[string]interface{})["volumeInfo"].(map[string]interface{})

		var book entity.ExternalBook
		if err := json.Unmarshal(jsonify(volumeInfo), &book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func jsonify(m map[string]interface{}) []byte {
	data, _ := json.Marshal(m)
	return data
}
