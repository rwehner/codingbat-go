package warmup1

import "testing"

func TestSleepIn(t *testing.T) {
	if !SleepIn(false, false) {
		t.Error(`sleepIn(false, false) = false`)
	}
	if !SleepIn(false, true) {
		t.Error(`sleepIn(false, true) = false`)
	}
	if SleepIn(true, false) {
		t.Error(`sleepIn(true, false) = true`)
	}
}
