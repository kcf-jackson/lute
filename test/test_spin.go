package main

import (
	"fmt"
	"strings"
	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New()
	
	fmt.Println("=== Testing with <wbr> tags ===")
	
	// Test SpinVditorIRDOM with caret in different positions
	
	// Test 1: Link with caret
	linkHTML := `<p data-block="0"><span data-type="a" class="vditor-ir__node"><span class="vditor-ir__marker vditor-ir__marker--bracket">[</span><span class="vditor-ir__link">text</span><span class="vditor-ir__marker vditor-ir__marker--bracket">]</span><span class="vditor-ir__marker vditor-ir__marker--paren">(</span><span class="vditor-ir__marker vditor-ir__marker--link">url<wbr></span><span class="vditor-ir__marker vditor-ir__marker--paren">)</span></span></p>`
	
	fmt.Println("Link with <wbr>:")
	spinResult := luteEngine.SpinVditorIRDOM(linkHTML)
	fmt.Println(spinResult)
	
	// Test 2: Wiki link with caret
	wikiHTML := `<p data-block="0"><span data-type="wiki-link" class="vditor-ir__node"><span class="vditor-ir__marker vditor-ir__marker--bracket">[[</span><span class="vditor-ir__link">Category:Examples<wbr></span><span class="vditor-ir__marker vditor-ir__marker--bracket">]]</span></span></p>`
	
	fmt.Println("\nWiki link with <wbr>:")
	spinResult2 := luteEngine.SpinVditorIRDOM(wikiHTML)
	fmt.Println(spinResult2)
	
	// Check for expand class
	checkExpandClass := func(html, desc string) {
		if strings.Contains(html, "vditor-ir__node--expand") {
			fmt.Printf("✓ %s has expand class\n", desc)
		} else {
			fmt.Printf("✗ %s missing expand class\n", desc)
		}
	}
	
	checkExpandClass(spinResult, "Link")
	checkExpandClass(spinResult2, "Wiki link")
}