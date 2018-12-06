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

func abs(i int) int {
    if i < 0{
        return i * -1
    } else {
        return i
    }
}

func parse_line(line string) (int, int) {
    r, _ := regexp.Compile(`(.*), (.*)`)

    res := r.FindStringSubmatch(line)

    x, _ := strconv.Atoi(res[1])
    y, _ := strconv.Atoi(res[2])

    return x, y
}

type Coordinate struct{
    index int
    x int
    y int
}

func get_coordinates(input []string) []Coordinate{
    coordinates := []Coordinate{}
    for i, line := range input {
        x, y := parse_line(line)
        coordinates = append(coordinates, Coordinate{i, x, y})
    }

    return coordinates
}

func get_max_values(coordinates []Coordinate) (int, int) {
    max_x := 0
    max_y := 0

    for _, coordinate := range coordinates {
        if coordinate.x > max_x {
            max_x = coordinate.x
        }

        if coordinate.y > max_y {
            max_y = coordinate.y
        }
    }

    return max_x, max_y
}

func find_closest_index(x int, y int, coordinates []Coordinate) int {
    //Return -1 if multiple

    min_distance := -1  //to start
    index := 0
    distance := 0

    for _, coordinate := range coordinates {
        distance = abs(coordinate.x - x) + abs(coordinate.y - y)

        if min_distance == -1 || distance < min_distance{
            min_distance = distance
            index = coordinate.index
        } else if distance == min_distance{
            index = -1
        }

    }

    return index
}

func find_sum_of_distances(x int, y int, coordinates []Coordinate) int {

    sum := 0
    for _, coordinate := range coordinates {
        sum += abs(coordinate.x - x) + abs(coordinate.y - y)
    }

    return sum
}

func part1(input []string){
    coordinates := get_coordinates(input)

    x_max, y_max := get_max_values(coordinates)

    surface_by_index := map[int]int{}  //-1 when infinite

    for x:=0; x<=x_max; x++ {
        for y:=0; y<=y_max; y++{

            closest_index := find_closest_index(x, y, coordinates)

            if closest_index != -1 {
                if x == 0 || x == x_max || y == 0 || y == y_max {
                    surface_by_index[closest_index] = -1
                } else {
                    if val, ok := surface_by_index[closest_index]; ok {
                        if val != -1{
                            surface_by_index[closest_index] = surface_by_index[closest_index] + 1
                        }
                    } else {
                        surface_by_index[closest_index] = 1
                    }

                }
            } 
        }
    }

    total_area := 0
    for _, val := range surface_by_index{
        if val > total_area{
            total_area = val
        }
    }

    fmt.Println("Solution part 1:", total_area)
}

func part2(input []string, minimum_distance int){
    coordinates := get_coordinates(input)

    x_max, y_max := get_max_values(coordinates)

    total_area := 0
    for x:=0; x<=x_max; x++ {
        for y:=0; y<=y_max; y++{

            if find_sum_of_distances(x, y, coordinates) < minimum_distance {
                total_area += 1
            }
        }
    }

    fmt.Println("Solution part 2:", total_area)
}

func main(){
    input := get_input("input.txt")
    minimum_distance := 10000
    part1(input)
    part2(input, minimum_distance)
}