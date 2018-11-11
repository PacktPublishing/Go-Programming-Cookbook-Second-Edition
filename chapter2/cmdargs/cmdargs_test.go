package main

import "testing"

func TestMenuConf_SetupMenu(t *testing.T) {
	type fields struct {
		Goodbye bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"base-case", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MenuConf{
				Goodbye: tt.fields.Goodbye,
			}
			f := m.SetupMenu()
			f.Usage()
		})
	}
}

func TestMenuConf_GetSubMenu(t *testing.T) {
	type fields struct {
		Goodbye bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"base-case", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MenuConf{
				Goodbye: tt.fields.Goodbye,
			}
			f := m.GetSubMenu()
			f.Usage()

		})
	}
}

func TestMenuConf_Greet(t *testing.T) {
	type fields struct {
		Goodbye bool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"hello", fields{false}, args{"test"}},
		{"hello", fields{true}, args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MenuConf{
				Goodbye: tt.fields.Goodbye,
			}
			m.Greet(tt.args.name)
		})
	}
}

func TestMenuConf_Version(t *testing.T) {
	type fields struct {
		Goodbye bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"base-case", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MenuConf{
				Goodbye: tt.fields.Goodbye,
			}
			m.Version()
		})
	}
}
