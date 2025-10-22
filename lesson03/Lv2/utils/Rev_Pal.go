package utils

func Reverse(s string) string {
	slen := len(s)
	buf := make([]byte, slen)
	for i := 0; i < slen; i++ {
		buf[i] = s[slen-i-1]
	}
	return string(buf)
}

func IsPalindrome(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	} else {
		return false
	}
}
