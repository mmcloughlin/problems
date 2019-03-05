package solution

func isMatch(s string, p string) bool {
	match, _ := matchstages(s, parse(p))
	return match
}

func matchstages(s string, stages []stage) (bool, bool) {
	if len(stages) == 0 {
		return len(s) == 0, true
	}
	stage, rest := stages[0], stages[1:]

	// consume min chars
	if len(s) < stage.min {
		return false, true
	}
	s = s[stage.min:]

	// non infinite case
	n := len(stage.sub)
	lim := len(s)
	if !stage.inf && n < lim {
		lim = n
	}

	subfound := false
	for i := 0; i+n <= lim; i++ {
		if s[i:i+n] != stage.sub {
			continue
		}
		subfound = true
		match, possible := matchstages(s[i+n:], rest)
		if match {
			return true, true
		}
		if !possible {
			return false, false
		}
	}

	return false, subfound
}

// stage of a pattern.
type stage struct {
	min int    // min chars to skip
	inf bool   // if we can skip an infinite number of chars
	sub string // substring to match
}

// parse a pattern into stages.
func parse(p string) []stage {
	stages := []stage{}
	for len(p) > 0 {
		var s stage

		// consume wildcards
		for ; len(p) > 0 && !isletter(p[0]); p = p[1:] {
			switch p[0] {
			case '?':
				s.min++
			case '*':
				s.inf = true
			}
		}

		// consume substring
		i := 0
		for ; i < len(p) && isletter(p[i]); i++ {
		}
		s.sub = p[:i]
		p = p[i:]

		stages = append(stages, s)
	}
	return stages
}

func isletter(b byte) bool {
	return 'a' <= b && b <= 'z'
}

/*
func isMatch(s string, p string) bool {
	// Check empty pattern.
	if len(p) == 0 {
		return len(s) == 0
	}

	// Consider next pattern character.
	op, p := p[0], p[1:]
	switch op {
	case '?':
		return len(s) > 0 && isMatch(s[1:], p)
	case '*':
		for ; len(p) > 0 && p[0] == '*'; p = p[1:] {
		}
		for i := 0; i <= len(s); i++ {
			if isMatch(s[i:], p) {
				return true
			}
		}
		return false
	default:
		return len(s) > 0 && s[0] == op && isMatch(s[1:], p)
	}
}
*/
