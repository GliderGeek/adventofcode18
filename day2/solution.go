package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
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

    contains_double := 0
    contains_triple := 0

    for _, line := range lines{

        present_1time := map[rune]bool{}
        present_2times := map[rune]bool{}
        present_3times := map[rune]bool{}
        present_more3times := map[rune]bool{}

        // runes := []rune(line)
        for _, rune := range []rune(line){

            if present_more3times[rune] {
                continue
            } else if present_3times[rune] {
                present_more3times[rune] = true
                delete(present_3times, rune)
            } else if present_2times[rune] {
                present_3times[rune] = true
                delete(present_2times, rune)
            } else if present_1time[rune] {
                present_2times[rune] = true
                delete(present_1time, rune)
            } else {
                present_1time[rune] = true
            }
        }

        if len(present_2times) > 0{
            contains_double ++
        }
        if len(present_3times) > 0{
            contains_triple ++
        }
    }

    total := contains_double * contains_triple

    fmt.Println("Solution part 1:", total)
}

func part2(){
    lines := get_input("input.txt")

    difference_indexes := map[int]bool{}
    found_it := false
    for i, line := range lines {

        for j, line2 := range lines {

            if i == j {continue}

            difference_indexes = map[int]bool{}
            for k := 0; k < len(line) - 1; k++ {
                if line[k] != line2[k] {
                    difference_indexes[k] = true
                }
            }

            if len(difference_indexes) == 1 {
                found_it = true

                equal := []byte{}
                for l := 0; l < len(line); l++ {

                    if difference_indexes[l]{
                        continue
                    } else {
                        equal = append(equal, line[l])
                    }
                }

                fmt.Print("Solution part 2: ")
                for l:=0; l<len(equal);l++ {
                    fmt.Printf("%c", equal[l])
                }
                fmt.Println("")
                break
            }
        }
        if found_it{break}
    }
}

func main() {
    part1()
    part2()
}