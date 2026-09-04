package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func staticRoot(t *testing.T) string {
	t.Helper()
	// Tests run with package dir as cwd (module root for this repo).
	root, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return filepath.Join(root, "static")
}

func TestStaticRobotsTxt(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(staticRoot(t), "robots.txt"))
	if err != nil {
		t.Fatalf("static/robots.txt missing: %v", err)
	}
	body := string(data)
	for _, want := range []string{"User-agent:", "Allow:", "https://guisolski.github.io"} {
		if !strings.Contains(body, want) {
			t.Errorf("robots.txt missing %q", want)
		}
	}
	if strings.Contains(strings.ToLower(body), "disallow: /") && !strings.Contains(body, "Allow: /") {
		t.Error("robots.txt should allow crawling of the public site")
	}
}

func TestStaticLLMsTxt(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(staticRoot(t), "llms.txt"))
	if err != nil {
		t.Fatalf("static/llms.txt missing: %v", err)
	}
	body := string(data)
	for _, want := range []string{
		"Guilherme Solski Alves",
		"Mercado Libre",
		"FINK AI",
		"https://guisolski.github.io/",
		"https://www.linkedin.com/in/guilherme-solski-alves/",
		"https://github.com/guisolski",
		"/assets/pdf/resume.pdf",
		"guilhermesolskialves@gmail.com",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("llms.txt missing %q", want)
		}
	}
}

func TestResumePDFAssetsExist(t *testing.T) {
	for _, rel := range []string{
		"assets/pdf/resume.pdf",
		"assets/pdf/Curriculum/portugues.pdf",
	} {
		path := filepath.Join(staticRoot(t), "..", rel)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("%s missing: %v", rel, err)
			continue
		}
		if info.Size() < 1000 {
			t.Errorf("%s looks too small (%d bytes)", rel, info.Size())
		}
	}
}
