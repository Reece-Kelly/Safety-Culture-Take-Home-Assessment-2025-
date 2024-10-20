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

	// Check if name is empty
	if name == "" {
		return []Folder{}, errors.New("empty name") // Not too sure if this line is correct
	}

	if orgID == uuid.Nil {
		return []Folder{}, errors.New("invalid orgID")
	}

	allFolders := f.folders

	// Add error checking here to see if:
	// - There is a folder with that name within the list of all folders
	// - There is a folder with that name within the specified organisation
	// - There are folders within the folders list (e.g. making sure it is not empty)

	// Other TODO:
	// Need to implement the error checking into the test function?

	childFolders := []Folder{}

	parentFolderExists := false

	for _, folder := range allFolders {
		if folder.Name == name {
			parentFolderExists = true
		}

		if strings.Contains(folder.Paths, name) && folder.Name != name {
			childFolders = append(childFolders, folder)
		}
	}

	if parentFolderExists == false {
		return []Folder{}, errors.New("parent folder does not exist")
	}

	return childFolders, nil
}
