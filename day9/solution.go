package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "strconv"
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

func parse_line(line string) (int, int) {
    r, _ := regexp.Compile(`(\d*) players; last marble is worth (\d*) points`)

    res := r.FindStringSubmatch(line)

    number_of_players, _ := strconv.Atoi(res[1])
    max_points, _ := strconv.Atoi(res[2])

    return number_of_players, max_points

}

type Node struct{
    value int
    previous_node *Node
    next_node *Node
}

func insert_node(left_node *Node, right_node *Node, value int) *Node{
    new_node := &Node{value, left_node, right_node}
    left_node.next_node = new_node
    right_node.previous_node = new_node
    return new_node
}

func remove_node(node *Node) (int, *Node) {
    left_node := node.previous_node
    right_node := node.next_node
    left_node.next_node = right_node
    right_node.previous_node = left_node
    return node.value, right_node
}

func calculate_max_score(number_of_players int, max_points int) int{

    var n1 *Node
    var n2 *Node
    n1 = &Node{0, n2, n2}
    n2 = &Node{1, n1, n1}
    n1.previous_node = n2
    n1.next_node = n2

    current_node := n2

    marble := 2
    var player int
    var removed_value int
    var node_to_be_removed *Node
    var left_node *Node
    var right_node *Node

    scores := map[int]int{}
    for player_id:=1; player_id<=number_of_players; player_id++{
        scores[player_id] = 0
    }

    for ; marble <= max_points; marble++ {
        player = marble % number_of_players
        if player == 0 {
            player = number_of_players
        }

        if marble % 23 == 0 {
            scores[player] += marble
            node_to_be_removed = current_node.previous_node.previous_node.previous_node.previous_node.previous_node.previous_node.previous_node
            removed_value, current_node = remove_node(node_to_be_removed)
            scores[player] += removed_value
        } else {
            left_node = current_node.next_node
            right_node = current_node.next_node.next_node
            current_node = insert_node(left_node, right_node, marble)
        }
    }

    max_score := 0
    for _, score := range scores {
        if score > max_score {
            max_score = score
        }
    }

    return max_score
}

func part1(input []string){
    number_of_players, max_points := parse_line(input[0])
    max_score := calculate_max_score(number_of_players, max_points)
    fmt.Println("Solution part 1:", max_score)
}

func part2(input []string){
    number_of_players, max_points := parse_line(input[0])
    max_score := calculate_max_score(number_of_players, max_points * 100)
    fmt.Println("Solution part 2:", max_score)

}

func main(){
    input := get_input("input.txt")
    part1(input)
    part2(input)
}