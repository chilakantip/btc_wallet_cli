package keys

import (
	"testing"
)

func TestSetupKeysFromSeed(t *testing.T) {
	pk, err := SetupKeysFromSeed("praveen")
	pk.ExportWIF()

}
