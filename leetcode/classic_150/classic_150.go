package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			nums[i] = nums[j]
			j--
		} else {
			i++
		}
	}
	return i
}

func removeDuplicates(nums []int) int {
	remove := func(nums []int, k int) int {
		i, j := k, k
		for j < len(nums) {
			if nums[j] != nums[j-k] {
				nums[i] = nums[j]
				i++
			}
			j++
		}
		return i
	}
	return remove(nums, 1)
}

func removeDuplicatesII(nums []int) int {
	remove := func(nums []int, k int) int {
		if len(nums) < k {
			return len(nums)
		}
		i, j := k, k
		for j < len(nums) {
			if nums[j] != nums[i-k] {
				nums[i] = nums[j]
				i++
			}
			j++
		}
		return i
	}
	return remove(nums, 2)
}

func majorityElement(nums []int) int {
	var res int
	var cnt int
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			res = nums[i]
		}
		if res == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse := func(nums []int) {
		l, r := 0, len(nums)-1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func maxProfit(prices []int) int {
	var cost = prices[0]
	var res int
	for i := 1; i < len(prices); i++ {
		cost = min(cost, prices[i])
		res = max(res, prices[i]-cost)
	}
	return res
}

func maxProfitII(prices []int) int {
	var res int
	for i := 1; i < len(prices); i++ {
		res += max(0, prices[i]-prices[i-1])
	}
	return res
}

func canJump(nums []int) bool {
	var maxPosition int
	for i := 0; i < len(nums); i++ {
		if i > maxPosition {
			return false
		}
		maxPosition = max(maxPosition, i+nums[i])
		if maxPosition >= len(nums)-1 {
			return true
		}
	}
	return false
}

func jump(nums []int) int {
	var res, maxP, end int
	for i := 0; i < len(nums)-1; i++ {
		maxP = max(maxP, i+nums[i])
		if end == i {
			res++
			end = maxP
		}
	}
	return res
}

func hIndex(citations []int) int {
	slices.SortFunc(citations, func(a, b int) int {
		return b - a
	})
	var res int
	for i := 0; i < len(citations); i++ {
		if citations[i] > i {
			res = i + 1
		}
	}
	return res
}

type RandomizedSet struct {
	nums []int
	idx  map[int]int
}

func ConstructorRandomizedSet() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		idx:  map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; !ok {
		this.nums = append(this.nums, val)
		this.idx[val] = len(this.nums) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.idx[val]
	if !ok {
		return false
	}
	newVal := this.nums[len(this.nums)-1]
	this.idx[newVal] = id
	this.nums[id] = newVal
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.idx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	fmt.Println(res)
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = r * res[i]
		r *= nums[i]
	}
	return res
}

func canCompleteCircuit(gas []int, cost []int) int {
	var totalGas, totalCost int
	var curGas, startIdx int
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		curGas += gas[i] - cost[i]
		if curGas < 0 {
			startIdx = i + 1
			curGas = 0
		}
	}
	if totalGas < totalCost {
		return -1
	}
	return startIdx
}

func candy(ratings []int) int {
	l, r := make([]int, len(ratings)), make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		l[i] = 1
		r[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			l[i] = l[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			r[i] = r[i+1] + 1
		}
	}
	var res int
	for i := 0; i < len(ratings); i++ {
		res += max(l[i], r[i])
	}
	return res
}

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	num := m[s[0]]
	var sum int
	for i := 1; i < len(s); i++ {
		nextNum := m[s[i]]
		if num < nextNum {
			sum -= num
		} else {
			sum += num
		}
		num = nextNum
	}
	return sum + num
}

func intToRoman(num int) string {
	intRoman := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	sortedValue := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var res string
	for i := 0; i < len(sortedValue); i++ {
		for num >= sortedValue[i] {
			num -= sortedValue[i]
			res += intRoman[sortedValue[i]]
		}
		if num == 0 {
			break
		}
	}
	return res
}

func lengthOfLastWord(s string) int {
	var res int
	j := len(s) - 1
	for j >= 0 && s[j] == ' ' {
		j--
	}
	for j >= 0 && s[j] != ' ' {
		res++
		j--
	}
	return res
}

func longestCommonPrefix(strs []string) string {
	lcp := func(ls, rs string) string {
		var i int
		for i < len(ls) && i < len(rs) && ls[i] == rs[i] {
			i++
		}
		return ls[:i]
	}
	var partition func(strs []string, l, r int) string
	partition = func(strs []string, l, r int) string {
		if l >= r {
			return strs[l]
		}
		mid := (l + r) >> 1
		ls := partition(strs, l, mid)
		rs := partition(strs, mid+1, r)
		return lcp(ls, rs)
	}
	return partition(strs, 0, len(strs)-1)
}

