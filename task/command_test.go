package task

import "testing"

func TestCommandTask_Execute(t *testing.T) {
	t.Parallel()

	cmd := NewCommandTask("go", "help", "build")
	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}

	cmd = NewCommandTask("unknown-command", "foo", "bar")
	if err := cmd.Execute(); err == nil {
		t.Error("expect to fail cmd but it succeeded")
	}
}
