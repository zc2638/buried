package buried

type writer struct {
	system []System
}

func (w *writer) Write(p []byte) (n int, err error) {

	for _, s := range w.system {
		if err := s.Publish(p); err != nil {
			return 0, err
		}
	}
	return len(p), nil
}