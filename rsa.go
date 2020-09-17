package go_rsa

import "math/big"

// 加密
func Enc(key *PrivKey, m *big.Int) *big.Int {
	c := new(big.Int)
	c.Exp(m, key.Pub.E, key.Pub.N)
	return c
}

// 解密
func Dec(key *PrivKey, c *big.Int) *big.Int {
	m := new(big.Int)
	m = m.Exp(c, key.D, key.N)
	return m
}
