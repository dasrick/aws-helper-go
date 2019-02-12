package awshelper

type Store interface {
	//GetAll(k string) (list []interface{}, err error)
	Get(k string, v interface{}) (found bool, err error)
	Store(k string, v interface{}) error
	//Delete(k string) error
	Close() error
}
