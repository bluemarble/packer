package chroot

import (
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
)

// StepChrootProvision provisions the instance within a chroot.
type StepChrootProvision struct {
	mounts []string
}

func (s *StepChrootProvision) Run(state map[string]interface{}) multistep.StepAction {
	hook := state["hook"].(packer.Hook)
	mountPath := state["mount_path"].(string)
	ui := state["ui"].(packer.Ui)

	// Create our communicator
	comm := &Communicator{
		Chroot: mountPath,
	}

	// Provision
	log.Println("Running the provision hook")
	if err := hook.Run(packer.HookProvision, ui, comm, nil); err != nil {
		state["error"] = err
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepChrootProvision) Cleanup(state map[string]interface{}) {}
