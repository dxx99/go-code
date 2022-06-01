package main

func main() {

}

func majorityElement(nums []int) int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
		if hash[nums[i]] > (len(nums)>>1) {
			return nums[i]
		}
	}

	return 0
}
