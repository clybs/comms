package connections

import (
	"reflect"
	"testing"
)

func TestMapper_CreateMap(t *testing.T) {
	container := make(map[string]string)
	type args struct {
		container map[string]string
		key       string
		value     string
	}
	type results struct {
		length int
	}
	tests := []struct {
		name    string
		m       *Mapper
		args    args
		results results
	}{
		{
			"ab",
			&Mapper{},
			args{container, "a", "b"},
			results{1},
		},
		{
			"bc",
			&Mapper{},
			args{container, "b", "c"},
			results{2},
		},
		{
			"ab again",
			&Mapper{},
			args{container, "a", "b"},
			results{2},
		},
		{
			"ba",
			&Mapper{},
			args{container, "b", "a"},
			results{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.CreateMap(tt.args.container, tt.args.key, tt.args.value)
			if len(container) != tt.results.length {
				t.Errorf("CreateMap(%v, %v, %v) was incorrect, got: %v, wantLength: %v.", tt.args.container, tt.args.key, tt.args.value, len(container), tt.results.length)
			}
		})
	}
}

func TestMapper_CreateConnections(t *testing.T) {
	container := map[string]string{
		"0": "2",
		"1": "1",
		"2": "0, 3, 4",
		"3": "2, 4",
		"4": "2, 3, 6",
		"5": "6",
		"6": "4, 5",
		"7": "8",
		"8": "8",
		"9": "1, 7",
	}

	connections := map[string][]string{
		"0": {"0", "2", "3", "4", "5", "6"},
		"1": {"1"},
		"2": {"0", "2", "3", "4", "5", "6"},
		"3": {"0", "2", "3", "4", "5", "6"},
		"4": {"0", "2", "3", "4", "5", "6"},
		"5": {"0", "2", "3", "4", "5", "6"},
		"6": {"0", "2", "3", "4", "5", "6"},
		"7": {"8"},
		"8": {"8"},
		"9": {"1", "7", "8"},
	}

	type args struct {
		container map[string]string
	}
	type results struct {
		length      int
		connections map[string][]string
	}
	tests := []struct {
		name    string
		m       *Mapper
		args    args
		results results
	}{
		{
			"simple",
			&Mapper{},
			args{container},
			results{len(connections), connections},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapper{}
			connections := m.CreateConnections(tt.args.container)
			if len(connections) != tt.results.length {
				t.Errorf("CreateConnections(%v) was incorrect, got: %v, results: %v.", tt.args.container, len(connections), tt.results.length)
			}
			if !reflect.DeepEqual(connections, tt.results.connections) {
				t.Errorf("CreateConnections(%v) was incorrect, got: %v, results: %v.", tt.args.container, len(connections), tt.results.length)
			}
		})
	}
}

func TestMapper_CreateGroups(t *testing.T) {
	connections := map[string][]string{
		"0": {"0", "2", "3", "4", "5", "6"},
		"1": {"1"},
		"2": {"0", "2", "3", "4", "5", "6"},
		"3": {"0", "2", "3", "4", "5", "6"},
		"4": {"0", "2", "3", "4", "5", "6"},
		"5": {"0", "2", "3", "4", "5", "6"},
		"6": {"0", "2", "3", "4", "5", "6"},
		"7": {"8"},
		"8": {"8"},
		"9": {"1", "7", "8"},
	}

	type args struct {
		connections map[string][]string
	}
	tests := []struct {
		name string
		m    *Mapper
		args args
		want [][]string
	}{
		{
			"simple",
			&Mapper{},
			args{connections},
			[][]string{
				{"0", "2", "3", "4", "5", "6"},
				{"1"},
				{"8"},
				{"1", "7", "8"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapper{}
			if got := m.CreateGroups(tt.args.connections); len(got) != len(tt.want) {
				t.Errorf("Mapper.CreateGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
