BASE = 10

def radixSort(array):
    if len(array) == 1:
        return array
    
    maxValue = max(array)
    # leftmost place value
    placeValue = BASE**(getNumberOfDigits(maxValue) - 1)

    bucketSort(array, placeValue)
       
    return array


def getNumberOfDigits(num):
    numberOfdigits = 0
    while num > 0:
        numberOfdigits += 1
        num /= BASE
    return numberOfdigits

# recursion
def bucketSort(array, placeValue):
    if placeValue < 1:
        return array
    
    if len(array) <= 1:
        return array

    buckets = []
    minValue = getMinValue(array, placeValue)

    offset = 0
    if minValue < 0:
        offset = 0 - minValue
    for i in range(BASE+offset):
        buckets.append([])

    for idx, num in enumerate(array):
        bucketIdx = getDigit(array, idx, placeValue) + offset
        buckets[bucketIdx].append(num)
    print("place value: ", placeValue, "bucket: ",  buckets)
    

    sortedIdx = 0
    for bucket in buckets:
        sortedBucket = bucketSort(bucket, placeValue/BASE)
        for num in sortedBucket:
            array[sortedIdx] = num
            sortedIdx += 1
    print("place value: ", placeValue, "array: ",  array)

    return array

def getMinValue(array, placeValue):
    minValue = getDigit(array, 0, placeValue)
    for i in range(1, len(array)):
        element = getDigit(array, i, placeValue)
        if element < minValue:
            minValue = element
    return minValue

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
'unsorted array: ', [2022, 5, 398, 159, 16, 1223, 7])
('place value: ', 1000, 'bucket: ', [[5, 398, 159, 16, 7], [1223], [2022], [], [], [], [], [], [], []])
('place value: ', 100, 'bucket: ', [[5, 16, 7], [159], [], [398], [], [], [], [], [], []])
('place value: ', 10, 'bucket: ', [[5, 7], [16], [], [], [], [], [], [], [], []])
('place value: ', 1, 'bucket: ', [[], [], [], [], [], [5], [], [7], [], []])
('place value: ', 1, 'array: ', [5, 7])
('place value: ', 10, 'array: ', [5, 7, 16])
('place value: ', 100, 'array: ', [5, 7, 16, 159, 398])
('place value: ', 1000, 'array: ', [5, 7, 16, 159, 398, 1223, 2022])
('sorted array: ', [5, 7, 16, 159, 398, 1223, 2022])


('unsorted array: ', [2022, -5, 398, 159, 16, -1223, 7])
('place value: ', 1000, 'bucket: ', [[-1223], [-5, 398, 159, 16, 7], [], [2022], [], [], [], [], [], [], []])
('place value: ', 100, 'bucket: ', [[-5, 16, 7], [159], [], [398], [], [], [], [], [], []])
('place value: ', 10, 'bucket: ', [[-5, 7], [16], [], [], [], [], [], [], [], []])
('place value: ', 1, 'bucket: ', [[-5], [], [], [], [], [], [], [], [], [], [], [], [7], [], []])
('place value: ', 1, 'array: ', [-5, 7])
('place value: ', 10, 'array: ', [-5, 7, 16])
('place value: ', 100, 'array: ', [-5, 7, 16, 159, 398])
('place value: ', 1000, 'array: ', [-1223, -5, 7, 16, 159, 398, 2022])
('sorted array: ', [-1223, -5, 7, 16, 159, 398, 2022])
"""
