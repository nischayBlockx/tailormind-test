package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchStudentData(id string) (*Student, error) {
	url := fmt.Sprintf("http://localhost:3000/api/v1/students/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch student data, status: %d", resp.StatusCode)
	}
	var student Student
	err = json.NewDecoder(resp.Body).Decode(&student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}
