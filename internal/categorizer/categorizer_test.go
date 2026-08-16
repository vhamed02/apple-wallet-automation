package categorizer_test

import (
	"testing"

	"github.com/apple-wallet-automation/internal/categorizer"
)

func testCategories() map[string][]string {
	return map[string][]string{
		"Groceries":     {"yerevan city", "carrefour", "supermarket", "grocery"},
		"Restaurant":    {"kfc", "mcdonald", "starbucks", "cafe", "restaurant"},
		"Transport":     {"uber", "bolt", "taxi", "parking", "fuel"},
		"Shopping":      {"amazon", "zara", "nike", "mall"},
		"Health":        {"pharmacy", "hospital", "gym"},
		"Entertainment": {"cinema", "netflix", "game"},
		"Utilities":     {"electric", "internet", "telecom"},
		"Travel":        {"hotel", "airbnb", "booking"},
	}
}

func TestCategorize_ExactKeyword(t *testing.T) {
	c := categorizer.New(testCategories())

	cases := []struct {
		merchant string
		want     string
	}{
		{"KFC", "Restaurant"},
		{"kfc", "Restaurant"},
		{"McDonald's", "Restaurant"},
		{"Starbucks Coffee", "Restaurant"},
		{"Yerevan City Komitas", "Groceries"},
		{"Carrefour Market", "Groceries"},
		{"Uber Eats", "Transport"},
		{"Bolt Taxi", "Transport"},
		{"Amazon Prime", "Shopping"},
		{"Zara Online", "Shopping"},
		{"City Pharmacy", "Health"},
		{"Netflix", "Entertainment"},
		{"Hilton Hotel", "Travel"},
		{"Airbnb", "Travel"},
	}

	for _, tc := range cases {
		t.Run(tc.merchant, func(t *testing.T) {
			got := c.Categorize(tc.merchant)
			if got != tc.want {
				t.Errorf("Categorize(%q) = %q, want %q", tc.merchant, got, tc.want)
			}
		})
	}
}

func TestCategorize_CaseInsensitive(t *testing.T) {
	c := categorizer.New(testCategories())

	cases := []struct {
		merchant string
		want     string
	}{
		{"KFC DOWNTOWN", "Restaurant"},
		{"YEREVAN CITY", "Groceries"},
		{"UBER TECHNOLOGIES", "Transport"},
		{"AMAZON.COM", "Shopping"},
	}

	for _, tc := range cases {
		t.Run(tc.merchant, func(t *testing.T) {
			got := c.Categorize(tc.merchant)
			if got != tc.want {
				t.Errorf("Categorize(%q) = %q, want %q", tc.merchant, got, tc.want)
			}
		})
	}
}

func TestCategorize_FallsBackToOther(t *testing.T) {
	c := categorizer.New(testCategories())

	cases := []string{
		"Random Unknown Shop",
		"XYZ Corp",
		"12345",
		"",
		"   ",
	}

	for _, merchant := range cases {
		t.Run(merchant, func(t *testing.T) {
			got := c.Categorize(merchant)
			if got != "Other" {
				t.Errorf("Categorize(%q) = %q, want %q", merchant, got, "Other")
			}
		})
	}
}

func TestCategorize_PartialMatch(t *testing.T) {
	c := categorizer.New(testCategories())

	cases := []struct {
		merchant string
		want     string
	}{
		{"New KFC Branch Yerevan", "Restaurant"},
		{"Best Cafe In Town", "Restaurant"},
		{"Downtown Parking Lot", "Transport"},
		{"24h Pharmacy Express", "Health"},
	}

	for _, tc := range cases {
		t.Run(tc.merchant, func(t *testing.T) {
			got := c.Categorize(tc.merchant)
			if got != tc.want {
				t.Errorf("Categorize(%q) = %q, want %q", tc.merchant, got, tc.want)
			}
		})
	}
}

func TestCategorize_EmptyCategories(t *testing.T) {
	c := categorizer.New(map[string][]string{})

	got := c.Categorize("KFC")
	if got != "Other" {
		t.Errorf("expected Other with empty categories, got %q", got)
	}
}
