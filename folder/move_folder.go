package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if name == "" {
		return []Folder{}, errors.New("empty name")
	}

	if dst == "" {
		return []Folder{}, errors.New("empty destination file")
	}

	allFolders := f.folders

	var destinationPath string = ""
	var destinationOrgID uuid.UUID

	for _, folder := range allFolders {
		if folder.Name == dst {
			destinationPath = folder.Paths
			destinationOrgID = folder.OrgId
			break
		}
	}

	if destinationPath == "" {
		return []Folder{}, errors.New("destination folder not found")
	}

	modifiedFolders := []Folder{}

	for _, folder := range allFolders {
		if folder.Name == name {
			if strings.Contains(destinationPath, folder.Paths) {
				return []Folder{}, errors.New("cannot move a folder to child folder")
			}

			if folder.OrgId != destinationOrgID {
				return []Folder{}, errors.New("cannot move a folder to different orgID")
			}
		}

		// This section modifies the path of the target folder
		if folder.Name == name {
			folder.Paths = destinationPath + "." + folder.Name

			// This section modifies the path of the child folders of the target folder
		} else if strings.Contains(folder.Paths, name) {
			indexOfParentFolderName := strings.Index(folder.Paths, name)                  // Find the index of the parent folder in the child folders path
			folder.Paths = destinationPath + "." + folder.Paths[indexOfParentFolderName:] // Create new path with destination path as the beginning
		}

		modifiedFolders = append(modifiedFolders, folder)
	}

	return modifiedFolders, nil
}
