from time import time
# from pylogic import compile_program, evaluate
from cylogic import compile_program, evaluate


def load_file(path):
    lines = []
    with open(path, 'r') as fh:
        for line in fh:
            lines.append(line.strip())
    return lines


def main():
    instructions = load_file('input_part2.txt')

    compiled_instructions, mem_map, memory = compile_program(instructions)

    start = time()
    evaluate(compiled_instructions, memory)
    duration = time() - start
    print(f"Value in register a: {memory[mem_map.get('a')]}")
    print(f"Finished in: {duration}!")


if __name__ == "__main__":
    main()
