package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	type test struct {
		x    int
		y    int
		want int
	}

	tests := []test{
		{x: 1, y: 2, want: 3},
		{x: -1, y: -2, want: -3},
		{x: -1, y: 2, want: 1},
		{x: 1, y: -2, want: -1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d", tt.x, tt.y), func(t *testing.T) {
			got := Add(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
