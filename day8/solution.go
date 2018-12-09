package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "strings"
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

func get_additional_metadata(number_list []int, left_index int) (int, int) {

    additional_metadata := 0


    number_of_children := number_list[left_index]
    left_index += 1

    number_of_metadata_entries := number_list[left_index]
    left_index += 1

    if number_of_children != 0 {
        child_additional_meta_data := 0
        for child_i := 0; child_i < number_of_children; child_i++{
            child_additional_meta_data, left_index = get_additional_metadata(number_list, left_index)
            additional_metadata += child_additional_meta_data
        }
    }

    end_of_metadata_index := left_index + number_of_metadata_entries
    for ; left_index < end_of_metadata_index; left_index ++ {
        additional_metadata += number_list[left_index]
    }

    return additional_metadata, left_index
}

func get_node(number_list []int, left_index int) (*Node, int) {

    number_of_children := number_list[left_index]
    left_index += 1

    number_of_metadata_entries := number_list[left_index]
    left_index += 1

    to_nodes := []*Node{}
    if number_of_children != 0 {
        var node *Node
        for child_i := 0; child_i < number_of_children; child_i++{   
            node, left_index = get_node(number_list, left_index)
            to_nodes = append(to_nodes, node)
        }
    } 

    metadata := []int{}
    end_of_metadata_index := left_index + number_of_metadata_entries
    for ; left_index < end_of_metadata_index; left_index ++ {
        metadata = append(metadata, number_list[left_index])
    }
    
    return &Node{to_nodes, metadata}, left_index
}

func get_node_value(node *Node) int {
    value := 0
    if len(node.to_nodes) == 0 {
        for _, metadata_value := range node.metadata {
            value += metadata_value
        }
    } else {
        var child_node *Node
        for _, metadata := range node.metadata {
            
            if metadata > 0 && metadata <= len(node.to_nodes) {
                child_node = node.to_nodes[metadata-1]
                value += get_node_value(child_node)
            }

        }
    }
    return value
}

func get_number_list(input string) []int {
    values := strings.Split(input, " ")
    number_list := make([]int, len(values))

    for i, val := range values {
        number, _ := strconv.Atoi(val)
        number_list[i] = number
    }

    return number_list
}

type Node struct{
    to_nodes []*Node
    metadata []int
}

func part1(input []string){
    number_list := get_number_list(input[0])
    metadata_total, _ := get_additional_metadata(number_list, 0)
    fmt.Println("Solution part 1:", metadata_total)
}


func part2(input []string){
    number_list := get_number_list(input[0])
    top_node, _ := get_node(number_list, 0)
    fmt.Println("Solution part 2:", get_node_value(top_node))

}

func main(){
    input := get_input("input.txt")
    part1(input)
    part2(input)
}