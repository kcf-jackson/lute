package main

import (
	"fmt"
	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New()
	
	// Test simple wiki link
	fmt.Println("=== Test 1: Simple wiki link ===")
	testWikiLink(luteEngine, "[[Category:Examples]]")
	
	// Test wiki link with display text  
	fmt.Println("\n=== Test 2: Wiki link with display text ===")
	testWikiLink(luteEngine, "[[Category:Examples|Examples]]")
	
	// Test multiple wiki links
	fmt.Println("\n=== Test 3: Multiple wiki links ===")
	testWikiLink(luteEngine, "[[Category:Examples]] and [[Page:Test]]")
}

func testWikiLink(luteEngine *lute.Lute, markdown string) {
	// Step 1: Convert to IR DOM
	irDOM := luteEngine.Md2VditorIRDOM(markdown)
	fmt.Printf("Original: %s\n", markdown)
	
	// Step 2: Simulate adding a space 
	irDOMWithSpace := irDOM[:len(irDOM)-4] + " " + irDOM[len(irDOM)-4:]
	
	// Step 3: Convert back to markdown (this is what happens during editing)
	backToMarkdown := luteEngine.VditorIRDOM2Md(irDOMWithSpace)
	fmt.Printf("After space manipulation: [%s]\n", backToMarkdown)
	
	// Step 4: Test the SpinVditorIRDOM function 
	spinResult := luteEngine.SpinVditorIRDOM(irDOMWithSpace)
	finalMarkdown := luteEngine.VditorIRDOM2Md(spinResult)
	fmt.Printf("After spin: [%s]\n", finalMarkdown)
	
	// Check if wiki links are preserved
	if backToMarkdown == finalMarkdown && backToMarkdown != "" {
		fmt.Println("✓ Wiki links preserved correctly")
	} else {
		fmt.Println("✗ Wiki links not preserved correctly")
	}
}