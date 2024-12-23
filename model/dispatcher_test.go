package model_test

import (
	"testing"

	"github.com/younker/cribl/model"
)

type TestDispatchObj struct {
	isHeavy bool
	isBulky bool
}

func (o TestDispatchObj) IsHeavy() bool {
	return o.isHeavy
}

func (o TestDispatchObj) IsBulky() bool {
	return o.isBulky
}

func TestDispatch(t *testing.T) {
	tests := []struct {
		name string
		disp model.Dispatchable
		want model.Stack
	}{
		{
			name: "rejected package due to weight & size",
			disp: TestDispatchObj{isHeavy: true, isBulky: true},
			want: model.Rejected,
		},
		{
			name: "special package due to weight",
			disp: TestDispatchObj{isHeavy: true},
			want: model.Special,
		},
		{
			name: "special package due to size",
			disp: TestDispatchObj{isBulky: true},
			want: model.Special,
		},
		{
			name: "is neither heavy nor bulky",
			disp: TestDispatchObj{},
			want: model.Standard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := model.Dispatch(tt.disp); got != tt.want {
				t.Errorf("Dispatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
