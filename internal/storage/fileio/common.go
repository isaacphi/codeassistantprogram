package fileio

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func SaveFile(basePath, name string, contents string) error {
	filePath := filepath.Join(basePath, name)
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, []byte(contents), 0644)
}

func SaveYAML(data interface{}, basePath, subdir, id string) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	filePath := filepath.Join(basePath, subdir, id+".yaml")
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, yamlData, 0644)
}

func LoadYAML(data interface{}, basePath, subdir, id string) error {
	filePath := filepath.Join(basePath, subdir, id+".yaml")
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlData, data)
}

func DeleteFile(basePath, subdir, id string) error {
	filePath := filepath.Join(basePath, subdir, id+".yaml")
	return os.Remove(filePath)
}

func ListFiles(basePath, subdir string) ([]string, error) {
	dir := filepath.Join(basePath, subdir)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".yaml" {
			ids = append(ids, strings.TrimSuffix(file.Name(), ".yaml"))
		}
	}

	return ids, nil
}
