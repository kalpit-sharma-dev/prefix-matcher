package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		os.Setenv("File_Name", "filename")
		wantErr := false
		gotResp, err := LoadConfig()
		wantResp := "filename"
		if (err != nil) != wantErr {
			t.Errorf("error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(gotResp, wantResp) {
			t.Errorf("got = %v, want %v", gotResp, wantResp)
		}
	})

}

func TestSetupConfig(t *testing.T) {

	err := SetupConfig("sample_prefixes_(3).txt")
	wantErr := false
	if (err != nil) != wantErr {
		t.Errorf("error = %v, wantErr %v", err, wantErr)
		return
	}
}
