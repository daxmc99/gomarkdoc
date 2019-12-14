package format

import (
	"fmt"
	"regexp"
	"strings"
)

// GitHubFlavoredMarkdown provides a Format which is compatible with GitHub
// Flavored Markdown's syntax and semantics. See GitHub's documentation for
// more details about their markdown format:
// https://guides.github.com/features/mastering-markdown/
type GitHubFlavoredMarkdown struct{}

// Bold converts the provided text to bold
func (f *GitHubFlavoredMarkdown) Bold(text string) (string, error) {
	return bold(text), nil
}

// CodeBlock wraps the provided code as a code block and tags it with the
// provided language (or no language if the empty string is provided).
func (f *GitHubFlavoredMarkdown) CodeBlock(language, code string) (string, error) {
	return gfmCodeBlock(language, code), nil
}

// Header converts the provided text into a header of the provided level. The
// level is expected to be at least 1.
func (f *GitHubFlavoredMarkdown) Header(level int, text string) (string, error) {
	return header(level, text)
}

var (
	gfmWhitespaceRegex = regexp.MustCompile(`\s`)
	gfmRemoveRegex     = regexp.MustCompile(`[^\pL-_\d]+`)
)

// LocalHref generates an href for navigating to a header with the given
// headerText located within the same document as the href itself.
func (f *GitHubFlavoredMarkdown) LocalHref(headerText string) (string, error) {
	result := strings.ToLower(headerText)
	result = strings.TrimSpace(result)
	result = gfmWhitespaceRegex.ReplaceAllString(result, "-")
	result = gfmRemoveRegex.ReplaceAllString(result, "")

	return fmt.Sprintf("#%s", result), nil
}

// Link generates a link with the given text and href values.
func (f *GitHubFlavoredMarkdown) Link(text, href string) (string, error) {
	return link(text, href), nil
}

// ListEntry generates an unordered list entry with the provided text at the
// provided zero-indexed depth. A depth of 0 is considered the topmost level of
// list.
func (f *GitHubFlavoredMarkdown) ListEntry(depth int, text string) (string, error) {
	return listEntry(depth, text), nil
}

// Accordion generates a collapsible content. The accordion's visible title
// while collapsed is the provided title and the expanded content is the body.
func (f *GitHubFlavoredMarkdown) Accordion(title, body string) (string, error) {
	return gfmAccordion(title, body), nil
}

// AccordionHeader generates the header visible when an accordion is collapsed.
//
// The AccordionHeader is expected to be used in conjunction with
// AccordionTerminator() when the demands of the body's rendering requires it to
// be generated independently. The result looks conceptually like the following:
//
//	accordion := format.AccordionHeader("Accordion Title") + "Accordion Body" + format.AccordionTerminator()
func (f *GitHubFlavoredMarkdown) AccordionHeader(title string) (string, error) {
	return gfmAccordionHeader(title), nil
}

// AccordionTerminator generates the code necessary to terminate an accordion
// after the body. It is expected to be used in conjunction with
// AccordionHeader(). See AccordionHeader for a full description.
func (f *GitHubFlavoredMarkdown) AccordionTerminator() (string, error) {
	return gfmAccordionTerminator(), nil
}

// Paragraph formats a paragraph with the provided text as the contents.
func (f *GitHubFlavoredMarkdown) Paragraph(text string) (string, error) {
	return paragraph(text), nil
}
