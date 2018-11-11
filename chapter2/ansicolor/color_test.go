package ansicolor

import "testing"

func TestColorText_String(t *testing.T) {
	type fields struct {
		TextColor Color
		Text      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"ColorNone", fields{ColorNone, "test"}, "test"},
		{"Red", fields{Red, "test"}, "\033[0;31mtest\033[0m"},
		{"Green", fields{Green, "test"}, "\033[0;32mtest\033[0m"},
		{"Yellow", fields{Yellow, "test"}, "\033[0;33mtest\033[0m"},
		{"Blue", fields{Blue, "test"}, "\033[0;34mtest\033[0m"},
		{"Magenta", fields{Magenta, "test"}, "\033[0;35mtest\033[0m"},
		{"Cyan", fields{Cyan, "test"}, "\033[0;36mtest\033[0m"},
		{"White", fields{White, "test"}, "\033[0;37mtest\033[0m"},
		{"Black", fields{Black, "test"}, "\033[0;30mtest\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ColorText{
				TextColor: tt.fields.TextColor,
				Text:      tt.fields.Text,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("ColorText.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
