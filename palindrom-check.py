# O(n)
def is_palindrome(input):
	if input < 0:
		return False
	divisor = 1
	while(input/divisor >= 10):
		divisor *= 10

	# print("divisor: ", divisor)
	while(input > 0):
		leading = input // divisor
		trailing = input % 10

		# print("leading:", leading)
		# print("trailing:", trailing)

		if leading!=trailing:
			return False

		input = (input % divisor) // 10
		# print("input:", input)

		divisor = divisor / 100

	return True

print("-121 -> ",is_palindrome(-121))
print("0 -> ",is_palindrome(0))
print("1221 -> ",is_palindrome(1221))
print("9010109 -> ",is_palindrome(9010109))
print("9010209 -> ",is_palindrome(9010209))