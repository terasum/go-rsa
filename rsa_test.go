package go_rsa

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestDec(t *testing.T) {
	key, err := GenKey()
	assert.Nil(t, err)
	m := big.NewInt(123)
	c := Enc(key, m)
	t.Logf("c: %s", c.String())

	m2 := Dec(key, c)
	t.Logf("m: %s", m2.String())

}
