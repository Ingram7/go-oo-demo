package app

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

type Mode struct {
	val string
}

func newMode(val string) *Mode {
	mode := new(Mode)
	mode.val = val
	return mode
}

func (m *Mode) IsDebug() bool {
	return m.val == DebugMode
}

func (m *Mode) IsRelease() bool {
	return m.val == ReleaseMode
}

func (m *Mode) IsTest() bool {
	return m.val == TestMode
}

func (m *Mode) String() string {
	return m.val
}
