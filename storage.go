package main

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

const storagePath = "./database.json"

type StorageFormat struct {
	Applications Applications `json:"applications"`
}

type Application struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	StartCommand string `json:"startCommand"`
}

type Applications = map[string]Application

type ApplicationCreateData struct {
	Name         string `json:"name"`
	StartCommand string `json:"startCommand"`
	Location     string `json:"location"`
}

func SetupStorage(forceRecreate bool) error {
	recreateFile := false
	stat, err := os.Stat(storagePath)

	if err != nil || stat.Size() == 0 {
		recreateFile = true
	}

	if recreateFile || forceRecreate {
		file, err := os.Create(storagePath)

		if err != nil {
			return err
		}

		initData := StorageFormat{
			Applications: make(map[string]Application, 0),
		}

		data, err := json.Marshal(initData)

		if err != nil {
			return err
		}

		file.Write([]byte(data))

		return nil
	}

	// TODO: Check integrity of the file

	return nil
}

func GetStorageData() (StorageFormat, error) {
	bytes, err := os.ReadFile(storagePath)

	if err != nil {
		return StorageFormat{}, err
	}

	data := StorageFormat{}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return StorageFormat{}, err
	}

	return data, nil
}

func rewriteStorage(newStorageData StorageFormat) error {
	file, err := os.OpenFile(storagePath, os.O_CREATE, 0660)

	if err != nil {
		return err
	}

	defer file.Close()

	byteData, err := json.Marshal(newStorageData)

	if err != nil {
		return err
	}

	_, err = file.Write(byteData)

	return err
}

func CreateApplication(createData ApplicationCreateData) error {
	storageData, err := GetStorageData()

	if err != nil {
		return err
	}

	newId := uuid.New().String()

	newApp := Application{
		ID: newId,
		Location: createData.Location,
		Name: createData.Name,
		StartCommand: createData.StartCommand,
	}

	storageData.Applications[newId] = newApp

	err = rewriteStorage(storageData)

	if err != nil {
		return err
	}

	return nil
}