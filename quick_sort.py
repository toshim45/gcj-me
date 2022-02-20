def do_quick_sort(numbers):
	quick_sort(0, len(numbers)-1, numbers)
	return numbers


# partition
# find number smaller than pivot from left
# find number greater than pivot from right
# if start and end have not crossed each other, swap the numbers on start and end
def partition(start, end, numbers):
	pivot_idx = start
	pivot = numbers[pivot_idx]

	while start < end:
		while start < len(numbers) and numbers[start] <= pivot:
			start += 1
		while numbers[end] > pivot:
			end -= 1
		if start < end:
			numbers[start],numbers[end] = numbers[end], numbers[start]

	numbers[end], numbers[pivot_idx] = numbers[pivot_idx], numbers[end]
	return end

def quick_sort(start, end, numbers):
	if start < end:
		part_idx = partition(start, end, numbers)
		quick_sort(start, part_idx-1, numbers)
		quick_sort(part_idx+1, end, numbers)

print("[ 10,7,8,9,1,5 ] -> ", do_quick_sort([ 10,7,8,9,1,5 ]))
print("[ 1,8,3,9,4,5,7 ] -> ", do_quick_sort([ 1,8,3,9,4,5,7 ]))
print("[ 7,2,8,4,3,6,9 ] -> ", do_quick_sort([ 7,2,8,4,3,6,9 ]))