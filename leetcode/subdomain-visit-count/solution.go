package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	counts := map[string]int{}
	for _, cpdomain := range cpdomains {
		parts := strings.Fields(cpdomain)
		n, _ := strconv.Atoi(parts[0])
		domain := parts[1]

		counts[domain] += n
		b := []byte(domain)
		for i := 0; i < len(b); i++ {
			if b[i] == '.' {
				counts[string(b[i+1:])] += n
			}
		}
	}

	result := []string{}
	for domain, n := range counts {
		result = append(result, fmt.Sprintf("%d %s", n, domain))
	}

	return result
}
