package gotree_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sh0e1/gotree"
)

func TestTree(t *testing.T) {
	tests := []struct {
		name    string
		w       *bytes.Buffer
		dirs    []string
		opt     *gotree.Option
		want    string
		wantErr error
	}{
		{
			name: "OptionIsNil",
			w:    &bytes.Buffer{},
			dirs: []string{"./testdata"},
			opt:  nil,
			want: `./testdata
├── file
└── sub1
   ├── file
   └── sub2
      └── file
`,
			wantErr: nil,
		},
		{
			name: "OptionLevelIsOne",
			w:    &bytes.Buffer{},
			dirs: []string{"./testdata"},
			opt:  &gotree.Option{Level: 1},
			want: `./testdata
├── file
└── sub1
`,
			wantErr: nil,
		},
		{
			name: "OptionAllDisplay",
			w:    &bytes.Buffer{},
			dirs: []string{"./testdata"},
			opt:  &gotree.Option{IsDisplayAllFiles: true, Level: -1},
			want: `./testdata
├── .invisible
├── file
└── sub1
   ├── file
   └── sub2
      └── file
`,
			wantErr: nil,
		},
		{
			name: "MultipleDirectories",
			w:    &bytes.Buffer{},
			dirs: []string{"./testdata", "./testdata/sub1/sub2"},
			opt:  &gotree.Option{},
			want: `./testdata
./testdata/sub1/sub2
`,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := gotree.Execute(tt.w, tt.dirs, tt.opt)
			if tt.wantErr != err {
				t.Errorf("want %v, got %v", tt.wantErr, err)
				return
			}
			if tt.wantErr != nil && err != nil && tt.wantErr.Error() != err.Error() {
				t.Errorf("want %v, got %v", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, tt.w.String()); diff != "" {
				t.Errorf("differs (-want +got):\n%s", diff)
			}
		})
	}
}
