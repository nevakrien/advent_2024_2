package main

import (
    "fmt"
    "os"
    "bufio"
)


func absInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func gatherLine(reader *bufio.Reader ) []int {
    numbers := make([]int, 0)
    var input int
    for  {
        n,err := fmt.Fscanf(reader,"%d",&input)
        if (err != nil || n==0) {
            return numbers
        }
        numbers = append(numbers,input)
    }
}

func main() {
    // Open the file
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        os.Exit(1)

    }
    defer file.Close() // Ensure the file is closed when the function exits
    reader := bufio.NewReader(file)
    
    var safe_count = 0
    var total_count = 0


    for ;; total_count++ {
        fmt.Println("start again")
        arr := gatherLine(reader)
        fmt.Println("arr = ",arr)

        if len(arr) == 0 {
            // fmt.Println("empty")
            break
        }
        var input_a,input_b int
        
        input_a = arr[0]
        input_b = arr[1]
        diff:=input_b-input_a
        
        if (input_a == input_b || absInt(diff)<1 || absInt(diff)>3 ){
            continue
        }

        var bad = false
        var prev_input = input_b
        //read inputs
        for i:=2;i<len(arr);i++ {
            input:=arr[i]

            diff2 := input-prev_input
            pos_cond := (diff2>0 && diff<0) || (diff2<0 && diff>0)
            size_cond := absInt(diff2)<1 || absInt(diff2)>3
            
            if (diff2 == 0 || pos_cond || size_cond ) {
                bad = true
                break
            }

            prev_input = input

        }
        if !bad {
            safe_count++
        }
    }
    
    fmt.Println("total_count:",total_count,"safe_count:",safe_count)
}