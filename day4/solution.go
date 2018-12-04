package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "regexp"
    "strconv"
    "sort"
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

func parse_line(line string) (int, int, int, int, int, int, string) {
    r, _ := regexp.Compile(`\[(\d\d\d\d)-(\d\d)-(\d\d) (..):(..)\] (.*)`)

    res := r.FindStringSubmatch(line)

    year, _ := strconv.Atoi(res[1])
    month, _ := strconv.Atoi(res[2])
    day, _ := strconv.Atoi(res[3])
    hours, _ := strconv.Atoi(res[4])
    minutes, _ := strconv.Atoi(res[5])
    remaining := res[6]

    guard_id := 0
    phase := ""
    if remaining == "falls asleep" {
        phase = "sleeps"
    } else if remaining == "wakes up"{
        phase = "wakes"
    } else {
        r, _ = regexp.Compile(`Guard #(\d*) begins shift`)
        res = r.FindStringSubmatch(remaining)
        guard_id, _ = strconv.Atoi(res[1])
        phase = "starts"
    }

    return year, month, day, hours, minutes, guard_id, phase
}

type Shift struct {
    year int
    month int
    day int
    hour int
    minute int
    guard_id int
    phase string
}

func get_shift(line string) Shift {
    year, month, day, hours, minutes, guard_id, phase := parse_line(line)
    shift := Shift{year, month, day, hours, minutes, guard_id, phase}
    return shift
}

func get_max_value(m map[int]int) (int, int){
    //get key for which value is maximum
    //return both key and max_value

    var max_value int = 0
    var key_at_max_value int = 0

    for key, val := range m {
        if val > max_value {
            max_value = val
            key_at_max_value = key
        }
    }

    return max_value, key_at_max_value
}

func get_sorted_shifts(input []string) []Shift{
    shifts := []Shift{}

    for _, line := range input {
        shift := get_shift(line)
        shifts = append(shifts, shift)   
    }

    sort.Slice(shifts, func(i, j int) bool {

        smaller := false
        if shifts[i].year < shifts[j].year {
            smaller = true
        } else if shifts[i].year > shifts[j].year {
            smaller = false
        } else if shifts[i].month < shifts[j].month {
            smaller = true
        } else if shifts[i].month > shifts[j].month {
            smaller = false
        } else if shifts[i].day < shifts[j].day {
            smaller = true
        } else if shifts[i].day > shifts[j].day {
            smaller = false
        } else if shifts[i].hour < shifts[j].hour {
            smaller = true
        } else if shifts[i].hour > shifts[j].hour {
            smaller = false
        } else if shifts[i].minute <= shifts[j].minute {
            smaller = true
        } else { // if shifts[i].minute > shifts[j]
            smaller = false
        }

        return smaller
    })

    return shifts
}

func init_or_add(m map[int]int, key int, value int) map[int]int {
    // initialise with value when key not present
    // add to value when key present

    if _, ok := m[key]; ok {
        m[key] = m[key] + value
    } else {
        m[key] = value
    }

    return m
}


func part1_and_part2(){
    input := get_input("input.txt")
    shifts := get_sorted_shifts(input)

    guard_total_minutes := map[int]int{}
    var guard_sleep_minutes = map[int]map[int]int{}

    previous_shift := shifts[0]
    shift := shifts[1]
    guard_id := previous_shift.guard_id

    for i:=1; i < len(shifts); i++ {
        shift = shifts[i]

        if shift.guard_id != 0 {
            guard_id = shift.guard_id
        } else if shift.phase == "wakes" && previous_shift.phase == "sleeps"{
            minutes := shift.minute - previous_shift.minute

            guard_total_minutes = init_or_add(guard_total_minutes, guard_id, minutes)

            for minute := previous_shift.minute; minute < shift.minute; minute ++ {
                if _, ok := guard_sleep_minutes[guard_id]; ok {
                    guard_sleep_minutes[guard_id] = init_or_add(guard_sleep_minutes[guard_id], minute, 1)
                } else {
                    guard_sleep_minutes[guard_id] = map[int]int{}
                    guard_sleep_minutes[guard_id][minute] = 1
                }
            }

        }

        previous_shift = shift
    }

    _, max_sleep_guard_id := get_max_value(guard_total_minutes)

    max_occurrences, max_occurences_minute := get_max_value(guard_sleep_minutes[max_sleep_guard_id])

    fmt.Println("Solution part 1:", max_sleep_guard_id * max_occurences_minute)

    max_occurrences = 0
    max_guard_id := 0
    max_minute := 0

    for guard_id, minute_occurrences := range guard_sleep_minutes {
        occurrences, minute := get_max_value(minute_occurrences)

        if occurrences > max_occurrences {
            max_occurrences = occurrences
            max_guard_id = guard_id
            max_minute = minute
        }

    }

    fmt.Println("Solution part 2:", max_guard_id * max_minute)

}

func main(){
    part1_and_part2()
}