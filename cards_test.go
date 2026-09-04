package main

import "testing"

func TestIsExternal(t *testing.T) {
	tests := []struct {
		name string
		href string
		want bool
	}{
		{"https link", "https://github.com/guisolski", true},
		{"http link", "http://example.com", true},
		{"mailto link", "mailto:guilhermesolskialves@gmail.com", false},
		{"tel link", "tel:+5541996286624", false},
		{"relative link", "/assets/pdf/faculdade.pdf", false},
		{"empty href", "", false},
		{"short string", "ht", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isExternal(tc.href); got != tc.want {
				t.Errorf("isExternal(%q) = %v, want %v", tc.href, got, tc.want)
			}
		})
	}
}

func TestIsAssetLink(t *testing.T) {
	tests := []struct {
		href string
		want bool
	}{
		{"/assets/pdf/cv.pdf", true},
		{"/assets/pdf/Curriculum/portugues.pdf", true},
		{"/assets/images/ifpr_certified.jpeg", true},
		{"https://github.com/guisolski", false},
		{"/relax", false},
		{"", false},
	}
	for _, tc := range tests {
		t.Run(tc.href, func(t *testing.T) {
			if got := isAssetLink(tc.href); got != tc.want {
				t.Errorf("isAssetLink(%q) = %v, want %v", tc.href, got, tc.want)
			}
		})
	}
}
