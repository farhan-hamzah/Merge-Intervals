package main
import (
    "fmt"
    "sort"
)

func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }

    // Urutkan berdasarkan elemen pertama dari setiap interval
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    var result [][]int
    current := intervals[0]

    for i := 1; i < len(intervals); i++ {
        next := intervals[i]
        if current[1] >= next[0] {
            // Jika tumpang tindih, gabungkan interval
            if current[1] < next[1] {
                current[1] = next[1]
            }
        } else {
            // Jika tidak tumpang tindih, simpan interval sebelumnya
            result = append(result, current)
            current = next
        }
    }

    // Tambahkan interval terakhir
    result = append(result, current)

    return result
}
func main() {
    intervals := [][]int{
        {1, 3},
        {2, 6},
        {8, 10},
        {15, 18},
    }

    merged := merge(intervals)

    fmt.Println("Hasil penggabungan interval:")
    for _, interval := range merged {
        fmt.Printf("[%d, %d]\n", interval[0], interval[1])
    }
}
