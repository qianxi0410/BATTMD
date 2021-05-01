package hot100

func groupAnagrams(strs []string) [][]string {
	hash := map[rune]uint64{
			'a':2,'b':3,'c':5,'d':7,'e':11,'f':13,'g':17,'h':19,'i':23,'j':29,'k':31,'l':37,'m':41,
			'n':43,'o':47,'p':53,'q':59,'r':61,'s':67,'t':71,'u':73,'v':79,'w':83,'x':89,'y':97,'z':101,
	}

	stringToInt := func(str string) uint64 {
		var res uint64 = 1
		for _, v := range str {
			res *= hash[v]
		}
		return res
	}

	m := make(map[uint64][]string)

	for _, v := range strs {
		i := stringToInt(v)
		m[i] = append(m[i], v)
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}