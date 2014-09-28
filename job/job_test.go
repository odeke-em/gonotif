package job

import "testing"

func TestJobInit(t *testing.T) {
    j := New(0)
    if j == nil {
        t.Errorf("Not expected null back")
    }

    retrId := j.GetId()

    if retrId != 0 {
        t.Errorf("Expected zero as the id, instead got: %d", retrId)
    }
}
