# O(n)
def is_bracket_balanced(input):
	stack = []
	brackets = {"(":")","[":"]","{":"}"}
	for i in input:
		if i in brackets.keys(): #match open bracket, push to stack
			stack.append(i)
		else:
			if not stack:
				return False
			latest_bracket_open = stack.pop()
			bracket_closed = brackets[latest_bracket_open]
			if bracket_closed != i:
				return False

	if stack:
		return False
	return True

print("({}) -> ",is_bracket_balanced("({})"))
print("[()]{}{[()()]()} -> ",is_bracket_balanced("[()]{}{[()()]()}"))
print("[(]) -> ",is_bracket_balanced("[(])"))