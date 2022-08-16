package magic8

import (
	"math/rand"
	"testing"
	"time"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestMagic8(t *testing.T) {
	seed := time.Now().Unix()
	got := New(seed).Answer()
	rand.Seed(seed)
	want := list_answer[rand.Intn(len(list_answer))]

	// if got == failureMessage {
	// 	t.Errorf("GOT:%s", got)
	// }

	// index := rand.Intn(len(answer))
	if want != got {
		t.Errorf("GOT:%s is not the correct ramdom, because we recieve %s", got, want)
	}
	// if !contains(want,got) {
	//   t.Errorf("GOT:%s is not in list", got)
	// }

}
