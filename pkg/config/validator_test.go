package config

import "testing"

func TestValidate(t *testing.T) {
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
						JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
					},
					Entries: []Entry{
						{
							ID:      "foo",
							FileExt: "md",
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
						BaseDirectory: "/tmp",
					},
					Entries: []Entry{
						{
							ID:      "foo",
							FileExt: "md",
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
						BaseDirectory:    "/tmp",
						JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
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
						BaseDirectory:    "/tmp",
						JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
					},
					Entries: []Entry{
						{
							FileExt: "md",
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
						BaseDirectory:    "/tmp",
						JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
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
						BaseDirectory:    "/tmp",
						JournalDirectory: "{{.Year}}/{{.Month}}/{{.Day}}",
					},
					Entries: []Entry{
						{
							ID:      "foo",
							FileExt: "md",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.cfg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
