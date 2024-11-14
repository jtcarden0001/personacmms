package category

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	tp "github.com/jtcarden0001/personacmms/cli/internal/types"
)

func CreateCategory(apiURL string, category *tp.Category) (*tp.Category, error) {
	body, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/categories", apiURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create category: %s", resp.Status)
	}

	var newCategory tp.Category
	if err := json.NewDecoder(resp.Body).Decode(&newCategory); err != nil {
		return nil, err
	}

	return &newCategory, nil
}

func DeleteCategory(apiURL string, title string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/categories/%s", apiURL, title), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete category: %s", resp.Status)
	}

	return nil
}

func GetCategory(apiURL string, title string) (*tp.Category, error) {
	resp, err := http.Get(fmt.Sprintf("%s/categories/%s", apiURL, title))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get category: %s", resp.Status)
	}

	var category tp.Category
	if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
		return nil, err
	}

	return &category, nil
}

func ListCategories(apiURL string) ([]tp.Category, error) {
	resp, err := http.Get(fmt.Sprintf("%s/categories", apiURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get categories: %s", resp.Status)
	}

	var categories []tp.Category
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func UpdateCategory(apiURL string, category *tp.Category) (*tp.Category, error) {
	body, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/categories/%s", apiURL, category.Title), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update category: %s", resp.Status)
	}

	var updatedCategory tp.Category
	if err := json.NewDecoder(resp.Body).Decode(&updatedCategory); err != nil {
		return nil, err
	}

	return &updatedCategory, nil
}
