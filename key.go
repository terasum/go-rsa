package go_rsa

import "math/big"
import "crypto/rand"

type Pubkey struct {
	E *big.Int
	N *big.Int
}

type PrivKey struct {
	Pub *Pubkey
	D *big.Int
	N *big.Int
}

func GenKey() (*PrivKey, error) {
	// 选取幂次 e
	e := big.NewInt(65537) // 65537

	// 选取素数 p,q
	p, err := rand.Prime(rand.Reader, 1024)
	if err != nil {
		return nil,err
	}

	q, err := rand.Prime(rand.Reader, 1024)
	if err != nil {
		return nil,err
	}

	n := new(big.Int)
	n.Mul(p, q)

	// 计算 φ(n) 欧拉函数
	PHIn := new(big.Int)
	PHIn.Mul(p.Sub(p,big.NewInt(1)), q.Sub(q, big.NewInt(1)))

	// 计算n 对 e 的模反元素
	d := new(big.Int)
	d.ModInverse(e, PHIn)

	// 构造私钥
	return &PrivKey{
		Pub: &Pubkey{N: n, E: e},
		D:   d,
		N:   n,
	}, nil

}
