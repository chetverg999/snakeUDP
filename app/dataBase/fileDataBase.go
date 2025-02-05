package dataBase

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileDataBase struct {
	BestScore map[int]int `json:"best_score"`
}

func LoadFileDataBase() (*FileDataBase, error) {
	file, err := os.OpenFile("app/tmp/dataBase.json", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return &FileDataBase{make(map[int]int)}, nil
	}
	fmt.Println("Loading dataBase from file")
	var fileData FileDataBase
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&fileData)
	if err != nil {
		return &FileDataBase{BestScore: make(map[int]int)}, nil
	}
	return &fileData, nil
}

func (f *FileDataBase) SaveFileDataBase() error {
	file, err := os.Create("app/tmp/dataBase.json")
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(f)
}

func (f *FileDataBase) InsertScore(playerId int, score int) error {
	if currentHighScore, exists := f.BestScore[playerId]; !exists || score > currentHighScore {
		f.BestScore[playerId] = score
		return f.SaveFileDataBase()
	}
	return nil
}
