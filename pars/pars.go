package pars

import (
	"sort"
	"strconv"
	"strings"
)

func Parse(s string) [][2]string {

	m := make(map[string]string)

	s = strings.Replace(s, "\n", "\t", -1)
	s = strings.Replace(s, "//", "", -1)
	sa := strings.Split(s, "\t")

	for i:=0; i<len(sa);i++  {
		if len(sa[i]) > 0 {

			sb := strings.Split(sa[i], " – ")
			if strings.Contains(sb[1], " (") {
				sbn := strings.Split(sb[1], " (")
				sb[1] = sbn[0]
			}
			m[sb[0]] = sb[1]
		}

	}

	keys := make([]int64, len(m))

	i := 0
	for k := range m {
		keys[i], _ = strconv.ParseInt(k,0,0)
		i++
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	v := make([][2]string, len(keys))
	for i:=0; i<len(keys);i++  {
		v[i][0] = strconv.FormatInt(keys[i], 10)
		v[i][1] = m[strconv.FormatInt(keys[i], 10)]
	}

	return v
}