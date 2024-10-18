package folder

import (
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	allFolders := f.folders

	childFolders := []Folder{}

	for _, folder := range allFolders {
		indexOfFirstFullStop := strings.Index(folder.Paths, ".")
		if folder.Paths[0:indexOfFirstFullStop] == name {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders
}
