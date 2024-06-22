package config

import "testing"

func TestValidateConfig(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base directory not set",
			args: args{
				cfg: Config{
					Paths: Paths{
						DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
					},
					Entries: []Entry{
						{
							ID:            "foo",
							FileExtension: "md",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "directory pattern not set",
			args: args{
				cfg: Config{
					Paths: Paths{
						BaseDir: "/tmp",
					},
					Entries: []Entry{
						{
							ID:            "foo",
							FileExtension: "md",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "no file types defined",
			args: args{
				cfg: Config{
					Paths: Paths{
						BaseDir:    "/tmp",
						DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "file type name not set",
			args: args{
				cfg: Config{
					Paths: Paths{
						BaseDir:    "/tmp",
						DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
					},
					Entries: []Entry{
						{
							FileExtension: "md",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "file extension not set",
			args: args{
				cfg: Config{
					Paths: Paths{
						BaseDir:    "/tmp",
						DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
					},
					Entries: []Entry{
						{
							ID: "foo",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "successful validation",
			args: args{
				cfg: Config{
					Paths: Paths{
						BaseDir:    "/tmp",
						DirPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
					},
					Entries: []Entry{
						{
							ID:            "foo",
							FileExtension: "md",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateConfig(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
