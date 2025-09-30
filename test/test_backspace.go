package main

import (
	"fmt"
	"strings"
	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New()
	
	fmt.Println("=== Simulating the exact user scenario ===")
	
	// Step 1: Start with wiki link
	markdown := "[[Category:Examples]]"
	fmt.Printf("1. Original markdown: %s\n", markdown)
	
	// Step 2: Convert to IR DOM (what happens when rendering)
	irDOM := luteEngine.Md2VditorIRDOM(markdown)
	fmt.Printf("2. IR DOM: %s\n", irDOM)
	
	// Step 3: Simulate adding a space after the wiki link (user types space)
	// In the actual editor, this would be done by inserting a text node
	irDOMWithSpace := strings.Replace(irDOM, "</span></p>", "</span> </p>", 1)
	fmt.Printf("3. After adding space: %s\n", irDOMWithSpace)
	
	// Step 4: Convert to markdown (what happens during the SpinVditorIRDOM)
	markdownWithSpace := luteEngine.VditorIRDOM2Md(irDOMWithSpace)
	fmt.Printf("4. Markdown with space: [%s]\n", markdownWithSpace)
	
	// Step 5: Convert back to IR DOM (re-rendering after change)
	newIRDOM := luteEngine.Md2VditorIRDOM(markdownWithSpace)
	fmt.Printf("5. Re-rendered IR DOM: %s\n", newIRDOM)
	
	// Step 6: Simulate pressing backspace to remove the space
	// This would involve DOM manipulation, but we'll simulate by converting back
	irDOMAfterBackspace := strings.Replace(newIRDOM, " </p>", "</p>", 1)
	fmt.Printf("6. After backspace: %s\n", irDOMAfterBackspace)
	
	// Step 7: Convert back to markdown (what happens after backspace)
	finalMarkdown := luteEngine.VditorIRDOM2Md(irDOMAfterBackspace)
	fmt.Printf("7. Final markdown: [%s]\n", finalMarkdown)
	
	// Check if the wiki link survived the round trip
	if strings.Contains(finalMarkdown, "[[Category:Examples]]") {
		fmt.Println("✓ SUCCESS: Wiki link preserved through space add/remove cycle")
	} else {
		fmt.Printf("✗ FAILURE: Wiki link corrupted. Got: %s\n", finalMarkdown)
	}
	
	// Test the SpinVditorIRDOM function specifically
	fmt.Println("\n=== Testing SpinVditorIRDOM directly ===")
	spinResult := luteEngine.SpinVditorIRDOM(irDOMWithSpace)
	spinMarkdown := luteEngine.VditorIRDOM2Md(spinResult)
	fmt.Printf("Spin result: [%s]\n", spinMarkdown)
	
	if strings.Contains(spinMarkdown, "[[Category:Examples]]") {
		fmt.Println("✓ SUCCESS: SpinVditorIRDOM preserves wiki links")
	} else {
		fmt.Printf("✗ FAILURE: SpinVditorIRDOM corrupted wiki link. Got: %s\n", spinMarkdown)
	}
}