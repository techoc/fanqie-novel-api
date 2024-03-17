package service

type ServiceGroup struct {
	BookService
	ChapterService
}

var ServiceGroupApp = new(ServiceGroup)
