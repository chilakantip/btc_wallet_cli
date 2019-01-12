package transactions

import (
	"testing"

	"github.com/chilakantip/btc_wallet_cli/keys"
)

func TestSetupKeysFromSeed(t *testing.T) {
	dd := keys.GetKeyTemplate()
	makeTxMsg(dd, "12rwuzBoErFB3kv9boDxs9CBGjXnoS3mH9", 0.000125)

}
