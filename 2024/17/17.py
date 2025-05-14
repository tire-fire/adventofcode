from abc import ABC, abstractmethod

with open('input', 'r') as f:
    _input = f.readlines()

class State:

    def __init__(self, registers: list[int]):
        self.registers = registers
        self.pointer = 0
        self.output = []

    def __str__(self) -> str:
        return ','.join(map(str, self.output))

    def __getitem__(self, index: int) -> int:
        return self.registers[index]

    def __setitem__(self, index: int, value: int):
        self.registers[index] = value

    def increment_pointer(self):
        self.pointer += 1

    def reset(self, a: int):
        self.registers[0] = a
        self.registers[1] = 0
        self.registers[2] = 0
        self.output = []
        self.pointer = 0

class Instruction(ABC):

    opcode = -1

    def __init__(self, operand):
        self.operand = operand

    @abstractmethod
    def run(self, state: State) -> State:
        ...

    def combo_operand(self, state: State) -> int:
        # Maps combo operands to their registers
        _mapping = {4: 0, 5: 1, 6: 2}
        if self.operand not in _mapping:
            return self.operand
        
        register_index = _mapping[self.operand]
        return state[register_index]

class AdvInstruction(Instruction):

    opcode = 0

    def run(self, state: State) -> State:
        denominator = self.combo_operand(state)
        division = state[0] // (2 ** denominator)
        state[0] = division
        state.increment_pointer()
        return state

class BxlInstruction(Instruction):

    opcode = 1

    def run(self, state: State) -> State:
        xor_operation = state[1] ^ self.operand
        state[1] = xor_operation
        state.increment_pointer()
        return state

class BstInstruction(Instruction):

    opcode = 2

    def run(self, state: State) -> State:
        combo_mod_8 = self.combo_operand(state) % 8
        state[1] = combo_mod_8
        state.increment_pointer()
        return state

class JnzInstruction(Instruction):

    opcode = 3

    def run(self, state: State) -> State:
        if state[0] != 0:
            state.pointer = self.operand
        else:
            state.increment_pointer()     
        return state

class BxcInstruction(Instruction):

    opcode = 4

    def run(self, state: State) -> State:
        xor_operation = state[1] ^ state[2]
        state[1] = xor_operation
        state.increment_pointer()
        return state
    
class OutInstruction(Instruction):

    opcode = 5

    def run(self, state: State) -> State:
        combo_mod_8 = self.combo_operand(state) % 8
        state.output.append(combo_mod_8)
        state.increment_pointer()
        return state
    
class BdvInstruction(Instruction):

    opcode = 6

    def run(self, state: State) -> State:
        denominator = self.combo_operand(state)
        division = state[0] // (2 ** denominator)
        state[1] = division
        state.increment_pointer()
        return state

class CdvInstruction(Instruction):

    opcode = 7

    def run(self, state: State) -> State:
        denominator = self.combo_operand(state)
        division = state[0] // (2 ** denominator)
        state[2] = division
        state.increment_pointer()
        return state


def load() -> tuple[State, list[Instruction]]:
    state = State([0, 0, 0])
    instructions = []
    opcode_mapping = {
        0: AdvInstruction,
        1: BxlInstruction,
        2: BstInstruction,
        3: JnzInstruction,
        4: BxcInstruction,
        5: OutInstruction,
        6: BdvInstruction,
        7: CdvInstruction
    }

    for line in _input:
        if line.startswith('Register'):
            num = int(line.rsplit(maxsplit=1)[-1])
            if 'A' in line:
                state[0] = num
            elif 'B' in line:
                state[1] = num
            elif 'C' in line:
                state[2] = num
        
        elif line.startswith('Program'):
            program = line.split(maxsplit=1)[-1].split(',')
            for i in range(0, len(program), 2):
                opcode = int(program[i])
                instruction = opcode_mapping[opcode]
                instructions.append(
                    instruction(operand=int(program[i+1]))
                )
            
    return state, instructions


def puzzle_one() -> str:
    state, instructions = load()
    while state.pointer < len(instructions):
        instructions[state.pointer].run(state)
    return state.__str__()

def oct_to_dec(oct: str) -> int:
    return int(oct, 8)

def dec_to_oct(dec: int) -> str:
    # format() with 'o' formats the integer as an octal string without the '0o' prefix.
    return format(dec, 'o')

def puzzle_two() -> int:
    state, instructions = load()

    output = []
    for instruction in instructions:
        output.append(instruction.opcode)
        output.append(instruction.operand)

    # Iterate backwards over the expected output, trying to find a value of A
    # that reproduces the substring output[i:]
    lower_bound = 0
    for i in range(len(output) - 1, -1, -1):

        # There is probably a more appropriate upper bound, but I can't figure it out.
        for a in range(lower_bound, lower_bound + (8 ** (len(output) - i))):
            state.reset(a)
            while state.pointer < len(instructions):
                instructions[state.pointer].run(state)

            if state.output == output[i:]:
                # If we matched the entire output, we've found our solution.
                if len(state.output) == len(output):
                    return a
                
                # If not complete, but we matched a partial suffix, we need to adjust the lower_bound.
                # Convert 'a' to base 8 and append '0' in octal, effectively "shifting" our search range.
                a_base_8 = dec_to_oct(a)
                lower_bound = oct_to_dec(a_base_8 + '0')
                break

if __name__ == "__main__":
    print(puzzle_one())
    print(puzzle_two())