func reverseWords(s string) string {
	fields := strings.Fields(s)
	i, j := 0, len(fields)-1
	for i < j {
		fields[i], fields[j] = fields[j], fields[i]
		i++
		j--
	}
	return strings.Join(fields, " ")
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strs := make([]string, numRows)
	var rowIdx int
	goingDown := false
	for i := 0; i < len(s); i++ {
		strs[rowIdx] += string(s[i])
		if rowIdx == 0 || rowIdx == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			rowIdx++
		} else {
			rowIdx--
		}
	}
	return strings.Join(strs, "")
}

func strStr(haystack string, needle string) int {
	buildLPS := func(pattern string) []int {
		lps := make([]int, len(pattern))
		lps[0] = 0
		i, length := 1, 0
		for i < len(pattern) {
			if pattern[i] == pattern[length] {
				length++
				lps[i] = length
				i++
				//0,1,0,1,2,0,0,0
				//a,a,b,a,a,c,c,c
			} else {
				if length != 0 {
					length = lps[length-1]
				} else {
					lps[i] = 0
					i++
				}
			}
		}
		return lps
	}
	kmpSearch := func(text, pattern string) []int {
		var res []int
		lps := buildLPS(pattern)
		i, j := 0, 0
		for i < len(text) {
			if text[i] == pattern[j] {
				i++
				j++
			}
			if j == len(pattern) {
				res = append(res, i-j)
				j = lps[j-1]
			} else if i < len(text) && pattern[j] != text[i] {
				if j != 0 {
					j = lps[j-1]
				} else {
					i++
				}
			}
		}
		return res
	}
	res := kmpSearch(haystack, needle)
	if len(res) == 0 {
		return -1
	}
	return res[0]
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		a := rune(s[i])
		b := rune(s[j])
		if !unicode.IsLetter(a) && !unicode.IsDigit(a) {
			i++
		} else if !unicode.IsLetter(b) && !unicode.IsDigit(b) {
			j--
		} else if unicode.ToLower(a) == unicode.ToLower(b) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return false
	}
	i, j := 0, 0
	for i < len(t) {
		if t[i] == s[j] {
			j++
		}
		if j == len(s) {
			return true
		}
		i++
	}
	return false
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

func maxArea(height []int) int {
	var res int
	l, r := 0, len(height)-1
	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func minSubArrayLen(target int, nums []int) int {
	var sum int
	var res = len(nums) + 1
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for l <= r && sum >= target {
			fmt.Println(sum)
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	var res int
	var duplicate = map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		duplicate[s[r]]++
		for l <= r && duplicate[s[r]] > 1 {
			duplicate[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func myPow(x float64, n int) float64 {
	var quickMul func(x float64, n int) float64
	//n=5,2,1,0
	quickMul = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		y := quickMul(x, n/2)
		if n%2 == 0 {
			return y * y
		}
		return y * y * x
	}
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
				res = max(res, f[i+1][j+1])
			}
		}
	}
	return res * res
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		f[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		f[0][j] = j
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if word1[i] == word2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i+1][j], f[i][j+1], f[i][j]) + 1
			}
		}
	}
	return f[m][n]
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, t := len(s1), len(s2), len(s3)
	if m+n != t {
		return false
	}
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[m][n]
}

// 逆波兰表达式
func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		val, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

