def turn_dir(dir, turn):
    match (dir, turn):
        case ('N', 'L'): return 'W'
        case ('N', 'R'): return 'E'
        case ('E', 'L'): return 'N'
        case ('E', 'R'): return 'S'
        case ('S', 'L'): return 'E'
        case ('S', 'R'): return 'W'
        case ('W', 'L'): return 'S'
        case ('W', 'R'): return 'N'

def move(x, y, dir, length):
    match dir:
        case 'N': return (x, y + length)
        case 'E': return (x + length, y)
        case 'S': return (x, y - length)
        case 'W': return (x - length, y)

x = 0
y = 0

dir = 'N'

# input = 'R2, L3'
# input = 'R2, R2, R2'
# input = 'R5, L5, R5, R3'
input = 'R8, R4, R4, R8'
# input = 'R3, L5, R2, L1, L2, R5, L2, R2, L2, L2, L1, R2, L2, R4, R4, R1, L2, L3, R3, L1, R2, L2, L4, R4, R5, L3, R3, L3, L3, R4, R5, L3, R3, L5, L1, L2, R2, L1, R3, R1, L1, R187, L1, R2, R47, L5, L1, L2, R4, R3, L3, R3, R4, R1, R3, L1, L4, L1, R2, L1, R4, R5, L1, R77, L5, L4, R3, L2, R4, R5, R5, L2, L2, R2, R5, L2, R194, R5, L2, R4, L5, L4, L2, R5, L3, L2, L5, R5, R2, L3, R3, R1, L4, R2, L1, R5, L1, R5, L1, L1, R3, L1, R5, R2, R5, R5, L4, L5, L5, L5, R3, L2, L5, L4, R3, R1, R1, R4, L2, L4, R5, R5, R4, L2, L2, R5, R5, L5, L2, R4, R4, L4, R1, L3, R1, L1, L1, L1, L4, R5, R4, L4, L4, R5, R3, L2, L2, R3, R1, R4, L3, R1, L4, R3, L3, L2, R2, R2, R2, L1, L4, R3, R2, R2, L3, R2, L3, L2, R4, L2, R3, L4, R5, R4, R1, R5, R3'
steps = input.split(', ')

visited = {(0, 0)}

for step in steps:
    turn = step[0]
    length = int(step[1:])

    dir = turn_dir(dir, turn)

    x, y = move(x, y, dir, length)

    if (x, y) in visited:
        break

    visited.add((x, y))

print(f'Final position: {x}, {y}')

result = abs(x) + abs(y)

print(f'Final distance: {result}')