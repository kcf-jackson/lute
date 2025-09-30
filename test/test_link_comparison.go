package main

import (
	"fmt"
	"strings"
	"github.com/88250/lute"
	"github.com/88250/lute/editor"
)

func main() {
	luteEngine := lute.New()
	
	fmt.Println("=== Comparing links vs wiki links ===")
	
	// Test regular links with caret in different positions
	linkTests := []string{
		"[te" + editor.Caret + "xt](url)",  // Caret in text
		"[text](ur" + editor.Caret + "l)",  // Caret in URL
	}
	
	// Test wiki links with caret in different positions  
	wikiTests := []string{
		"[[Category:Exa" + editor.Caret + "mples]]",  // Caret in target
		"[[Category:Examples|Te" + editor.Caret + "xt]]",  // Caret in text
	}
	
	fmt.Println("Regular links:")
	for i, test := range linkTests {
		fmt.Printf("Test %d: %s\n", i+1, test)
		irDOM := luteEngine.Md2VditorIRDOM(test)
		
		if strings.Contains(irDOM, "vditor-ir__node--expand") {
			fmt.Println("✓ Has expand class")
		} else {
			fmt.Println("✗ No expand class")
		}
		fmt.Printf("Result: %s\n\n", irDOM)
	}
	
	fmt.Println("Wiki links:")
	for i, test := range wikiTests {
		fmt.Printf("Test %d: %s\n", i+1, test)
		irDOM := luteEngine.Md2VditorIRDOM(test)
		
		if strings.Contains(irDOM, "vditor-ir__node--expand") {
			fmt.Println("✓ Has expand class")
		} else {
			fmt.Println("✗ No expand class")
		}
		fmt.Printf("Result: %s\n\n", irDOM)
	}
}