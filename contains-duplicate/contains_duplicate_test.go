package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        expected bool
    }{
        {
            name:     "No duplicates",
            nums:     []int{1, 2, 3, 4},
            expected: false,
        },
        {
            name:     "Duplicates present",
            nums:     []int{1, 2, 3, 1},
            expected: true,
        },
        {
            name:     "Empty array",
            nums:     []int{},
            expected: false,
        },
        {
            name:     "Single element",
            nums:     []int{1},
            expected: false,
        },
        {
            name:     "All elements the same",
            nums:     []int{5, 5, 5, 5},
            expected: true,
        },
        {
            name:     "Large input with duplicates",
            nums:     []int{100, 200, 300, 400, 100},
            expected: true,
        },
        {
            name:     "Large input without duplicates",
            nums:     []int{100, 200, 300, 400, 500},
            expected: false,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := containsDuplicate(test.nums)
            if result != test.expected {
                t.Errorf("containsDuplicate(%v) = %v; want %v", test.nums, result, test.expected)
            }
        })
    }
}

