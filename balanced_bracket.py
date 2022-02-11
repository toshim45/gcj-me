def is_bracket_balanced(input):
	stack = []
	brackets = {"(":")","[":"]","{":"}"}
	for i in input:
		if i in brackets.keys():
			stack.append(i)
		else:
			if not stack:
				return False
			

	if stack:
		return False
	return True

print("({}) -> ",is_bracket_balanced("({})"))