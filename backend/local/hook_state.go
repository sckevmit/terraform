package local

import (
	"sync"

	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/states/statemgr"
	"github.com/hashicorp/terraform/terraform"
)

// StateHook is a hook that continuously updates the state by calling
// WriteState on a state.State.
type StateHook struct {
	terraform.NilHook
	sync.Mutex

	StateMgr statemgr.Writer
}

var _ terraform.Hook = (*StateHook)(nil)

func (h *StateHook) PostStateUpdate(s *states.State) (terraform.HookAction, error) {
	h.Lock()
	defer h.Unlock()

	if h.StateMgr != nil {
		// Write the new state
		if err := h.StateMgr.WriteState(s); err != nil {
			return terraform.HookActionHalt, err
		}
	}

	// Continue forth
	return terraform.HookActionContinue, nil
}
