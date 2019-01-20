package sms

var letters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations := []string{""}
	for _, digit := range []byte(digits) {
		chars := []byte(letters[digit])
		next := make([]string, 0, len(combinations)*len(chars))
		for _, combination := range combinations {
			for _, char := range chars {
				s := combination + string([]byte{char})
				next = append(next, s)
			}
		}
		combinations = next
		if len(combinations) == 0 {
			break
		}
	}
	return combinations
}
