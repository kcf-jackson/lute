package main

import (
	"fmt"
	"strings"
	"github.com/88250/lute"
	"github.com/88250/lute/editor"
)

func main() {
	luteEngine := lute.New()
	
	fmt.Println("=== Testing expand mode behavior ===")
	
	// Test 1: Link with caret inside
	linkWithCaret := "[text](url" + editor.Caret + ")"
	fmt.Printf("Link with caret: %s\n", linkWithCaret)
	linkIRDOM := luteEngine.Md2VditorIRDOM(linkWithCaret)
	fmt.Printf("Link IR DOM: %s\n", linkIRDOM)
	
	// Test 2: Wiki link with caret inside  
	wikiWithCaret := "[[Category:Examples" + editor.Caret + "]]"
	fmt.Printf("\nWiki link with caret: %s\n", wikiWithCaret)
	wikiIRDOM := luteEngine.Md2VditorIRDOM(wikiWithCaret)
	fmt.Printf("Wiki link IR DOM: %s\n", wikiIRDOM)
	
	// Test 3: Wiki link with caret inside display text
	wikiWithCaretText := "[[Category:Examples|Exam" + editor.Caret + "ples]]"
	fmt.Printf("\nWiki link with caret in text: %s\n", wikiWithCaretText)
	wikiTextIRDOM := luteEngine.Md2VditorIRDOM(wikiWithCaretText)
	fmt.Printf("Wiki link text IR DOM: %s\n", wikiTextIRDOM)
	
	// Check if expand class is present
	checkExpandClass := func(html, desc string) {
		if strings.Contains(html, "vditor-ir__node--expand") {
			fmt.Printf("✓ %s has expand class\n", desc)
		} else {
			fmt.Printf("✗ %s missing expand class\n", desc)
		}
	}
	
	checkExpandClass(linkIRDOM, "Link")
	checkExpandClass(wikiIRDOM, "Wiki link")
	checkExpandClass(wikiTextIRDOM, "Wiki link with text")
}