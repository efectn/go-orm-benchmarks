package helper

var (
	OrmMaxIdle int
	OrmMaxConn int
	OrmSource  string
	DebugMode  bool
)

const (
	BatchSize      = 100
	InitialCounter = int64(1000)
	UpdatedCounter = int64(1001)
	GuardCounter   = int64(-1)
)
