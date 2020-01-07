package decorator

import (
	"fmt"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	type fields struct {
		Ingredient IngredientAdd
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "case0",
			fields: fields{
				Ingredient: &Meat{&Onion{&PizzaDecorator{}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.Ingredient

			got, _ := p.AddIngredient()
			fmt.Println(got)

		})
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	type fields struct {
		Ingredient IngredientAdd
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Meat{
				Ingredient: tt.fields.Ingredient,
			}
			got, err := m.AddIngredient()
			if (err != nil) != tt.wantErr {
				t.Errorf("Meat.AddIngredient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Meat.AddIngredient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	type fields struct {
		Ingredient IngredientAdd
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Onion{
				Ingredient: tt.fields.Ingredient,
			}
			got, err := m.AddIngredient()
			if (err != nil) != tt.wantErr {
				t.Errorf("Onion.AddIngredient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Onion.AddIngredient() = %v, want %v", got, tt.want)
			}
		})
	}
}
