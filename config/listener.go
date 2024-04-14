package config

import "github.com/lyzzz123/illusion/log"

type ListenerTest struct {
}

func (listenerTest *ListenerTest) PreRun() error {
	log.Info("PreRun")
	return nil
}

func (listenerTest *ListenerTest) PostRun() error {
	log.Info("PostRun")
	return nil
}

func (listenerTest *ListenerTest) GetPriority() int {
	return 5
}
