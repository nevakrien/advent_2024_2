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
        var input_a,input_b int
        n,err := fmt.Fscanf(reader,"%d %d",&input_a,&input_b)
        if n==0 {
            break
        }
        if(n!=2||err!=nil){
            fmt.Println("bad input:",err)
            os.Exit(1)
        }

        diff:=input_b-input_a
        
        if (input_a == input_b || absInt(diff)<1 || absInt(diff)>3 ){
            reader.ReadLine()
            continue
        }

        var bad = false
        var prev_input = input_b
        //read inputs
        for {
            var input int
            n,err := fmt.Fscanf(reader,"%d",&input)
            if n==0 {
                break
            }
            if(n!=1||err!=nil){
                fmt.Println("bad input:",err)
                os.Exit(1)
            }

            diff2 := input-prev_input
            pos_cond := (diff2>0 && diff<0) || (diff2<0 && diff>0)
            size_cond := absInt(diff2)<1 || absInt(diff2)>3
            
            if (diff2 == 0 || pos_cond || size_cond ) {
                reader.ReadLine()
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