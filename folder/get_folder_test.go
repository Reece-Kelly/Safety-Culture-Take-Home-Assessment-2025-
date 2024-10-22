package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name      string
		orgID     uuid.UUID
		folders   []folder.Folder
		want      []folder.Folder
		wantError error
	}{
		// TODO: Add tests here

		//Test to check that function does not include parent folder in output
		{
			name:  "clear-arclight",
			orgID: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
			folders: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "creative-scalphunter.clear-arclight.topical-micromax",
				},
			},
			want: []folder.Folder{
				{
					Name:  "topical-micromax",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "creative-scalphunter.clear-arclight.topical-micromax",
				},
			},
			wantError: nil,
		},

		// Test to check if "empty name" error checking is working
		{
			name:      "",
			orgID:     uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
			folders:   []folder.Folder{},
			want:      []folder.Folder{},
			wantError: errors.New("empty name"),
		},

		// Test to check that "parent folder does not exist" error checking is working
		{
			name:  "invalid folder",
			orgID: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
			folders: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "creative-scalphunter.clear-arclight",
				},
			},
			want:      []folder.Folder{},
			wantError: errors.New("parent folder does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, gotError := f.GetAllChildFolders(tt.orgID, tt.name)

			assert.Equal(t, tt.want, get)

			if tt.wantError != nil {
				assert.EqualError(t, gotError, tt.wantError.Error())
			} else {
				assert.NoError(t, gotError)
			}
		})
	}
}
