package categorizer

import "strings"

type Categorizer struct {
	categories map[string][]string
}

func New(categories map[string][]string) *Categorizer {
	return &Categorizer{categories: categories}
}

func (c *Categorizer) Categorize(merchant string) string {
	normalized := strings.ToLower(strings.TrimSpace(merchant))

	for category, keywords := range c.categories {
		for _, kw := range keywords {
			if strings.Contains(normalized, strings.ToLower(kw)) {
				return category
			}
		}
	}

	return "Other"
}
