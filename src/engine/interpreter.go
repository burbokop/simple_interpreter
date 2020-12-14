package engine

type Handler interface {
	Post(cmd Command)
}

type Command interface {
	Execute(handler Handler)
}

type EventLoop struct {
}

func (el *EventLoop) Start() {

}

func (el *EventLoop) Post(cmd Command) {

}
