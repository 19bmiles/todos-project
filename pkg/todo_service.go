package pkg

type ToDo interface {
	Create(text string, isDone bool)
}
