package quicksort

func QuickSort(elements []int) {
	sort(elements, 0, len(elements)-1)
}

func sort(elements []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(elements, lo, hi)
	sort(elements, lo, p-1)
	sort(elements, p+1, hi)
}

func partition(elements []int, lo, hi int) int {
	i := lo
	j := hi + 1
	pivot := elements[lo]
	for {
		for {
			i++
			if elements[i] < pivot {
				if i == hi {
					break
				}
			} else {
				break
			}
		}

		for {
			j--
			if elements[j] > pivot {
				if j == lo {
					break
				}
			} else {
				break
			}
		}

		if i >= j {
			break
		}
		elements[i], elements[j] = elements[j], elements[i]
	}
	elements[lo], elements[j] = elements[j], elements[lo]
	return j
}

func Quick3way(elements []int) {
	sort3way(elements, 0, len(elements)-1)
}

func sort3way(elements []int, lo, hi int) {
	if lo >= hi {
		return
	}
	lt := lo
	gt := hi
	pivot := elements[lo]
	i := lo + 1
	for i <= gt {
		if elements[i] < pivot {
			elements[lt], elements[i] = elements[i], elements[lt]
			lt++
			i++
		} else if elements[i] > pivot {
			elements[i], elements[gt] = elements[gt], elements[i]
			gt--
		} else {
			i++
		}
	}
	sort3way(elements, lo, lt-1)
	sort3way(elements, gt+1, hi)
}