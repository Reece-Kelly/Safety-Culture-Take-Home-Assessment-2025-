package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	allFolders := f.folders

	modifiedFolders := []Folder{}

	if name == "" {
		return []Folder{}, errors.New("empty name")
	}

	var destinationPath string = ""

	for _, folder := range allFolders {
		if folder.Name == dst {
			destinationPath = folder.Paths
		}

	}

	for _, folder := range allFolders {
		if folder.Name == name && strings.Contains(destinationPath, folder.Paths) {
			return []Folder{}, errors.New("cannot move a folder to child folder")
		}
	}

	var newPath string

	for _, folder := range allFolders {
		// This section modifies the path of the target folder
		if folder.Name == name {
			newPath = destinationPath + "." + folder.Name // Not sure if this is the correct way to add strings together
			folder.Paths = newPath

			// This section modifies the path of the child folders of the parent folder
		} else if strings.Contains(folder.Paths, name) {
			indexOfParentFolderName := strings.Index(folder.Paths, name)                  // Find the index of the parent folder in the child folders path
			folder.Paths = destinationPath + "." + folder.Paths[indexOfParentFolderName:] // Create new path with destination path as the beginning
		}

		modifiedFolders = append(modifiedFolders, folder)
	}

	return modifiedFolders, nil
}
