package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "regexp"
    "sort"
)


type Node struct{
    name string
    from_nodes map[string]*Node
    to_nodes map[string]*Node
}

func add_edge(start_node *Node, end_node *Node) {
    start_node.to_nodes[end_node.name] = end_node
    end_node.from_nodes[start_node.name] = start_node
}

func newNode(name string) *Node {
    return &Node{name, make(map[string]*Node), make(map[string]*Node)}
}

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

func addPairOfNodes(nodes map[string]*Node, node1_name string, node2_name string) map[string]*Node{
    node1, node2 := newNode(node1_name), newNode(node2_name)

    if _, ok := nodes[node1.name]; ok {
        node1 = nodes[node1.name]
    } else {
        nodes[node1.name] = node1
    }

    if _, ok := nodes[node2.name]; ok {
        node2 = nodes[node2.name]
    } else {
        nodes[node2.name] = node2
    }
    
    add_edge(node1, node2)

    return nodes
}

func remove_top_node(nodes map[string]*Node, node *Node){
    // remove all references to this node
    for _, to_node := range node.to_nodes{
        delete(to_node.from_nodes, node.name)
    }

    //remove node itself
    delete(nodes, node.name)
}

func get_top_node_names(nodes map[string]*Node) []string {
    // sorted in alphabetic order

    top_node_names := []string{}

    if len(nodes) == 1 {
        k := ""
        for key, _ := range nodes{
            k = key
        }

        top_node_names = append(top_node_names, nodes[k].name)
        return top_node_names
    } else {
        for _, node := range nodes {
            if len(node.from_nodes) == 0{
                top_node_names = append(top_node_names, node.name)
            }
        }
    }

    sort.Slice(top_node_names, func(i, j int) bool {
        return int(top_node_names[i][0]) < int(top_node_names[j][0])
    })

    return top_node_names
}


func get_node_name_for_removal(nodes map[string]*Node) string {
    top_node_names := get_top_node_names(nodes)
    return top_node_names[0]
}


func get_nodes(input []string) map[string]*Node {
    nodes := map[string]*Node{}

    r, _ := regexp.Compile(`Step (.) must be finished before step (.) can begin.`)
    res := []string{}    

    for _, line := range input {
        res = r.FindStringSubmatch(line)
        node1 := res[1]
        node2 := res[2]

        nodes = addPairOfNodes(nodes, node1, node2)
    }

    return nodes
}

func part1(input []string){

    nodes := get_nodes(input)

    execution_order := []string{}

    fmt.Print("Solution part1: ")
    for len(nodes) > 0 {
        node_name_for_removal := get_node_name_for_removal(nodes)

        execution_order = append(execution_order, node_name_for_removal)

        fmt.Print(node_name_for_removal)
        remove_top_node(nodes, nodes[node_name_for_removal])
    }

    fmt.Println("")
}

type Worker struct {
    time_left int
    current_node_name string
}

func part2(input []string, number_of_workers int, additional_time int){

    nodes := get_nodes(input)

    top_node_names := get_top_node_names(nodes)
    taken_nodes := map[string]bool{}

    workers := []*Worker{}

    for i:=0; i<number_of_workers; i++ {
        workers = append(workers, &Worker{0, ""})
    }

    total_time_left := 1
    time := 0

    for ;total_time_left > 0; time++ {

        total_time_left = 0
        for _, worker := range workers {

            if worker.time_left != 0{
                worker.time_left = worker.time_left - 1

                if worker.time_left == 0{ //finished node
                    remove_top_node(nodes, nodes[worker.current_node_name])
                    delete(taken_nodes, worker.current_node_name)
                    worker.current_node_name = ""
                    top_node_names = get_top_node_names(nodes)
                }
            } 

            if worker.time_left == 0 && len(top_node_names) != 0 {
                for _, name := range top_node_names {
                    if !taken_nodes[name] {
                        worker.time_left = additional_time + int(name[0]) - int("A"[0]) + 1
                        taken_nodes[name] = true
                        worker.current_node_name = name
                        break
                    }
                }
            }

            total_time_left += worker.time_left
        }
    }

    fmt.Println("Solution part 2:", time-1)

}

func main(){
    input := get_input("input.txt")
    number_of_workers := 5
    additional_time := 60

    part1(input)
    part2(input, number_of_workers, additional_time)
}