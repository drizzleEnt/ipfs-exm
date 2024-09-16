package test

import (
	"fmt"
	"testing"

	"github.com/drizzleent/ipffs-exm/internal/model"
	"github.com/drizzleent/ipffs-exm/internal/service/file"
	"github.com/stretchr/testify/require"
)

type WrongUser struct {
}

func TestCatJson(t *testing.T) {
	t.Parallel()

	var (
	//serviceErr = fmt.Errorf("failed cat QmYMF2kVjafVc5NYqM cat: invalid path \"QmYMF2kVjafVc5NYqM\": path does not have enough components")
	)

	tests := []struct {
		name string
		cid  string
		want model.User
		err  error
	}{
		{
			name: "Seccses case",
			cid:  "QmYMF2kVjafVc5NYqMhQpJn7MXTNraAEPdtf6cS3k4ZoUQ",
			want: model.User{
				ID:    1,
				Name:  "Kiwi",
				Price: 5000,
			},
			err: nil,
		},
		{
			name: "Error case bad cid",
			cid:  "QmYMF2kVjafVc5NYqM",
			want: model.User{
				ID:    0,
				Name:  "",
				Price: 0,
			},
			err: fmt.Errorf("failed cat QmYMF2kVjafVc5NYqM cat: invalid path \"QmYMF2kVjafVc5NYqM\": path does not have enough components"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			srv := file.NewFileService()
			var res model.User
			err := srv.CatJSON(tt.cid, &res)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestCatFile(t *testing.T) {
	t.Parallel()

	var (
	//serviceErr = fmt.Errorf("failed cat QmYMF2kVjafVc5NYqM cat: invalid path \"QmYMF2kVjafVc5NYqM\": path does not have enough components")
	)

	tests := []struct {
		name string
		cid  string
		want string
		err  error
	}{
		{
			name: "Seccses case",
			cid:  "QmYMF2kVjafVc5NYqMhQpJn7MXTNraAEPdtf6cS3k4ZoUQ",
			want: `{"id":1,"name":"Kiwi","price":5000}`,
			err:  nil,
		},
		{
			name: "Error case bad cid",
			cid:  "QmYMF2kVjafVc5NYqM",
			want: ``,
			err:  fmt.Errorf("failed cat QmYMF2kVjafVc5NYqM cat: invalid path \"QmYMF2kVjafVc5NYqM\": path does not have enough components"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			srv := file.NewFileService()
			cid, err := srv.CatFile(tt.cid)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, cid)
		})
	}
}
