package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	test := []struct {
		name string
		want *Graph
	}{
		{
			name: "success",
			want: &Graph{
				nodes: make(map[string]*Node),
				edges: make(map[string][]string),
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n NewGraph() = %v, \n want = %v", got, tt.want)
			}
		})
	}
}

func TestAddNode(t *testing.T) {
	type args struct {
		id string
	}
	test := []struct {
		name    string
		args    args
		mock    func() (interface{}, error)
		want    interface{}
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: "calculation",
			},
			mock: func() (interface{}, error) {
				return 10 + 20, nil
			},
			want:    30,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				id: "calculation",
			},
			mock: func() (interface{}, error) {
				return nil, errors.New("test-error")
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph()
			graph.AddNode(tt.args.id, tt.mock)
			graph.ExecuteNode(tt.args.id)

			if got := graph.nodes[tt.args.id]; (got.Err != nil) != tt.wantErr {
				t.Errorf("\n AddNode() = %v, \n want = %v", got.Err, nil)
			}
			if got := graph.nodes[tt.args.id]; !reflect.DeepEqual(got.Result, tt.want) {
				t.Errorf("\n AddNode() = %v, \n want = %v", got, &Node{
					ID:       tt.args.id,
					Function: tt.mock,
					Executed: false,
				})
			}
		})
	}
}

func TestAddEdge(t *testing.T) {
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
			graph := NewGraph()
			graph.AddNode(tt.args.from, tt.mock)
			graph.AddNode(tt.args.to, tt.mock2)

			graph.AddEdge(tt.args.from, tt.args.to)
			if got := graph.edges[tt.args.from]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n AddEdge() = %v, \n want = %v", got, tt.want)
			}
		})
	}

}

func TestExcuteNode(t *testing.T) {
	type args struct {
		id string
	}
	test := []struct {
		name    string
		args    args
		mock    func() (interface{}, error)
		want    interface{}
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: "calculation",
			},
			mock: func() (interface{}, error) {
				return 10 + 20, nil
			},
			want:    30,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				id: "calculation",
			},
			mock: func() (interface{}, error) {
				return nil, errors.New("test-error")
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph()
			graph.AddNode(tt.args.id, tt.mock)
			graph.ExecuteNode(tt.args.id)

			if got := graph.nodes[tt.args.id]; (got.Err != nil) != tt.wantErr {
				t.Errorf("\n ExecuteNode() = %v, \n want = %v", got.Err, nil)
			}
			if got := graph.nodes[tt.args.id]; !reflect.DeepEqual(got.Result, tt.want) {
				t.Errorf("\n ExecuteNode() = %v, \n want = %v", got, &Node{
					ID:       tt.args.id,
					Function: tt.mock,
					Executed: true,
				})
			}
		})
	}
}

func TestExecuteGraph(t *testing.T) {
	type args struct {
		to   string
		from string
	}
	test := []struct {
		name    string
		args    args
		mock    func() (interface{}, error)
		mock2   func() (interface{}, error)
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				to:   "number_node",
				from: "calculation",
			},
			mock: func() (interface{}, error) {
				return 10 + 20, nil
			},
			mock2: func() (interface{}, error) {
				return 42, nil
			},
			want: map[string]interface{}{
				"calculation": 30,
				"number_node": 42,
			},
			wantErr: false,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph()
			graph.AddNode(tt.args.from, tt.mock)
			graph.AddNode(tt.args.to, tt.mock2)
			graph.AddEdge(tt.args.from, tt.args.to)

			results := graph.ExecuteGraph()
			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("\n ExecuteGraph() = %v, \n want = %v", results, tt.want)
			}
		})
	}
}

func TestGetNodeResult(t *testing.T) {
	type args struct {
		id      string
		wrongId string
	}
	test := []struct {
		name    string
		args    args
		mock    func() (interface{}, error)
		want    interface{}
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: "calculation",
			},
			mock: func() (interface{}, error) {
				return 10 + 20, nil
			},
			want:    30,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				id:      "calculation",
				wrongId: "wrong-calculation",
			},
			mock: func() (interface{}, error) {
				return nil, errors.New("test-error")
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error-node_not_found",
			args: args{
				id: "calculation",
			},
			mock: func() (interface{}, error) {
				return nil, errors.New("test-error")
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph()

			graph.AddNode(tt.args.id, tt.mock)
			graph.ExecuteNode(tt.args.id)

			targetNode := tt.args.id
			if tt.name == "error-node_not_found" {
				targetNode = tt.args.wrongId
			}
			got, err := graph.GetNodeResult(targetNode)
			if (err != nil) != tt.wantErr {
				t.Errorf("\n GetNodeResult() = %v, \n want = %v", err, nil)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n GetNodeResult() = %v, \n want = %v", got, tt.want)
			}
		})
	}
}
