package repository

import (
	"reflect"
	"sync"
	"testing"
)

func TestSearchWord(t *testing.T) {
	db := Trie{}
	wantResp := 0
	repo := NewDatabaseRepository(&db)
	t.Run("success", func(t *testing.T) {
		gotResp := repo.SearchWord("gg")

		if !reflect.DeepEqual(gotResp, wantResp) {
			t.Errorf("got = %v, want %v", gotResp, wantResp)
		}
	})
}

func TestInsert(t *testing.T) {
	PrefixChan := make(chan string, 1000)
	wg := &sync.WaitGroup{}
	r := &sync.RWMutex{}
	wantErr := false
	PrefixChan <- "kalpit"
	PrefixChan <- "kalaait"
	close(PrefixChan)
	t.Run("success", func(t *testing.T) {
		wg.Add(1)
		err := Data.Insert(PrefixChan, wg, r)
		if (err != nil) != wantErr {
			t.Errorf("error = %v, wantErr %v", err, wantErr)
			return
		}
	})
}
