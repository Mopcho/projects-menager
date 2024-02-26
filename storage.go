package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

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

func getDefaultDir() string {
	var defaultDir string
	if osname := os.Getenv("OS"); osname == "Windows_NT" {
		appdata := os.Getenv("APPDATA")
		defaultDir = filepath.Join(appdata, "MopManager")
	} else {
		home := os.Getenv("HOME")
		defaultDir = filepath.Join(home, ".mopmanager")
	}

	return defaultDir
}

func getDbPath() string {
	defaultDir := getDefaultDir()
	dbFile := filepath.Join(defaultDir, "database.json")
	return dbFile
}

func SetupStorage(forceRecreate bool) error {
	defaultDir := getDefaultDir()
	dbPath := getDbPath()

	// Ensure the directory exists
	err := os.MkdirAll(defaultDir, 0660)

	if err != nil {
		return err
	}
	
	recreateFile := false
	stat, err := os.Stat(dbPath)

	if err != nil || stat.Size() == 0 {
		recreateFile = true
	}

	if recreateFile || forceRecreate {
		file, err := os.Create(dbPath)

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
	bytes, err := os.ReadFile(getDbPath())

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
	file, err := os.OpenFile(getDbPath(), os.O_RDWR|os.O_TRUNC, 0660)

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

func GetApplications() (Applications, error) {
	storageData, err := GetStorageData()

	if err != nil {
		return nil, err
	}

	return storageData.Applications, nil
}

func DeleteApplication(id string) error {
	storageData, err := GetStorageData()

	if err != nil {
		return err
	}

	if _, ok := storageData.Applications[id]; !ok {
		return errors.New("App does not exist")
	}

	delete(storageData.Applications, id)

	err = rewriteStorage(storageData)

	return err
}
