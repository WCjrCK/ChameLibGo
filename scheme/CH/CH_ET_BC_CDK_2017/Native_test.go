package CH_ET_BC_CDK_2017

import (
	"fmt"
	"testing"

	"github.com/Hugu1e/ChameLibGo/utils"
)

func run_scheme(t *testing.T, lamuda int64){
	pk, sk := KeyGen(lamuda)
	m1 := utils.GenerateBigNumber(lamuda)
	m2 := utils.GenerateBigNumber(lamuda)
	if m1.Cmp(m2) == 0 {
		t.Errorf("m1 == m2")
	}

	h1, r1, etd1 := Hash(pk, m1, lamuda)
	t.Log("h1.h: ", h1.H_1.H.String())
	t.Log("r1.r: ", r1.R_1.R.String())
	t.Log("etd1.sk: ", etd1.Sk_ch_2.D.String())
	if !Check(h1, r1, pk, m1) {
		t.Errorf("H(m1) invalid")
	}
	if Check(h1, r1, pk, m2) {
		t.Error()
	}

	h2, r2, _ := Hash(pk, m2, lamuda)
	if !Check(h2, r2, pk, m2) {
		t.Errorf("H(m2) invalid")
	}
	if Check(h2, r2, pk, m1) {
		t.Error()
	}

	r1_p := Adapt(h1, r1, etd1, pk, sk, m1, m2);
	t.Log("r1_p.r: ", r1_p.R_1.R.String())
	if !Check(h1, r1_p, pk, m2) {
		t.Errorf("adapt m2 invalid")
	}
	if Check(h1, r1_p, pk, m1) {
		t.Error()
	}
}

func Test_Native(t *testing.T) {
	cases := []struct {
		lamuda int64
	}{
		{32},
		{64},
		{128},
		{256},
		{512},
		{1024},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d lamuda %d", i+1, c.lamuda), func(t *testing.T) {
			run_scheme(t, c.lamuda)
		})
	}
}


