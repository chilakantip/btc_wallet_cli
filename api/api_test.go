package api

import (
	"testing"
)

func TestGetUTXO(t *testing.T) {

	dd, _ := GetUTXO("12rwuzBoErFB3kv9boDxs9CBGjXnoS3mH9")
	t.Fatal(dd.UnspentOutputs[0].TxHash)

}
