def do_quick_sort(numbers):
	quick_sort(0, len(numbers)-1, numbers)
	return numbers


# partition with pivot at start
# find number greater than pivot from left
# find number smaller than pivot from right
# if start and end have not crossed each other, swap the numbers on start and end
def partition_pivot_at_start(start, end, numbers):
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
	# means end idx_is pivot_idx
	return end

# partition with pivot at end
# https://www.youtube.com/watch?v=MZaf_9IZCrc
def partition_pivot_at_end(start, end, numbers):
	i = (start-1)
	pivot = numbers[end]

	for j in range(start, end):
		if numbers[j] <= pivot:
			i = i+1 # increment index of smaller element
			numbers[i], numbers[j] = numbers[j], numbers[i] # swap

	numbers[i+1], numbers[end] = numbers[end], numbers[i+1]
	return (i+1)

def quick_sort(start, end, numbers):
	if start < end:
		part_idx = partition_pivot_at_start(start, end, numbers) 
		quick_sort(start, part_idx-1, numbers)
		quick_sort(part_idx+1, end, numbers)

print("[ 10,7,8,9,1,5 ] -> ", do_quick_sort([ 10,7,8,9,1,5 ]))
print("[ 1,8,3,9,4,5,7 ] -> ", do_quick_sort([ 1,8,3,9,4,5,7 ]))
print("[ 7,2,8,4,3,6,9 ] -> ", do_quick_sort([ 7,2,8,4,3,6,9 ]))
print("[ 2,6,5,0,8,7,1,3] -> ", do_quick_sort([ 2,6,5,0,8,7,1,3 ]))
print("[ 2,1,0] -> ", do_quick_sort([ 2,1,0 ]))
print("[ 8,7,6,5 ] -> ", do_quick_sort([ 7,6,8,5 ]))