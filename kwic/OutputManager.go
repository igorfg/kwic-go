package kwic

type OutputManager interface {
	Format([]string)
	Exhibit() error
}
