# merge two sorted list of integer
# time complexity -> O(A+B)
	
# with additional space:result
def merge_sorted_list(numbers_a, numbers_b):
	a = 0
	b = 0
	result = []
	while (a < len(numbers_a) or b < len(numbers_b)):
		if a == len(numbers_a):
			result.append(numbers_b[b])
			b+=1
			continue
		if b == len(numbers_b):
			result.append(numbers_a[a])
			a+=1
			continue

		if (a < len(numbers_a) and numbers_a[a] < numbers_b[b]):
			result.append(numbers_a[a])
			a+=1
		elif (b < len(numbers_b) and numbers_b[b] < numbers_a[a]):
			result.append(numbers_b[b])
			b+=1
	return result

# without additional space
def merge_sorted_list_in_place(numbers_a, numbers_b):
	# use binary tree
	return [1,2,3,4,5,6]

print("[1,4,6],[2,3,5,7,8] -> ",merge_sorted_list([1,4,6],[2,3,5,7,8]))
print("[5,7,9],[4,6,8] -> ",merge_sorted_list([5,7,9],[4,6,8]))
print("[5,10,15],[2,3,20] -> ",merge_sorted_list([5,10,15],[2,3,20]))
print("[2,3,5,7,8],[1,4,6] -> ",merge_sorted_list([2,3,5,7,8],[1,4,6]))