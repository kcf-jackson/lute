// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"testing"

	"github.com/88250/lute"
)

var wikiLinkTests = []parseTest{
	{"wiki-basic", "[[file_name]]", "<p><a href=\"file_name\" class=\"wiki-link\">file_name</a></p>\n"},
	{"wiki-section", "[[file_name#section]]", "<p><a href=\"file_name#section\" class=\"wiki-link\">file_name#section</a></p>\n"},
	{"wiki-display", "[[file_name|display text]]", "<p><a href=\"file_name\" class=\"wiki-link\">display text</a></p>\n"},
	{"wiki-full", "[[file_name#section|display text]]", "<p><a href=\"file_name#section\" class=\"wiki-link\">display text</a></p>\n"},
	{"wiki-context", "This is a [[wiki link]] in text", "<p>This is a <a href=\"wiki link\" class=\"wiki-link\">wiki link</a> in text</p>\n"},
	{"wiki-multiple", "Multiple [[link1]] and [[link2|text]] here", "<p>Multiple <a href=\"link1\" class=\"wiki-link\">link1</a> and <a href=\"link2\" class=\"wiki-link\">text</a> here</p>\n"},
}

func TestWikiLink(t *testing.T) {
	luteEngine := lute.New()

	for _, test := range wikiLinkTests {
		html := luteEngine.MarkdownStr("", test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal markdown text\n\t%q", test.name, test.to, html, test.from)
		}
	}
}

func TestWikiLinkSerialization(t *testing.T) {
	luteEngine := lute.New()

	testCases := []string{
		"[[file_name]]",
		"[[file_name#section]]", 
		"[[file_name|display text]]",
		"[[file_name#section|display text]]",
		"This is a [[wiki link]] in text",
		"Multiple [[link1]] and [[link2|text]] here",
	}

	for _, testCase := range testCases {
		formatted := luteEngine.FormatStr("", testCase)
		expectedFormat := testCase + "\n" // Format typically adds a newline
		if expectedFormat != formatted {
			t.Fatalf("serialization test failed\nexpected\n\t%q\ngot\n\t%q", expectedFormat, formatted)
		}
	}
}