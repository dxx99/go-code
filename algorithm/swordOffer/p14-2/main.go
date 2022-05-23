package main


func main() {

}

func cuttingRope(n int) int {
	const mod = 1e9+7

	if n <= 3 {
		return 1*(n-1)
	}

	ans := 1
	for n > 4 {
		ans *= 3
		if ans > mod {
			ans %= mod
		}
		n -= 3
	}
	ans *= n

	return ans % mod
}
