// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`"capiTalIze tHe titLe"`, 
			`"Capitalize The Title"`,
		},
		{
			`"First leTTeR of EACH Word"`, 
			`"First Letter of Each Word"`,
		},
		{
			`"i lOve leetcode"`, 
			`"i Love Leetcode"`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, capitalizeTitle, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-69/problems/capitalize-the-title/