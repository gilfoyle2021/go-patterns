package strategy

import "testing"

func TestConsolSquare_Print(t *testing.T) {
	tests := []struct {
		name    string
		c       *ConsolSquare
		wantErr bool
	}{
		{
			name: "case0",
			c:    &ConsolSquare{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConsolSquare{}
			if err := c.Print(); (err != nil) != tt.wantErr {
				t.Errorf("ConsolSquare.Print() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestImageSquare_Print(t *testing.T) {
	type fields struct {
		DestinationFilePath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "case0",
			fields: fields{
				DestinationFilePath: "xxx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ImageSquare{
				DestinationFilePath: tt.fields.DestinationFilePath,
			}
			i.Print()
			if err := i.Print(); (err != nil) != tt.wantErr {
				t.Errorf("ImageSquare.Print() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
