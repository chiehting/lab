package main

import (
	"fmt"
	"time"
)

func longestPalindrome1(s string) string {
	maxLength := 0
	ans := ""
	// x - i = j - 1
	for c := 0; c < len(s); c++ {
		for l := 0; l < len(s); l++ {
			leftA := c - l // "aba"
			rightA := c + l
			if leftA < 0 || rightA+1 > len(s) {
				break
			}
			if s[leftA] != s[rightA] {
				break
			}
			if l+1 > maxLength {
				maxLength = l + 1
				ans = s[leftA : rightA+1]
			}
		}
		for l := 0; l < len(s); l++ {
			leftB := c - l // "abba"
			rightB := c + l + 1
			if leftB < 0 || rightB+1 > len(s) {
				break
			}
			if s[leftB] != s[rightB] {
				break
			}

			if l+2 > maxLength {
				maxLength = l + 2
				ans = s[leftB : rightB+1]
			}
		}
	}
	return ans
}

func main() {
	t1 := time.Now()
	s := "mqizdjrfqtmcsruvvlhdgzfrmxgmmbguroxcbhalzggxhzwfznfkrdwsvzhieqvsrbyedqxwmnvovvnesphgddoikfwuujrhxwcrbttfbmlayrlmpromlzwzrkjkzdvdkpqtbzszrngczvgspzpfnvwuifzjdrmwfadophxscxtbavrhfkadhxrmvlmofbzqshqxazzwjextdpuszwgrxirmmlqitjjpijptmqfbggkwaolpbdglmsvlwdummsrdyjhmgrasrblpjsrpkkgknsucsshjuxunqiouzrdwwooxclutkrujpfebjpoodvhknayilcxjrvnykfjhvsikjabsdnvgguoiyldshbsmsrrlwmkfmyjbbsylhrusubcglaemnurpuvlyyknbqelmkkyamrcmjbncpafchacckhymtasylyfjuribqxsekbjkgzrvzjmjkquxfwopsbjudggnfbuyyfizefgxamocxjgkwxidkgursrcsjwwyeiymoafgyjlhtcdkgrikzzlenqgtdukivvdsalepyvehaklejxxmmoycrtsvzugudwirgywvsxqapxyjedbdhvkkvrxxsgifcldkspgdnjnnzfalaslwqfylmzvbxuscatomnmgarkvuccblpoktlpnazyeazhfucmfpalbujhzbykdgcirnqivdwxnnuznrwdjslwdwgpvjehqcbtjljnxsebtqujhmteknbinrloregnphwhnfidfsqdtaexencwzszlpmxjicoduejjomqzsmrgdgvlrfcrbyfutidkryspmoyzlgfltclmhaeebfbunrwqytzhuxghxkfwtjrfyxavcjwnvbaydjnarrhiyjavlmfsstewtxrcifcllnugldnfyswnsewqwnvbgtatccfeqyjgqbnufwttaokibyrldhoniwqsflvlwnjmffoirzmoxqxunkuepj"
	s = "abba"
	a := longestPalindrome1(s)
	fmt.Println(a)
	fmt.Println(time.Now().Sub(t1))
}
