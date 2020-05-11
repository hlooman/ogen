package primitives

import (
	"bytes"
	"testing"

	"github.com/go-test/deep"
	"github.com/olympus-protocol/ogen/utils/chainhash"
)

func TestWorkerCopy(t *testing.T) {
	worker := Worker{
		OutPoint: OutPoint{
			TxHash: chainhash.Hash{1},
			Index:  2,
		},
		PubKey:       [48]byte{5},
		Balance:      10,
		PayeeAddress: "test2",
		Status:       StatusActive,
	}
	worker2 := worker.Copy()

	worker.OutPoint.Index = 3
	if worker2.OutPoint.Index == 3 {
		t.Fatal("mutating outpoint mutates copy")
	}

	worker.PubKey[0] = 6
	if worker2.PubKey[0] == 6 {
		t.Fatal("mutating pubkey mutates copy")
	}

	worker.Balance = 7
	if worker2.Balance == 7 {
		t.Fatal("mutating balance mutates copy")
	}

	worker.PayeeAddress = "test3"
	if worker2.PayeeAddress == "test3" {
		t.Fatal("mutating payeeaddress mutates copy")
	}
}

func TestWorkerDeserializeSerialize(t *testing.T) {
	worker := Worker{
		OutPoint: OutPoint{
			TxHash: chainhash.Hash{1},
			Index:  2,
		},
		PubKey:       [48]byte{5},
		Balance:      10,
		PayeeAddress: "test2",
		Status:       StatusActive,
	}

	buf := bytes.NewBuffer([]byte{})
	err := worker.Serialize(buf)
	if err != nil {
		t.Fatal(err)
	}

	var worker2 Worker
	if err := worker2.Deserialize(buf); err != nil {
		t.Fatal(err)
	}

	if diff := deep.Equal(worker2, worker); diff != nil {
		t.Fatal(diff)
	}
}