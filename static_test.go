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
	for _, want := range []string{
		"User-agent:",
		"Allow:",
		"https://guisolski.github.io",
		// The sitemap is generated into dist/ at build time; robots.txt is
		// the only place that tells a crawler it exists.
		"Sitemap: https://guisolski.github.io/sitemap.xml",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("robots.txt missing %q", want)
		}
	}
	if strings.Contains(strings.ToLower(body), "disallow: /") && !strings.Contains(body, "Allow: /") {
		t.Error("robots.txt should allow crawling of the public site")
	}
}

// A JPEG served as .png is the kind of thing that renders fine in a browser
// (which sniffs) and is rejected by the scrapers that read the JSON-LD image.
func TestProfileImageMatchesItsExtension(t *testing.T) {
	path := filepath.Join(staticRoot(t), "..", strings.TrimPrefix(profileImage, "/"))
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("%s missing: %v", profileImage, err)
	}
	if len(data) < 2000 {
		t.Errorf("%s looks too small (%d bytes) to be a portrait", profileImage, len(data))
	}

	magic := map[string]string{
		".jpg":  "\xff\xd8\xff",
		".jpeg": "\xff\xd8\xff",
		".png":  "\x89PNG",
	}
	ext := filepath.Ext(profileImage)
	want, ok := magic[ext]
	if !ok {
		t.Fatalf("unhandled portrait extension %q", ext)
	}
	if !strings.HasPrefix(string(data), want) {
		t.Errorf("%s is not a %s: header %x", profileImage, ext, data[:4])
	}
}

func TestResumePDFAssetsExist(t *testing.T) {
	for _, rel := range []string{
		"assets/pdf/cv.pdf",
		"assets/pdf/resume.pdf", // kept identical so old bookmarks do not 404
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

	cv, err := os.ReadFile(filepath.Join(staticRoot(t), "..", "assets/pdf/cv.pdf"))
	if err != nil {
		t.Fatalf("cv.pdf: %v", err)
	}
	resume, err := os.ReadFile(filepath.Join(staticRoot(t), "..", "assets/pdf/resume.pdf"))
	if err != nil {
		t.Fatalf("resume.pdf: %v", err)
	}
	if string(cv) != string(resume) {
		t.Error("assets/pdf/resume.pdf must stay identical to assets/pdf/cv.pdf for old bookmarks")
	}
}
