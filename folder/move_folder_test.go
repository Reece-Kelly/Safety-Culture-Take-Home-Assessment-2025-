package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	// TODO: your tests here
	t.Parallel()
	tests := [...]struct {
		name      string
		dst     string
		folders   []folder.Folder
		want      []folder.Folder
	}{
		// TODO: Add tests here
		{
			name: "nearby-secret",
			dst: "creatives-calphunter",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "creative-scalphunter",
				},
				{
					Name: "noble-vixen",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "noble-vixen",
				},
				{
					Name: "nearby-secret",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "noble-vixen.nearby-secret",
				},
				{
					Name: "magnetic-sinister-six",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "noble-vixen.nearby-secret.magnetic-sinister-six",
				},
			},
			want: []folder.Folder{
				{
					Name:  "creative scalphunter",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "creative-scalphunter",
				},
				{
					Name: "noble-vixen",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "noble-vixen",
				},
				{
					Name: "nearby-secret",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "creative-scalphunter.nearby-secret",
				},
				{
					Name: "magnetic-sinister-six",
					OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
					Paths: "creative-scalphunter.nearby-secret.magnetic-sinister-six",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.MoveFolder(tt.name, tt.dst)

			assert.Equal(t, tt.want, get)
		})
	}
}
