package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	if name == "" {
		return []Folder{}, errors.New("empty name")
	}

	if orgID == uuid.Nil {
		return []Folder{}, errors.New("invalid orgID")
	}

	allFolders := f.folders

	childFolders := []Folder{}

	parentFolderExists := false

	for _, folder := range allFolders {
		if folder.Name == name {
			parentFolderExists = true
			if folder.OrgId != orgID {
				return []Folder{}, errors.New("parent folder does not exist in specified organization")
			}
		}

		if strings.Contains(folder.Paths, name) && folder.Name != name {
			childFolders = append(childFolders, folder)
		}
	}

	if !parentFolderExists {
		return []Folder{}, errors.New("parent folder does not exist")
	}

	return childFolders, nil
}
