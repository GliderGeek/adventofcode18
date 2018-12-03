package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "regexp"
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

type piece struct {
    id string
    x_min int
    x_max int
    y_min int
    y_max int
}

func get_piece(line string) piece {
    r, _ := regexp.Compile("#(.*) @ (.*),(.*): (.*)x(.*)")

    res := r.FindStringSubmatch(line)
    _, id, x1, y1, width, height := res[0], res[1], res[2], res[3], res[4], res[5]

    x_min, _ := strconv.Atoi(x1)
    width_i, _ := strconv.Atoi(width) 
    height_i, _ := strconv.Atoi(height)
    y_min, _ := strconv.Atoi(y1)

    x_max := x_min + width_i
    y_max := y_min + height_i

    return piece{id, x_min, x_max, y_min, y_max}
}



func part1_and_part2(){
    input := get_input("input.txt")

    var coordinates = map[int]map[int]bool{}
    pieces := map[piece]bool{}

    total := 0
    for _, line := range input {
        piece := get_piece(line)
        pieces[piece] = true

        for x:=piece.x_min; x<piece.x_max; x++{
            for y:=piece.y_min; y<piece.y_max; y++{

                if _, ok := coordinates[x]; ok {
                    if val, ok := coordinates[x][y]; ok {
                        if val{
                            total = total + 1
                            coordinates[x][y] = false
                        }
                    } else {
                        coordinates[x][y] = true
                    }
                } else{
                    coordinates[x] = map[int]bool{}
                    coordinates[x][y] = true
                }
            }
        }
    }

    fmt.Println("Solution part 1:", total)

    var found_the_one bool
    for piece, _ := range pieces{
        found_the_one = true
        for x:=piece.x_min; x<piece.x_max; x++{
            for y:=piece.y_min; y<piece.y_max; y++{
                if coordinates[x][y] == false {
                    found_the_one = false
                    break
                }
            }
            if found_the_one == false{break}
        }

        if found_the_one {
            fmt.Println("Solution part 2:", piece.id)
        }
    }
}

func main(){
    part1_and_part2()
}