package go_rsa

import "testing"
import "github.com/stretchr/testify/assert"

func TestGenKey(t *testing.T) {
	pn, err := GenKey()
	assert.Nil(t, err)
	t.Logf("%v", pn)
}