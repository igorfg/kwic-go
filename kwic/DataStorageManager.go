package kwic

type DataStorageManager interface {
	Init()
	Line(int) string
	Length() int
}
