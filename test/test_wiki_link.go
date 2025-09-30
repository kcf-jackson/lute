package main

import (
	"fmt"

	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New()

	// Test wiki link in IR mode
	markdown := "This is a [[Category:Examples]] test."
	fmt.Println("Original markdown:", markdown)

	// Convert to IR DOM
	irDOM := luteEngine.Md2VditorIRDOM(markdown)
	fmt.Println("IR DOM:", irDOM)

	// Convert back to markdown (this should preserve the wiki link)
	backToMarkdown := luteEngine.VditorIRDOM2Md(irDOM)
	fmt.Println("Back to markdown:", backToMarkdown)

	// Test the problematic case: add space and remove it
	irDOMWithSpace := irDOM + " "
	fmt.Println("IR DOM with space:", irDOMWithSpace)

	// Remove the space by trimming (simulate backspace)
	irDOMTrimmed := irDOM // This simulates the problematic operation
	backToMarkdownTrimmed := luteEngine.VditorIRDOM2Md(irDOMTrimmed)
	fmt.Println("After trim operation:", backToMarkdownTrimmed)

	// Test SpinVditorIRDOM (which is used when editing)
	spinResult := luteEngine.SpinVditorIRDOM(irDOM)
	fmt.Println("Spin result:", spinResult)
}