package buried

type System interface {
	Publish(b []byte) error
}

type Result interface {}