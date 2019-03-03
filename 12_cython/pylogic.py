INC_REG = 1
DEC_REG = 2
CPY_REG = 3
CPY_VAL = 4
JNZ_REG = 5
JNZ_VAL = 6


def compile_program(program):
    registers = []
    reg_map = {}
    parse_instructions = []

    def get_reg_index(reg_name):
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


def evaluate(program, memory):
    pc = 0    
    prog_len = len(program)
    while pc < prog_len:
        # pc = _eval_instruction(program[pc], memory, pc)
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


def _eval_instruction(inst, memory, pc):  # pc => program counter
    npc = pc + 1  # This is the normal case -- only jump is different

    action = inst[0]
    param1 = inst[1]

    if action == INC_REG:
        memory[param1] += 1
    elif action == DEC_REG:
        memory[param1] -= 1
    elif action == CPY_REG:
        memory[ inst[2] ] = memory[param1]
    elif action == CPY_VAL:
        memory[ inst[2] ] = param1
    elif action == JNZ_VAL and param1 != 0:
        npc = pc + inst[2]
    elif action == JNZ_REG and memory[param1] != 0:
        npc = pc + inst[2]

    return npc


def is_num(val):
    for c in val:
        if c not in '0123456789-':
            return False
    return True
