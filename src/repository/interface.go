package repository

type IRepository interface {
	//Insert(word chan string, wg *sync.WaitGroup, r *sync.RWMutex) error
	SearchWord(word string) int
}
