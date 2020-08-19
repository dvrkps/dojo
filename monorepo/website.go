package monorepo

// Website holds data.
type Website struct {
	Key   string
	Index func(w *Website) (string, error)
	Parse func(w *Website) (string, error)
}
