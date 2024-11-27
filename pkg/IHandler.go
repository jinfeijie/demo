package pkg

type (
	Handler interface {
		Serve()
		Name() string
	}
)
