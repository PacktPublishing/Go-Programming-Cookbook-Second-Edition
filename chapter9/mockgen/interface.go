package mockgen

// GetSetter implements get a set of a
// key value pair
type GetSetter interface {
	Set(key, val string) error
	Get(key string) (string, error)
}
