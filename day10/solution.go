package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "strconv"
    "regexp"
    "strings"
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

func parse_line(line string) *Point {
    r, _ := regexp.Compile(`position=<(.*),(.*)> velocity=<(.*),(.*)>`)

    res := r.FindStringSubmatch(line)

    x, _ := strconv.Atoi(strings.TrimSpace(res[1]))
    y, _ := strconv.Atoi(strings.TrimSpace(res[2]))
    vx, _ := strconv.Atoi(strings.TrimSpace(res[3]))
    vy, _ := strconv.Atoi(strings.TrimSpace(res[4]))

    return &Point{x, y, vx, vy}
}

type Point struct{
    x int
    y int
    vx int
    vy int
}

// func print_time_step(points map[*Point]bool) {
//     min_x = 0

// }

func advance_points (points map[*Point]bool) map[*Point]bool {
    for point, _ := range points {
        point.x += point.vx
        point.y += point.vy
    }

    return points
}

func probably_text (points map[*Point]bool, desired_range_y int) (int, int, int, int, bool) {
    max_y := -10000
    min_y := 10000
    min_x := 10000
    max_x := -10000

    for point, _ := range points {

        if point.y > max_y {
            max_y = point.y
        } 

        if point.y < min_y {
            min_y = point.y
        } 

        if point.x > max_x {
            max_x = point.x
        }

        if point.x < min_x {
            min_x = point.x
        }

        if max_y - min_y >= desired_range_y {
            break
        }

    }

    text := (max_y - min_y < desired_range_y)
    return min_x, max_x, min_y, max_y, text
}


func print_points(points map[*Point]bool, min_x int, max_x int, min_y int, max_y int) {
    present := false
    for y:=min_y; y<=max_y; y++{
        for x:=min_x; x<=max_x; x++ {

            present = false
            for point, _ := range points {
                if point.x == x && point.y == y {
                    present = true
                    break
                }
            }

            if present {
                fmt.Print("#")
            } else {
                fmt.Print(" ")
            }

        }
        fmt.Println("")
    }
}

func part1_part2(input []string){

    points := map[*Point]bool{}

    var point *Point
    for _, line := range input{
        point = parse_line(line)
        points[point] = true
    }

    i := 1

    desired_range_y := 10
    max_i := 1000000

    var min_x int
    var max_x int
    var min_y int
    var max_y int
    var text bool

    for ; i< max_i; i++{
        points = advance_points(points)
        min_x, max_x, min_y, max_y, text = probably_text(points, desired_range_y)
        if text {
            break
        }
    }
    
    fmt.Println("Solution part 1:")
    print_points(points, min_x, max_x, min_y, max_y)

    fmt.Println("Solution part 2:", i)
}

func main(){
    input := get_input("input.txt")
    part1_part2(input)
}