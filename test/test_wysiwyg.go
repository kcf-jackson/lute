package main

import (
	"fmt"
	"strings"
	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New()
	
	fmt.Println("=== Testing WYSIWYG mode ===")
	
	// Test simple wiki link
	testWysiwygWikiLink(luteEngine, "[[Category:Examples]]")
	
	// Test wiki link with display text
	testWysiwygWikiLink(luteEngine, "[[Category:Examples|Examples]]")
	
	// Test multiple wiki links
	testWysiwygWikiLink(luteEngine, "[[Category:Examples]] and [[Page:Test]]")
}

func testWysiwygWikiLink(luteEngine *lute.Lute, markdown string) {
	fmt.Printf("\\nTesting: %s\\n", markdown)
	
	// Convert to WYSIWYG DOM
	wysiwygDOM := luteEngine.Md2VditorDOM(markdown)
	fmt.Printf("WYSIWYG DOM: %s\\n", wysiwygDOM)
	
	// Convert back to markdown 
	backToMarkdown := luteEngine.VditorDOM2Md(wysiwygDOM)
	fmt.Printf("Back to markdown: [%s]\\n", backToMarkdown)
	
	// Test SpinVditorDOM
	spinResult := luteEngine.SpinVditorDOM(wysiwygDOM)
	spinMarkdown := luteEngine.VditorDOM2Md(spinResult)
	fmt.Printf("Spin result: [%s]\\n", spinMarkdown)
	
	// Check if original pattern is preserved
	isPreserved := strings.Contains(backToMarkdown, markdown) && strings.Contains(spinMarkdown, markdown)
	if isPreserved {
		fmt.Println("✓ SUCCESS: Wiki link preserved")
	} else {
		fmt.Printf("✗ FAILURE: Wiki link corrupted\\n")
	}
}