# string operation func is forbidden
def reverse_int(input):
	output = 0

	is_negative = False
	if input < 0:
		is_negative = True
		input *= -1

	while input > 0:
		trailing = input % 10
		leading = output * 10

		output = leading + trailing 

		if(output > 2147483647 or output <-2147483648):
			return 0
		input = input // 10

	return output*-1 if is_negative else output

print("12345 -> ",reverse_int(12345))
print("12305 -> ",reverse_int(12305))
print("56700 -> ",reverse_int(56700))
print("-12345 -> ",reverse_int(-12345))
print("-10345 -> ",reverse_int(-10345))
print("32109876543210987 -> ",reverse_int(-32109876543210987))
print("-32109876543210987 -> ",reverse_int(-32109876543210987))
