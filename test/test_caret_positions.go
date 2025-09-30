package main

import (
	"fmt"
	"strings"
	"github.com/88250/lute"
	"github.com/88250/lute/editor"
)

func main() {
	luteEngine := lute.New()
	
	fmt.Println("=== Testing text extraction ===")
	
	// Let's try to understand what causes expand mode
	// by examining the internal structure
	
	// Test with caret at the beginning
	testCases := []string{
		editor.Caret + "[[Category:Examples]]",  // Caret before
		"[[" + editor.Caret + "Category:Examples]]",  // Caret after opening brackets  
		"[[Category:Examples" + editor.Caret + "]]",  // Caret in target
		"[[Category:Examples]]" + editor.Caret,  // Caret after
		"[[Category:Examples|" + editor.Caret + "Text]]",  // Caret in text
	}
	
	for i, testCase := range testCases {
		fmt.Printf("\nTest %d: %s\n", i+1, testCase)
		irDOM := luteEngine.Md2VditorIRDOM(testCase)
		
		// Check for expand class
		if strings.Contains(irDOM, "vditor-ir__node--expand") {
			fmt.Println("✓ Has expand class")
		} else {
			fmt.Println("✗ No expand class")
		}
		
		// Show the result
		fmt.Printf("Result: %s\n", irDOM)
	}
}