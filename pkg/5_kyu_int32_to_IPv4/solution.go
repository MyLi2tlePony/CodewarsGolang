package int32_to_IPv4

import (
	"fmt"
)

func solution(n uint32) string {
	firstOctet := uint8(n >> (3 * 8))
	secondOctet := uint8(n >> (2 * 8))
	thirdOctet := uint8(n >> (1 * 8))
	fourthOctet := uint8(n)

	return fmt.Sprintf("%d.%d.%d.%d", firstOctet, secondOctet, thirdOctet, fourthOctet)
}
