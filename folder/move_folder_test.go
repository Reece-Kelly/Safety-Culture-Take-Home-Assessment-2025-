package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	sampleData := folder.GetSampleData()

	// TODO: your tests here
	t.Parallel()
	tests := [...]struct {
		name      string
		dst       string
		folders   []folder.Folder
		want      []folder.Folder
		wantError error
	}{
		// TODO: Add tests here

		// Test to check that method works with imported data
		{
			name:    "topical-micromax",
			dst:     "creative-scalphunter",
			folders: sampleData[0:4],
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
					Paths: "creative-scalphunter.topical-micromax",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "creative-scalphunter.topical-micromax.bursting-lionheart",
				},
			},
			wantError: nil,
		},

		// Test to check error handling for error "empty name"
		{
			name:      "",
			dst:       "topical-micromax",
			folders:   sampleData[0:4],
			want:      []folder.Folder{},
			wantError: errors.New("empty name"),
		},

		// Test to check error handling for error "empty destination file"
		{
			name:      "topical-micromax",
			dst:       "",
			folders:   sampleData[0:4],
			want:      []folder.Folder{},
			wantError: errors.New("empty destination file"),
		},

		// Test to check error handling for error "destination folder not found"
		{
			name:      "clear-arclight",
			dst:       "invalid-folder",
			folders:   sampleData[0:4],
			want:      []folder.Folder{},
			wantError: errors.New("destination folder not found"),
		},

		// Test to check error handling for error "cannot move a folder to child folder"
		{
			name:      "clear-arclight",
			dst:       "topical-micromax",
			folders:   sampleData[0:4],
			want:      []folder.Folder{},
			wantError: errors.New("cannot move a folder to child folder"),
		},

		// Test to check error handling for error "cannot move a folder to different orgID"
		{
			name:      "bursting-lionheart",
			dst:       "striking-black-panther",
			folders:   sampleData[0:5],
			want:      []folder.Folder{},
			wantError: errors.New("cannot move a folder to different orgID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, gotError := f.MoveFolder(tt.name, tt.dst)

			assert.Equal(t, tt.want, get)

			if tt.wantError != nil {
				assert.EqualError(t, gotError, tt.wantError.Error())
			} else {
				assert.NoError(t, gotError)
			}
		})
	}
}
