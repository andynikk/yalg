package sprint4

import (
	"math"
	"strconv"
	"strings"
)

func Hash(s string, p, m int) int {
	hash := 0
	pPow := 1
	for i := 0; i < len(s); i++ {
		hash = (hash + (int(s[0])-int('a')+1)*pPow) % m
		pPow = (pPow * p) % m
	}
	return hash
}

func Task4(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	a, _ := strconv.Atoi(lines[0])
	m, _ := strconv.Atoi(lines[1])
	s := lines[2]
	h := 0

	if len(s) == 0 {
		return "0"
	}

	st := len(s) - 1
	//var hashA float64 = 0
	//hashA := 1
	for i := 0; i < len(s); i++ {

		sm := int(s[i]) % m
		am := math.Pow(float64(a), float64(st))

		hashA := float64(sm) * am

		//power[0] = 1; power[i] = (power[i-1]*a) % m
		h += int(math.Mod(hashA, float64(m)))

		//hashA = int(math.Mod(aa, float64(m)))
		//if hashA < 0 {
		//	hashA %= m
		//	hashA += m
		//}

		//h += int(aa)
	}
	h = h % m
	//h = int(math.Mod(hashA, float64(m)))
	//if h < 0 {
	//	h %= m
	//	h += m
	//}
	/////////////////////////////////////////////////////////////////

	//h = int(s[0])
	//for i := 1; i < len(s); i++ {
	//	h = (h*a + int(s[i])) % m
	//}

	/////////////////////////////////////////////////////////////////
	//txt = "abcdef"
	//h1 := int(txt[0])
	//for i := 1; i < len(txt); i++ {
	//	h1 = (h1*a + int(txt[i])) % R
	//}
	//aa := 97 * math.Pow(float64(a), 5)
	//hashA := math.Mod(aa, float64(R))
	//hg := ((h1-int(hashA))*a + int('g')) % R
	//if hg < 0 {
	//	hg %= R
	//	hg += R
	//}
	//
	//fmt.Printf("%d\n", int(hashA))
	//fmt.Printf("%d\n", hg)
	//
	//txt = "bcdefg"
	//h2 := int(txt[0])
	//for i := 1; i < len(txt); i++ {
	//	h2 = (h2*a + int(txt[i])) % R
	//}
	//
	//txt = "bcdef"
	//h3 := int(txt[0])
	//for i := 1; i < len(txt); i++ {
	//	h3 = (h3*a + int(txt[i])) % R
	//}
	//
	//txt = "kvmzcab"
	//h = int(txt[0])
	//for i := 1; i < len(txt); i++ {
	//	h = (h*a + int(txt[i])) % R
	//}

	return strconv.Itoa(h) + "\n"
}
