package main

import (
    "fmt"
    "strconv"
    "regexp"
    "io/ioutil"
)

func get_input(file_name string) string {
    raw_content, _ := ioutil.ReadFile(file_name)
    return string(raw_content)
}

func get_value_groups(re *regexp.Regexp, content string, number_of_values_per_match int) [][]int{
    matches := re.FindAllStringSubmatch(content, -1)
    values := make([][]int, len(matches))
    for i := range matches {
        values[i] = make([]int, number_of_values_per_match)
        for j:=1; j< len(matches[i]); j++{
            val, _ := strconv.Atoi(matches[i][j])
            values[i][j-1] = val
        }
    }
    return values
}

func get_samples(content string) []Sample{
    var re = regexp.MustCompile(
`Before: \[(\d*), (\d*), (\d*), (\d*)\]
(\d*) (\d*) (\d*) (\d*)
After: *\[(\d*), (\d*), (\d*), (\d*)\]`)

    value_groups := get_value_groups(re, content, 12)
    samples := make([]Sample, len(value_groups))

    for i, values := range value_groups {
        samples[i] = Sample{
            start_registers: [4]int{values[0], values[1], values[2], values[3]},
            instruction: Instruction{
                opcode: values[4],
                content: [3]int{values[5], values[6], values[7]},
                },
            expected_result: [4]int{values[8], values[9], values[10], values[11]},
        }
    }

    return samples
}

func get_instructions(test_input string) []Instruction{
    var re = regexp.MustCompile(`(\d*) (\d*) (\d*) (\d*)`)

    value_groups := get_value_groups(re, test_input, 4)
    instructions := make([]Instruction, len(value_groups))

    for i, values := range value_groups {
        instructions[i] =  Instruction{
            opcode: values[0],
            content: [3]int{values[1], values[2], values[3]},
        }
    }

    return instructions
}

func addr(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]] + registers[instruction[1]]
    return result
}

func addi(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]] + instruction[1]
    return result
}

func mulr(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]] * registers[instruction[1]]
    return result
}

func muli(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]] * instruction[1]
    return result
}

func banr(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]] & registers[instruction[1]]
    return result
}

func bani(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]] & instruction[1]
    return result
}

func borr(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]] | registers[instruction[1]]
    return result
}

func bori(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]] | instruction[1]
    return result
}

func setr(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = registers[instruction[0]]
    return result
}

func seti(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}
    result[instruction[2]] = instruction[0]
    return result
}

func gtir(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}

    if instruction[0] > registers[instruction[1]] {
        result[instruction[2]] = 1
    } else {
        result[instruction[2]] = 0
    }
    return result
}

func gtri(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}

    if registers[instruction[0]] > instruction[1] {
        result[instruction[2]] = 1
    } else {
        result[instruction[2]] = 0
    }
    return result
}

func gtrr(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}

    if registers[instruction[0]] > registers[instruction[1]] {
        result[instruction[2]] = 1
    } else {
        result[instruction[2]] = 0
    }
    return result
}

func eqir(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}

    if instruction[0] == registers[instruction[1]] {
        result[instruction[2]] = 1
    } else {
        result[instruction[2]] = 0
    }
    return result
}

func eqri(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}

    if registers[instruction[0]] == instruction[1] {
        result[instruction[2]] = 1
    } else {
        result[instruction[2]] = 0
    }
    return result
}

func eqrr(registers [4]int, instruction [3]int) [4]int{
    result := [4]int{registers[0], registers[1], registers[2], registers[3]}

    if registers[instruction[0]] == registers[instruction[1]] {
        result[instruction[2]] = 1
    } else {
        result[instruction[2]] = 0
    }
    return result
}


func call_opcode(registers [4]int, opcode int, instruction [3]int) [4]int{

    switch ; opcode {
    case 0: return addr(registers, instruction)
    case 1: return addi(registers, instruction)
    case 2: return mulr(registers, instruction)
    case 3: return muli(registers, instruction)
    case 4: return banr(registers, instruction)
    case 5: return bani(registers, instruction)
    case 6: return borr(registers, instruction)
    case 7: return bori(registers, instruction)
    case 8: return setr(registers, instruction)
    case 9: return seti(registers, instruction)
    case 10: return gtir(registers, instruction)
    case 11: return gtri(registers, instruction)
    case 12: return gtrr(registers, instruction)
    case 13: return eqir(registers, instruction)
    case 14: return eqri(registers, instruction)
    case 15: return eqrr(registers, instruction)
    } 

    result := [4]int{0, 0, 0, 0}
    return result
}

type Sample struct{
    start_registers [4]int
    instruction Instruction
    expected_result [4]int
}

type Instruction struct{
    opcode int
    content [3]int
}

func part1(samples []Sample){

    var new_registers [4]int
    var number_of_valid_opcodes int
    number_of_samples := 0

    for _, sample := range samples {

        number_of_valid_opcodes = 0

        for opcode:=0; opcode<16; opcode++ {
            new_registers = call_opcode(sample.start_registers, opcode, sample.instruction.content)
            if new_registers == sample.expected_result{
                number_of_valid_opcodes += 1
            }
        }

        if number_of_valid_opcodes >= 3 {
            number_of_samples += 1
        }
    }

    fmt.Println("Solution part 1:", number_of_samples)
}

func get_invalid_codes(samples []Sample) map[int]map[int]bool{
    //which opcodes do not map to the expected result
    //key=opscode in sample, value is code in own listing

    invalid_codes := map[int]map[int]bool{} //key=opscode in sample, value is code in own listing
    for i:=0; i<16; i++{
        invalid_codes[i] = map[int]bool{}
    }

    var new_registers [4]int

    for _, sample := range samples {
        for opcode:=0; opcode<16; opcode++ {
            new_registers = call_opcode(sample.start_registers, opcode, sample.instruction.content)
            if new_registers != sample.expected_result{
                invalid_codes[sample.instruction.opcode][opcode] = true
            }
        }
    }

    return invalid_codes
}

func get_code_mapping(invalid_codes map[int]map[int]bool) map[int]int{
    //Get mapping from opcode in instruction to opcode own opcode

    code_mapping := map[int]int{}           //key=opscode in sample, value is code in own listing
    var opcode int
    opcode_found := true

    for ;opcode_found; {
        opcode_found = false
        for key, _ := range invalid_codes {

            if len(invalid_codes[key]) == 15{
                for i:=0; i<16; i++ {
                    if !invalid_codes[key][i] {
                        opcode = i
                        opcode_found = true
                        break
                    }
                }
                code_mapping[key] = opcode
                delete(invalid_codes, key)

                //add found opcode to invalid codes
                for key, _ := range invalid_codes {
                    invalid_codes[key][opcode] = true
                }
            } 
        }
    }

    return code_mapping
}

func part2(samples []Sample, instructions []Instruction){

    invalid_codes := get_invalid_codes(samples)       //key=opscode in sample, value is code in own listing
    code_mapping := get_code_mapping(invalid_codes)   //key=opscode in sample, value is code in own listing
    
    var opcode int
    registers := [4]int{0, 0, 0, 0}
    for _, instruction := range instructions {
        opcode = code_mapping[instruction.opcode]
        registers = call_opcode(registers, opcode, instruction.content)
    }

    fmt.Println("Solution part 2:", registers[0])

}

func main(){
    input := get_input("input1.txt")
    samples := get_samples(input)

    test_input := get_input("input2.txt")
    instructions := get_instructions(test_input)

    part1(samples)
    part2(samples, instructions)
}