from time import time


INC_REG = 1
DEC_REG = 2
CPY_REG = 3
CPY_VAL = 4
JNZ_REG = 5
JNZ_VAL = 6


def load_file(path: str) -> list[str]:
    lines = []
    with open(path, 'r') as fh:
        for line in fh:
            lines.append(line.strip())
    return lines


def main():
    instructions = load_file('input_part2.txt')

    start = time()
    compiled_instructions, mem_map, memory = compile_program(instructions)
    duration = time() - start
    print(f"Compilation: {duration}!")


    start = time()
    evaluate(compiled_instructions, memory)
    duration = time() - start
    print(f"Value in register a: {memory[mem_map.get('a')]}")
    print(f"Finished in: {duration}!")


def compile_program(program: list[str]) -> tuple[list[tuple[int, int, int]], dict[str, int], list[int]]:
    """
    compile_program takes the input (a list of strings that roughly correlate to assembly instructions)
    and coverts it into:
        - a tuple of [int, int, int], which is (the action to take, the first parameter, the second parameter)
        - a dictionary that maps the name of the register (e.g. "a") into an index
        - a list of integers, which is initialized to all 0s. Each integer corresponds to one of the registers
          named in the dictionary. Note that the first index is not guaranteed to correspond to register "a"
    We assume the input it correct.

    To use this data, run the "program", coupled with "memory". Mutate the memory. When you are done, look up the
    register value with the dictionary.
    """
    registers: list[int] = []
    reg_map: dict[str, int] = {}
    parse_instructions: list[tuple[int, int, int]] = []

    def get_reg_index(reg_name: str) -> int:
        val = reg_map.get(reg_name)
        if val is None:
            val = len(registers)
            reg_map[reg_name] = val
            registers.append(0)  # initialize registers
        return val

    for line in program:
        instruction = line.split(' ')
        action = instruction[0]
        if action == 'inc':
            register = get_reg_index(instruction[1])
            parse_instructions.append((INC_REG, register))
        elif action == 'dec':
            register = get_reg_index(instruction[1])
            parse_instructions.append((DEC_REG, register))
        elif action == 'cpy':
            frm = instruction[1]
            to_idx = get_reg_index(instruction[2])
            if is_num(frm):
                parse_instructions.append((CPY_VAL, int(frm), to_idx))
            else:
                parse_instructions.append((CPY_REG, get_reg_index(frm), to_idx))
        elif action == 'jnz':
            val = instruction[1]
            direction = int(instruction[2])
            if is_num(val):
                parse_instructions.append((JNZ_VAL, int(val), direction))
            else:
                parse_instructions.append((JNZ_REG, get_reg_index(val), direction))

    return parse_instructions, reg_map, registers


def evaluate(program: list[tuple[int, int, int]], memory: list[int]) -> None:
    pc = 0    
    prog_len = len(program)
    while pc < prog_len:
        npc = pc + 1  # This is the normal case -- only jump is different

        cmd = program[pc][0]
        param1 = program[pc][1]

        if cmd == INC_REG:
            memory[param1] += 1
        elif cmd == DEC_REG:
            memory[param1] -= 1
        elif cmd == CPY_REG:
            memory[ program[pc][2] ] = memory[param1]
        elif cmd == CPY_VAL:
            memory[ program[pc][2] ] = param1
        elif cmd == JNZ_VAL and param1 != 0:
            npc = pc + program[pc][2]
        elif cmd == JNZ_REG and memory[param1] != 0:
            npc = pc + program[pc][2]
        pc = npc


def is_num(val: str) -> bool:
    for c in val:
        if c not in '0123456789-':
            return False
    return True


if __name__ == "__main__":
    main()
