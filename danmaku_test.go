package bilibili

import "testing"

func TestDecryption(t *testing.T) {
	d := Danmaku{
		SenderID: "9ae0daaf",
	}

	if uid, err := d.DecryptUserID(); err != nil {
		t.Error(err)
	} else if uid != 12345678 {
		t.Error("decryption failed")
	}

}
