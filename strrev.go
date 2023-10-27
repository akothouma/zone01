package piscine

func StrRev(s string) string {
  var s1 string
  for i := range s{
	s1 += string(s[len(s)-i-1])
  }
 return s
}
