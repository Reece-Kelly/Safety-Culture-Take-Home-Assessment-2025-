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
	sampleData := folder.GetSampleData()

	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{

		// Test to check method works with matching OrgID
		{
			name:    "Matching OrgID",
			orgID:   uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
			folders: sampleData[0:3],
			want: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "creative-scalphunter",
				},
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
		},

		// Test to check method works with non-matching orgID
		{
			name:    "Non-Matching OrgID",
			orgID:   uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
			folders: sampleData[0:3],
			want:    []folder.Folder{},
		},

		// Test to check method works with folders with different OrgIDs in input slice
		{
			name:    "Multiple OrgIDs in input slice",
			orgID:   uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
			folders: append(sampleData[0:3], sampleData[79:82]...),
			want: []folder.Folder{
				{
					Name:  "noble-vixen",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "noble-vixen",
				},
				{
					Name:  "nearby-secret",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "noble-vixen.nearby-secret",
				},
				{
					Name:  "magnetic-sinister-six",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "noble-vixen.nearby-secret.magnetic-sinister-six",
				},
			},
		},

		// Test to check method works with no folders in slice
		{
			name:    "No folders",
			orgID:   uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
			folders: []folder.Folder{},
			want:    []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)

			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	sampleData := folder.GetSampleData()

	t.Parallel()
	tests := [...]struct {
		name      string
		orgID     uuid.UUID
		folders   []folder.Folder
		want      []folder.Folder
		wantError error
	}{

		//Test to check that method does not include parent folder in output
		{
			name:    "clear-arclight",
			orgID:   uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
			folders: sampleData[1:3],
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
			name:      "invalid folder",
			orgID:     uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
			folders:   sampleData[0:3],
			want:      []folder.Folder{},
			wantError: errors.New("parent folder does not exist"),
		},

		// Test to check that "parent folder does not exist in specified organization" error checking is working
		{
			name:      "clear-arclight",
			orgID:     uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
			folders:   sampleData[0:3],
			want:      []folder.Folder{},
			wantError: errors.New("parent folder does not exist in specified organization"),
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
