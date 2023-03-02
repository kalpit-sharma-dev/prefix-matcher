package service

type IService interface {
	SearchWord(word string) (string, error)
}