func longestPalindrome(s string) string {
	expand := func(str string, i, j int) (int, int) {
		l, r := i, j
		for l >= 0 && r <= len(str)-1 && str[l] == str[r] {
			l--
			r++
		}
		return l + 1, r - 1
	}
	var res string
	for i := 0; i < len(s); i++ {
		l1, r1 := expand(s, i, i)
		if len(res) < r1-l1+1 {
			res = s[l1 : r1+1]
		}
		l2, r2 := expand(s, i, i+1)
		if len(res) < r2-l2+1 {
			res = s[l2 : r2+1]
		}
	}
	return res
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([][]int, m)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		f[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		f[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				f[i][j] = grid[i][j]
			} else if i == 0 {
				f[i][j] = f[i][j-1] + grid[i][j]
			} else if j == 0 {
				f[i][j] = f[i-1][j] + grid[i][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
			}
		}
	}
	return f[m-1][n-1]
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	p, q := x, 0
	for p > 0 {
		q = q*10 + p%10
		p /= 10
	}
	return x == q
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func trailingZeroes(n int) int {
	var res int
	pow := 5
	for n/pow > 0 {
		res += n / pow
		pow *= 5
	}
	return res
}

// 数字二进制的公共前缀
func rangeBitwiseAnd(l int, r int) int {
	var shift int
	for l != r {
		l, r = l>>1, r>>1
		shift++
	}
	return r << shift
}

func singleNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func singleNumberII(nums []int) int {
	once, twice := 0, 0
	for i := 0; i < len(nums); i++ {
		once = (once ^ nums[i]) &^ twice
		twice = (twice ^ nums[i]) &^ once
	}
	return once
}

func hammingWeight(n int) int {
	var res int
	for n != 0 {
		res++
		n &= n - 1
	}
	return res
}

func reverseBits(n uint32) uint32 {
	var res uint32
	for i := 0; i < 32 && n > 0; i++ {
		res |= n & 1 << (31 - i)
		n >>= 1
	}
	return res
}

func addBinary(a string, b string) string {
	toVal := func(c byte) int {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
		return int(c - 'A' + 10)
	}
	toChar := func(n int) byte {
		if n < 10 {
			return byte('0' + n)
		}
		return byte('A' + n - 10)
	}
	add := func(a string, b string, numeration int) string {
		var res []byte
		var carry int
		i, j := len(a)-1, len(b)-1
		for i >= 0 || j >= 0 {
			sum := carry
			if i >= 0 {
				sum += toVal(a[i])
				i--
			}
			if j >= 0 {
				sum += toVal(b[j])
				j--
			}
			carry = sum / numeration
			sum %= numeration
			res = append([]byte{toChar(sum)}, res...)
		}
		if carry > 0 {
			res = append([]byte{toChar(carry)}, res...)
		}
		return string(res)
	}
	return add(a, b, 2)
}

type minHeap struct {
	nums []int
}

func (h *minHeap) up(i int) {
	parent, smallest := (i-1)/2, i
	//注意这里是小于等于
	if parent < 0 || h.nums[parent] <= h.nums[smallest] {
		return
	}
	h.nums[parent], h.nums[smallest] = h.nums[smallest], h.nums[parent]
	h.up(parent)
}

// append to tail->child want to up
func (h *minHeap) push(val int) {
	h.nums = append(h.nums, val)
	h.up(len(h.nums) - 1)
}

func (h *minHeap) down(i int) {
	l, r, smallest := 2*i+1, 2*i+2, i
	if l < len(h.nums) && h.nums[l] < h.nums[smallest] {
		smallest = l
	}
	if r < len(h.nums) && h.nums[r] < h.nums[smallest] {
		smallest = r
	}
	if smallest == i {
		return
	}
	h.nums[smallest], h.nums[i] = h.nums[i], h.nums[smallest]
	h.down(smallest)
}

// swap head,tail->pop tail->parent want to down
func (h *minHeap) pop() (int, bool) {
	if len(h.nums) == 0 {
		return -1, false
	}
	n := len(h.nums)
	h.nums[0], h.nums[n-1] = h.nums[n-1], h.nums[0]
	val := h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]
	h.down(0)
	return val, true
}

func (h *minHeap) peek() int {
	if len(h.nums) == 0 {
		return 0
	}
	return h.nums[0]
}

func (h *minHeap) len() int {
	return len(h.nums)
}

func findKthLargest(nums []int, k int) int {
	h := &minHeap{nums: []int{}}
	for i := 0; i < len(nums); i++ {
		if h.len() < k {
			h.push(nums[i])
		} else if h.peek() < nums[i] {
			h.pop()
			h.push(nums[i])
		}
	}
	return h.nums[0]
}

type elem struct {
	a, b int
	x, y int
}
type minElemHeap struct {
	nums []elem
}

func (h *minElemHeap) val(i int) int {
	return h.nums[i].a + h.nums[i].b
}

func (h *minElemHeap) up(i int) {
	parent, smallest := (i-1)/2, i
	if parent < 0 || h.val(parent) <= h.val(smallest) {
		return
	}
	h.nums[parent], h.nums[smallest] = h.nums[smallest], h.nums[parent]
	h.up(parent)
}

// append to tail->child want to up
func (h *minElemHeap) push(elem elem) {
	h.nums = append(h.nums, elem)
	h.up(len(h.nums) - 1)
}

func (h *minElemHeap) down(i int) {
	l, r, smallest := 2*i+1, 2*i+2, i
	if l < len(h.nums) && h.val(l) < h.val(smallest) {
		smallest = l
	}
	if r < len(h.nums) && h.val(r) < h.val(smallest) {
		smallest = r
	}
	if smallest == i {
		return
	}
	h.nums[smallest], h.nums[i] = h.nums[i], h.nums[smallest]
	h.down(smallest)
}

// swap head,tail->pop tail->parent want to down
func (h *minElemHeap) pop() (elem, bool) {
	if len(h.nums) == 0 {
		return elem{}, false
	}
	n := len(h.nums)
	h.nums[0], h.nums[n-1] = h.nums[n-1], h.nums[0]
	val := h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]
	h.down(0)
	return val, true
}

func (h *minElemHeap) peek() int {
	if len(h.nums) == 0 {
		return 0
	}
	return h.val(0)
}

func (h *minElemHeap) len() int {
	return len(h.nums)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &minElemHeap{nums: []elem{}}
	for i := 0; i < len(nums1) && h.len() <= k; i++ {
		h.push(elem{nums1[i], nums2[0], i, 0})
	}
	var res [][]int
	for h.len() > 0 && len(res) < k {
		val, _ := h.pop()
		res = append(res, []int{val.a, val.b})
		x, y := val.x, val.y
		if y+1 < len(nums2) {
			h.push(elem{nums1[x], nums2[y+1], x, y + 1})
		}
	}
	return res
}
