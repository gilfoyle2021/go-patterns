package builder

import (
	"testing"
)

func TestSmallCar_Drive(t *testing.T) {

	type fields struct {
		Color Color
		Speed Speed
		Wheel Wheel
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &SmallCar{
				Color: tt.fields.Color,
				Speed: tt.fields.Speed,
				Wheel: tt.fields.Wheel,
			}
			if err := sc.Drive(); (err != nil) != tt.wantErr {
				t.Errorf("SmallCar.Drive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
