package globals

type (
	ArgsCtx struct {
		CommandString *string
		Refresh       *int
		ProxyRefresh  *bool
		Bin           *string
	}
)

var ClientArgs *ArgsCtx = new(ArgsCtx)
