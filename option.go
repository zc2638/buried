package buried

type Option func(*buried)

// 设置接收组件
func NewSystem(s System) Option {
	return func(i *buried) {
		i.w.system = append(i.w.system, s)
	}
}