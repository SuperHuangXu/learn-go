package arrays

func Sum(num []int) int {
  res := 0
  for _, n := range num {
    res += n
  }
  return res
}

func SumAll(numbersToSum ...[]int) []int {
  lengthOfNumbers := len(numbersToSum)
  sums := make([]int, lengthOfNumbers)

  for i, numbers := range numbersToSum {
    sums[i] = Sum(numbers)
  }
  return sums
}

func SumAllTails(numbers ...[]int) []int {
  var sums []int
  for _, numbers := range numbers {
    if len(numbers) == 0 {
      sums = append(sums, 0)
    } else {
      tail := numbers[1:]
      sums = append(sums, Sum(tail))
    }
  }
  return sums
}
