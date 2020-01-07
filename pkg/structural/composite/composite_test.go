package composite

import "testing"

func TestAthelete_Exec(t *testing.T) {
	type fields struct {
		Swim  Swimer
		Train Trainer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "case1",
			fields: fields{
				Swim:  &Swimimpl{},
				Train: &Trainimpl{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ath := &Athelete{
				Swimer:  tt.fields.Swim,
				Trainer: tt.fields.Train,
			}
			ath.Swim()
			ath.Train()
			ath.Exec()
		})
	}
}
