package bilibili

import (
	"net/http"
	"reflect"
	"testing"
)

func TestFindCIDSingleP(t *testing.T) {
	cids, err := FindCID(106)
	if err != nil {
		t.Fatal(err)
	}

	if cids[0] != 3635863 {
		t.Fatalf("CID for video with single p failed: %v", cids)
	}
}

func TestFindCIDMultiP(t *testing.T) {
	cids, err := FindCID(100)
	if err != nil {
		t.Fatal(err)
	}

	var shoudReturnSlice []int

	for i := 3631785; i <= 3631794; i++ {
		shoudReturnSlice = append(shoudReturnSlice, i)
	}

	if !reflect.DeepEqual(shoudReturnSlice, cids) {
		t.Fatalf("CID for video with series p failed: %v", cids)
	}
}

func TestConnectionError(t *testing.T) {
	err := ConnectionError{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad Request",
	}

	if s := err.Error(); s != "Bad Request. Status code: 400" {
		t.Fatal("ConnectionError failed to raise.")
	}
}
