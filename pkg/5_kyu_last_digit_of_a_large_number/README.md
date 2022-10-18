Define a function that takes in two non-negative integers aaa and bbb and returns the last decimal digit of aba^ba
b
. Note that aaa and bbb may be very large!

For example, the last decimal digit of 979^79
7
is 999, since 97=47829699^7 = 47829699
7
=4782969. The last decimal digit of (2200)2300({2^{200}})^{2^{300}}(2
200
)
2
300

, which has over 109210^{92}10
92
decimal digits, is 666. Also, please take 000^00
0
to be 111.

You may assume that the input will always be valid.

Examples
lastDigit 4 1             `shouldBe` 4
lastDigit 4 2             `shouldBe` 6
lastDigit 9 7             `shouldBe` 9
lastDigit 10 (10^10)      `shouldBe` 0
lastDigit (2^200) (2^300) `shouldBe` 6