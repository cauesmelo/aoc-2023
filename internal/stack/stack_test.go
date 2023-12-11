package stack

import "testing"

func TestPush(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "Push 1", want: []int{1},
		},
		{
			name: "Push 1, 2", want: []int{1, 2},
		},
		{
			name: "Push 10, 20, 2", want: []int{10, 20, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack[int]{}

			for _, n := range tt.want {
				s.Push(n)
			}

			curr := s.top
			arrPos := len(tt.want) - 1
			lenStack := 0

			for curr != nil {
				if tt.want[arrPos] != curr.value {
					t.Errorf("got %d; want %d", s.top.value, tt.want)
				}

				curr = curr.prev
				arrPos--
				lenStack++
			}

			if len(tt.want) != lenStack {
				t.Errorf("got len %d; want len %d", lenStack, len(tt.want))
			}
		})
	}
}

func TestPop(t *testing.T) {
	s := Stack[int]{}

	s.Push(50)
	s.Push(2)
	s.Push(1)

	tests := []struct {
		name string
		want int
	}{
		{
			name: "Pop 1", want: 1,
		},
		{
			name: "Pop 2", want: 2,
		},
		{
			name: "Pop 50", want: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := s.Pop()

			if tt.want != v {
				t.Errorf("got %d; want %d", v, tt.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Len 1", want: 1,
		},
		{
			name: "Len 4", want: 4,
		},
		{
			name: "Len 50", want: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := Stack[int]{}

			for i := 0; i < tt.want; i++ {
				s.Push(i)
			}

			stackLen := s.Len()

			if tt.want != stackLen {
				t.Errorf("got %d; want %d", stackLen, tt.want)
			}
		})
	}
}
