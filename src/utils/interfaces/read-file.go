package readfile

type ReadFile interface {
	GetValueByKey(key string) (string, error)
}
