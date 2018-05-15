package kwic

// DataStorageManager : interface base para armazenar as linhas
type DataStorageManager interface {
	Init() error
	Line(int) string
	Length() int
}
