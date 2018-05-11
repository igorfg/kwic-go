package kwic

// DataStorageManager : interface base para armazenar as linhas
type DataStorageManager interface {
	Init()
	Line(int) string
	Length() int
}
