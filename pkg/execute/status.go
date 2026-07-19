package execute

type ExitStatus int

const (
	ExitStatusSuccess ExitStatus = iota
	ExitStatusInvalidCommandLineArgs
	ExitStatusCriticalException
	ExitStatusNotImplemented
)
