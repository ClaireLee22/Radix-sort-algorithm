BASE = 10 # 0 - 9 digits

def radixSort(array):
    if len(array) == 0:
        return array

    maxValue = max(array)

    placeValue = 1
    while maxValue//placeValue > 0:
        countingSort(array, placeValue)
        print("place value: ", placeValue, "array: ",  array)
        placeValue *= BASE
    
    return array

# use counting sort as subroutine
def countingSort(array, placeValue):
    minValue = getMinValue(array, placeValue)
    offset = 0 # deal with negative values
    if minValue < 0:
        offset = 0 - minValue
    counts = [0]*(BASE + offset) # minValue-9 digits
    sortedArray = [0]*len(array)

    for i in range(len(array)):
        countsIdx = getDigit(array, i, placeValue) + offset
        counts[countsIdx] += 1

    for i in range(1, len(counts)):
        counts[i] += counts[i-1]

    for i in range(len(array)-1, -1, -1):
        countsIdx = getDigit(array, i, placeValue) + offset
        position = counts[countsIdx]
        sortedIdx = position - 1
        sortedArray[sortedIdx] = array[i]
        counts[countsIdx] -= 1

    # copy values to original array
    for i in range(len(array)):
        array[i] = sortedArray[i]

# get minimum value of each digit group
def getMinValue(array, placeValue):
    minValue = getDigit(array, 0, placeValue)
    for i in range(1, len(array)):
        element = getDigit(array, i, placeValue)
        if element < minValue:
            minValue = element
    return minValue

# ex. number: 2022, place value: 100 => digit = 0
# ex. number: -1223, place value: 100 => digit = -2
def getDigit(array, idx, placeValue):
    element = array[idx]
    digit = abs(element) // placeValue % BASE
    return digit if element > 0 else -digit

if __name__ == "__main__":
    arrays = [[2022, 5, 398, 159, 16, 1223, 7], [2022, -5, 398, 159, 16, -1223, 7]]
    for array in arrays:
        print("unsorted array: ", array)
        print("sorted array: ", radixSort(array))
        print("\n")

"""
output:
('unsorted array: ', [2022, 5, 398, 159, 16, 1223, 7])
('place value: ', 1, 'array: ', [2022, 1223, 5, 16, 7, 398, 159])
('place value: ', 10, 'array: ', [5, 7, 16, 2022, 1223, 159, 398])
('place value: ', 100, 'array: ', [5, 7, 16, 2022, 159, 1223, 398])
('place value: ', 1000, 'array: ', [5, 7, 16, 159, 398, 1223, 2022])
('sorted array: ', [5, 7, 16, 159, 398, 1223, 2022])


('unsorted array: ', [2022, -5, 398, 159, 16, -1223, 7])
('place value: ', 1, 'array: ', [-5, -1223, 2022, 16, 7, 398, 159])
('place value: ', 10, 'array: ', [-1223, -5, 7, 16, 2022, 159, 398])
('place value: ', 100, 'array: ', [-1223, -5, 7, 16, 2022, 159, 398])
('place value: ', 1000, 'array: ', [-1223, -5, 7, 16, 159, 398, 2022])
('sorted array: ', [-1223, -5, 7, 16, 159, 398, 2022])

"""