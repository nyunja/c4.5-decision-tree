package node

import (
	"testing"

	test "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

func TestGetMajorityClass(t *testing.T) {
	type args struct {
		instances     []test.Instance
		targetFeature string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single majority class",
			args: args{
				instances: []test.Instance{
					{"label": "A"},
					{"label": "A"},
					{"label": "B"},
				},
				targetFeature: "label",
			},
			want: "A",
		},
		{
			name: "Tie between two classes",
			args: args{
				instances: []test.Instance{
					{"label": "A"},
					{"label": "B"},
					{"label": "A"},
					{"label": "B"},
				},
				targetFeature: "label",
			},
			want: "A",
		},
		{
			name: "Different class distribution",
			args: args{
				instances: []test.Instance{
					{"label": "X"},
					{"label": "Y"},
					{"label": "X"},
					{"label": "X"},
					{"label": "Y"},
					{"label": "X"},
				},
				targetFeature: "label",
			},
			want: "X",
		},
		{
			name: "All instances have same class",
			args: args{
				instances: []test.Instance{
					{"label": "C"},
					{"label": "C"},
					{"label": "C"},
					{"label": "C"},
				},
				targetFeature: "label",
			},
			want: "C",
		},
		{
			name: "Empty instances list",
			args: args{
				instances:     []test.Instance{},
				targetFeature: "label",
			},
			want: "",
		},
		{
			name: "Missing target feature in some instances",
			args: args{
				instances: []test.Instance{
					{"label": "A"},
					{"other": "X"},
					{"label": "A"},
					{"label": "B"},
				},
				targetFeature: "label",
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMajorityClass(tt.args.instances, tt.args.targetFeature); got != tt.want {
				t.Errorf("GetMajorityClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMajorityClassFromNode(t *testing.T) {
	type args struct {
		node *test.Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single leaf node",
			args: args{
				node: &test.Node{
					Children: []*test.Node{
						{IsLeaf: true, Class: "A"},
					},
				},
			},
			want: "A",
		},
		{
			name: "Multiple leaf nodes, majority exists",
			args: args{
				node: &test.Node{
					Children: []*test.Node{
						{IsLeaf: true, Class: "A"},
						{IsLeaf: true, Class: "A"},
						{IsLeaf: true, Class: "B"},
					},
				},
			},
			want: "A",
		},
		{
			name: "No leaf nodes",
			args: args{
				node: &test.Node{
					Children: []*test.Node{
						{IsLeaf: false, Class: "A"},
						{IsLeaf: false, Class: "B"},
					},
				},
			},
			want: "", // No leaf nodes, should return empty string
		},
		{
			name: "Empty node",
			args: args{
				node: &test.Node{},
			},
			want: "", // No children, should return empty string
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMajorityClassFromNode(tt.args.node); got != tt.want {
				t.Errorf("GetMajorityClassFromNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
