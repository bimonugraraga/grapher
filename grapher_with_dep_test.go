package grapher

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewGraphWithDeps(t *testing.T) {
	test := []struct {
		name string
		want *GraphWithDeps
	}{
		{
			name: "success",
			want: &GraphWithDeps{
				Graph: Graph{
					nodes: make(map[string]*Node),
					edges: make(map[string][]string),
				},
				dependencies: make(map[string][]string),
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraphWithDeps(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n NewGraphWithDeps() = %v, \n want = %v", got, tt.want)
			}
		})
	}
}

func TestAddEdgeWithDepth(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	test := []struct {
		name  string
		args  args
		mock  func() (interface{}, error)
		mock2 func() (interface{}, error)
		want  []string
	}{
		{
			name: "success",
			args: args{
				from: "calculation",
				to:   "number_node",
			},
			mock: func() (interface{}, error) {
				return 10 + 20, nil
			},
			mock2: func() (interface{}, error) {
				return 42, nil
			},
			want: []string{"number_node"},
		},
		{
			name: "error",
			args: args{
				from: "calculation",
				to:   "number_node",
			},
			mock: func() (interface{}, error) {
				return 10 + 20, nil
			},
			mock2: func() (interface{}, error) {
				return 42, errors.New("test-error")
			},
			want: []string{"number_node"},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraphWithDeps()
			graph.AddNode(tt.args.from, tt.mock)
			graph.AddNode(tt.args.to, tt.mock2)

			graph.AddEdge(tt.args.from, tt.args.to)
			if got := graph.edges[tt.args.from]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n AddEdge() = %v, \n want = %v", got, tt.want)
			}
		})
	}

}

func TestAddEdgeV2(t *testing.T) {
	type args struct {
		from []string
	}
	test := []struct {
		name    string
		args    args
		mock    func() (interface{}, error)
		mock2   func() (interface{}, error)
		want    interface{}
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				from: []string{"calculation", "number_node"},
			},
			mock: func() (interface{}, error) {
				return 10 + 20, nil
			},
			mock2: func() (interface{}, error) {
				return 42, nil
			},
			want:    "success",
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				from: []string{"calculation", "number_node"},
			},
			mock: func() (interface{}, error) {
				return 10 + 20, nil
			},
			mock2: func() (interface{}, error) {
				return 42, errors.New("test-error")
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraphWithDeps()
			graph.AddNode(tt.args.from[0], tt.mock)
			graph.AddNode(tt.args.from[1], tt.mock2)

			graph.AddNode("test-node", func() (interface{}, error) {
				return "success", nil
			})
			graph.AddEdgeV2(tt.args.from, "test-node")
			graph.ExecuteWithDependencies()
			got, err := graph.GetNodeResult("test-node")
			if (err != nil) != tt.wantErr {
				t.Errorf("\n GetNodeResult() = %v, \n want = %v", err, nil)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n GetNodeResult() = %v, \n want = %v", got, tt.want)
			}
		})
	}
}

func TestExecuteWithDependencies(t *testing.T) {

}
