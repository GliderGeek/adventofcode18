package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "regexp"
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

func reduce_polymer (polymer string, combinations []string) string {

    old_length := len(polymer)
    new_length := old_length - 1  // just to start

    r, _ := regexp.Compile("")

    for new_length < old_length{

        old_length = new_length

        for _, combination := range combinations{
            r, _ = regexp.Compile(combination)
            polymer = r.ReplaceAllString(polymer, "")
        }

        new_length = len(polymer)
    }

    return polymer
}


func get_letter_combinations() []string {
    // Aa, aA, bB, Bb etc... (duplicated with different order)

    combinations := []string{}
    small_i := 0
    capital_i := 0

    for i := 0; i< 26; i++ {
        small_i = int("a"[0]) + i
        capital_i = int("A"[0]) + i
        combinations = append(combinations, string(small_i) + string(capital_i))
        combinations = append(combinations, string(capital_i) + string(small_i))
    }
    return combinations
}

func get_letter_pairs() []string {
    // aA, bB, cC (no duplicates)

    pairs := []string{}
    small_i := 0
    capital_i := 0

    for i := 0; i< 26; i++ {
        small_i = int("a"[0]) + i
        capital_i = int("A"[0]) + i
        pairs = append(pairs, string(small_i) + string(capital_i))
    }
    return pairs
}

func part1(polymer string){
    combinations := get_letter_combinations()
    polymer = reduce_polymer(polymer, combinations)
    fmt.Println("Solution part 1:", len(polymer))
}

func part2(polymer string) {

    pairs := get_letter_pairs()
    combinations := get_letter_combinations()

    new_polymer := polymer

    min_length := -1

    r, _ := regexp.Compile("")
    for _, pair := range pairs{

        r, _ = regexp.Compile("[" + pair + "]")
        new_polymer = r.ReplaceAllString(polymer, "")

        new_polymer = reduce_polymer(new_polymer, combinations)

        if len(new_polymer) < min_length || min_length == -1 {
            min_length = len(new_polymer)
        }
    }

    fmt.Println("Solution part 2:", min_length)
}

func main(){
    input := get_input("input.txt")
    polymer := input[0]
    part1(polymer)
    part2(polymer)
}
