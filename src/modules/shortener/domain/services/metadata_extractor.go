package services

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// LinkMetadata represents consolidated rich-preview metadata for a URL.
// Fields follow common Open Graph / Twitter Card naming conventions.
type LinkMetadata struct {
	URL         string            `json:"url"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	SiteName    string            `json:"siteName"`
	Favicon     string            `json:"favicon"`
	Raw         map[string]string `json:"raw"` // all discovered meta/property tags
}

// ExtractMedata (typo kept intentionally per user request) fetches a web page and
// extracts Open Graph, Twitter Card, standard meta tags, title, and favicon.
// Precedence order: OpenGraph > Twitter > generic <title>/<meta name> fallback.
// The caller should provide a context with timeout/cancel.
func ExtractMedata(ctx context.Context, rawURL string) (*LinkMetadata, error) { //nolint: revive // name kept as requested
	if strings.TrimSpace(rawURL) == "" {
		return nil, errors.New("empty url")
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = "https"
		rawURL = u.String()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	// Friendly UA – some sites gate metadata on UA.
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; LinkMetadataBot/1.0; +https://example.com/bot)")
	// Accept HTML types explicitly.
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Limit read size to prevent excessive memory usage.
	const maxBytes = 2 << 20 // 2 MiB
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBytes))
	if err != nil {
		return nil, err
	}

	root, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	md := &LinkMetadata{URL: rawURL, Raw: make(map[string]string)}

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if md.Title == "" && n.FirstChild != nil {
					md.Title = strings.TrimSpace(n.FirstChild.Data)
				}
			case "meta":
				var name, property, content string
				for _, a := range n.Attr {
					lk := strings.ToLower(a.Key)
					switch lk {
					case "name":
						name = strings.ToLower(a.Val)
					case "property":
						property = strings.ToLower(a.Val)
					case "content":
						content = strings.TrimSpace(a.Val)
					}
				}
				if content != "" {
					if property != "" {
						md.Raw[property] = content
					} else if name != "" {
						md.Raw[name] = content
					}
				}
			case "link":
				var relVal, hrefVal string
				for _, a := range n.Attr {
					lk := strings.ToLower(a.Key)
					switch lk {
					case "rel":
						relVal = strings.ToLower(a.Val)
					case "href":
						hrefVal = a.Val
					}
				}
				if hrefVal != "" && md.Favicon == "" && relHasIcon(relVal) {
					md.Favicon = resolveURL(u, hrefVal)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(root)

	// Precedence.
	md.Title = firstNonEmpty(md.Raw["og:title"], md.Raw["twitter:title"], md.Title, md.Raw["title"])
	md.Description = firstNonEmpty(md.Raw["og:description"], md.Raw["twitter:description"], md.Raw["description"]) // meta description
	md.Image = absolutize(u, firstNonEmpty(md.Raw["og:image"], md.Raw["twitter:image"], md.Raw["twitter:image:src"]))
	md.SiteName = firstNonEmpty(md.Raw["og:site_name"], md.Raw["twitter:site"]) // twitter:site often like @handle

	if md.Favicon == "" {
		md.Favicon = u.Scheme + "://" + u.Host + "/favicon.ico"
	}
	return md, nil
}

// Utility helpers.
func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func relHasIcon(rel string) bool {
	if rel == "" {
		return false
	}
	parts := strings.Fields(rel)
	for _, p := range parts {
		switch p {
		case "icon", "shortcut", "shortcut icon", "apple-touch-icon", "apple-touch-icon-precomposed", "mask-icon":
			return true
		}
	}
	return false
}

func resolveURL(base *url.URL, ref string) string {
	if ref == "" {
		return ref
	}
	r, err := url.Parse(ref)
	if err != nil {
		return ref
	}
	if r.IsAbs() {
		return r.String()
	}
	return base.ResolveReference(r).String()
}

func absolutize(base *url.URL, v string) string {
	if v == "" {
		return v
	}
	return resolveURL(base, v)
}
