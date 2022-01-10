package cli

type cliHandler interface {
	Launch() error
	Keyinit() error
	Exit() error
}
