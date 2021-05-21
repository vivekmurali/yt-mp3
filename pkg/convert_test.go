package pkg

import "testing"

func TestConvert(t *testing.T) {

	err := convert("../video.mp4")
	if err != nil {
		t.Errorf("%v", err)
	}
}
