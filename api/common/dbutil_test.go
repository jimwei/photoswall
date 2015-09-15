package common

import (
	"testing"
)

func TestDb(t *testing.T) {
	err :=DB.DB().Ping()
	if err != nil{
		t.Error("ping fail.",err.Error())
	}else{
		t.Log("through.")
	}

	
}
