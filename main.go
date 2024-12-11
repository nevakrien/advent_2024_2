package main

import (
    "fmt"
    "os"
    "bufio"
)

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

        var base bool
        var increasing *bool

        if input_a==input_b {
            // fmt.Println("total_count ",total_count)
            // fmt.Println("unexpected a == b == ",input_a)
            // os.Exit(1)
            increasing = nil
        }else{
            base = (input_a<input_b)
            increasing = &base
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

            if(increasing!=nil && (prev_input<input) != *increasing){
                //bad flush number
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