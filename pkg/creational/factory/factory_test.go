package factory

import (
	"reflect"
	"testing"
)

func TestGetPaymentMethod(t *testing.T) {
	type args struct {
		paymentType PayMentType
	}
	tests := []struct {
		name    string
		args    args
		want    PaymentMethod
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Card01",
			args:    args{paymentType: Card},
			want:    new(CardpaymentMethod),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPaymentMethod(tt.args.paymentType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPaymentMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPaymentMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
