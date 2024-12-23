package model_test

import (
	"testing"

	"github.com/younker/thoughtful_ai/model"
)

func TestPackage_IsValid(t *testing.T) {
	tests := []struct {
		name       string
		mass       model.Mass
		dimensions model.Dimensions
		want       bool
	}{
		{
			name: "invalid due to missing values",
			want: false,
		},
		{
			name: "invalid due to missing dimension (depth)",
			mass: model.Mass(42),
			dimensions: model.Dimensions{
				Height: 23,
				Width:  32,
			},
			want: false,
		},
		{
			name: "invalid due to missing mass",
			dimensions: model.Dimensions{
				Height: 11,
				Width:  22,
				Depth:  33,
			},
			want: false,
		},
		{
			name: "all values present",
			mass: model.Mass(42),
			dimensions: model.Dimensions{
				Height: 11,
				Width:  22,
				Depth:  33,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model.Package{
				Mass:       tt.mass,
				Dimensions: tt.dimensions,
			}
			if got := m.IsValid(); got != tt.want {
				t.Errorf("Package.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackage_Volume(t *testing.T) {
	type fields struct {
		Height int
		Width  int
		Depth  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "zero",
			fields: fields{},
			want:   0,
		},
		{
			name: "one",
			fields: fields{
				Height: 1,
				Width:  1,
				Depth:  1,
			},
			want: 1,
		},
		{
			name: "two",
			fields: fields{
				Height: 2,
				Width:  2,
				Depth:  2,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model.Package{
				Dimensions: model.Dimensions{
					Height: tt.fields.Height,
					Width:  tt.fields.Width,
					Depth:  tt.fields.Depth,
				},
			}
			if got := m.Volume(); got != tt.want {
				t.Errorf("PackageDimensions.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackage_IsBulky(t *testing.T) {
	tests := []struct {
		name       string
		dimensions model.Dimensions
		want       bool
	}{
		{
			name: "has no size",
			want: false,
		},
		{
			name: "does not meet bulky range",
			dimensions: model.Dimensions{
				Height: model.MAX_VOLUME - 1,
				Width:  1,
				Depth:  1,
			},
			want: false,
		},
		{
			name: "matches max volume value",
			dimensions: model.Dimensions{
				Height: model.MAX_VOLUME,
				Width:  1,
				Depth:  1,
			},
			want: true,
		},
		{
			name: "exceeding bulky range by one",
			dimensions: model.Dimensions{
				Height: model.MAX_VOLUME + 1,
				Width:  1,
				Depth:  1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model.Package{
				Mass:       model.Mass(1),
				Dimensions: tt.dimensions,
			}
			if got := m.IsBulky(); got != tt.want {
				t.Errorf("Package.IsBulky() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackage_IsHeavy(t *testing.T) {
	tests := []struct {
		name string
		mass model.Mass
		want bool
	}{
		{
			name: "with no mass",
			mass: model.Mass(0),
			want: false,
		},
		{
			name: "does not meet mass upper bound",
			mass: model.Mass(model.MAX_MASS - 1),
			want: false,
		},
		{
			name: "matches mass upper bound",
			mass: model.Mass(model.MAX_MASS),
			want: true,
		},
		{
			name: "exceeds mass upper bound",
			mass: model.Mass(model.MAX_MASS + 1),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model.Package{
				Mass:       tt.mass,
				Dimensions: model.Dimensions{},
			}
			if got := m.IsHeavy(); got != tt.want {
				t.Errorf("Package.IsHeavy() = %v, want %v", got, tt.want)
			}
		})
	}
}
