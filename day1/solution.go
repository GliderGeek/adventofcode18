package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func get_input(file_name string) []string {
    file, err := os.Open(file_name)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    lines := []string{}

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return lines
}

func part1(){
    lines := get_input("input.txt")

    x := 0

    for _, line := range lines {
        i, _ := strconv.Atoi(line)
        x = x + i
    }

    fmt.Print("Result part 1 is: ")
    fmt.Println(x)
}

func part2(){
    
    lines := get_input("input.txt")

    x := 0
    vals := map[int]bool{
        0: true,
    }

    loops := 1000
    result_found := false

    for i := 0; i < loops; i++ {

        for _, line := range lines {

            num, _ := strconv.Atoi(line)            
            x = x + num

            if vals[x] {
                fmt.Print("Result part 2 is: ")
                fmt.Println(x)
                result_found = true
                break
            } else {
                vals[x] = true
            }
        }

        if result_found {
            break
        }
    }
}

func main() {
    part1()
    part2()
}



