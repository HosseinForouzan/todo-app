package entity

import "testing"

func TestTaskStatus_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		status TaskStatus
		want   bool
	}{
		{
			name:   "todo is valid",
			status: StatusTodo,
			want:   true,
		},
		{
			name:   "in progress is valid",
			status: StatusInProgress,
			want:   true,
		},
		{
			name:   "done is valid",
			status: StatusDone,
			want:   true,
		},
		{
			name:   "invalid status",
			status: TaskStatus("cancelled"),
			want:   false,
		},
		{
			name:   "empty status",
			status: TaskStatus(""),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.status.IsValid()

			if got != tt.want {
				t.Errorf(
					"IsValid() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}