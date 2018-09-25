package todoapi

import (
	"fmt"
	"testing"
)

func TestHCreateEnv(t *testing.T) {
	env, err := CreateEnv()

	if err != nil {
		t.Errorf("CreateEnv %v", err)
		return
	}
	fmt.Printf("env = %v", env)

	if env.Bind == "" {
		t.Errorf("env.Bind is null")
		return
	}

	if env.MasterURL == "" {
		t.Errorf("env.MasterURL is null")
		return
	}

	if env.SlaveURL == "" {
		t.Errorf("env.Bind is null")
		return
	}
}
