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

func checkLine(arr []int,ignore int) int {        
    var input_a,input_b,aidx,bidx int

    if(ignore == 0){
        input_a = arr[1]
        input_b = arr[2]
        aidx = 1
        bidx = 2
    } else if(ignore==1) {
        input_a = arr[0]
        input_b = arr[2]
        aidx = 0
        bidx = 2
    } else{
        input_a = arr[0]
        input_b = arr[1]
        aidx = 0
        bidx = 1
    }
    
    // fmt.Println("a =",input_a,"b =",input_b)

    
    diff:=input_b-input_a
    
    if (input_a == input_b || absInt(diff)<1 || absInt(diff)>3 ){
        // fmt.Println("first comperison failed")
        return aidx
    }

    var prev_input = input_b
    for i:=bidx+1;i<len(arr);i++ {
        if i==ignore {
            // fmt.Println("skiping value",arr[i])
            continue
        }
        input:=arr[i]

        diff2 := input-prev_input
        pos_cond := (diff2>0 && diff<0) || (diff2<0 && diff>0)
        size_cond := absInt(diff2)<1 || absInt(diff2)>3
        
        if (diff2 == 0 || pos_cond || size_cond ) {
            // fmt.Println("failed at comperison",prev_input,"?",input)
            return i-1
        }

        prev_input = input

    }

    return -1
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
        
        // fmt.Println("start again")
        arr := gatherLine(reader)
        fmt.Println("arr = ",arr)

        if len(arr) == 0 {
            // fmt.Println("empty")
            break
        }
        
        // breakPoint := checkLine(arr,-1)
        // if  breakPoint == -1 {
        //     safe_count++
        //     // fmt.Println("pass",total_count,"no fixup")
        //     continue
        // }

        // if checkLine(arr,breakPoint) == -1{
        //     safe_count++
        //     // fmt.Println("pass",total_count)
        //     continue
        // }

        // if checkLine(arr,breakPoint+1) == -1{
        //     safe_count++
        //     // fmt.Println("pass",total_count)
        //     continue
        // }
        for j:=-1;j<len(arr);j++{
            if  checkLine(arr,j) == -1 {
                safe_count++
                // fmt.Println("pass",total_count,"no fixup")
                goto end;
            }
        }


        fmt.Println("defective",total_count)
        end:
    }
    
    fmt.Println("total_count:",total_count,"safe_count:",safe_count)
}