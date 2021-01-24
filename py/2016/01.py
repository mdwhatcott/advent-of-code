# With the exception of 1 line (marked below) and the final
# print statements this code was written in the paradigm of
# 'functional' programming. There are no reassignment
# statements or function calls with side-effects.

ORIGIN = (0, 0)

N = (0, 1)
S = (0, -1)
W = (-1, 0)
E = (1, 0)

R = {N:E, E:S, S:W, W:N}
L = {N:W, W:S, S:E, E:N}

def turn(facing, direction):
    return R[facing] if direction[0] == "R" else L[facing]

def all_steps(steps_so_far, legs, facing):
    if not legs:
        return steps_so_far

    leg = legs[0]
    direction = leg[0]
    now_facing = turn(facing, direction)
    distance = int(leg[1:])
    steps = [now_facing] * distance
    return all_steps(steps_so_far + steps, legs[1:], now_facing)

def step(at, step):
    return (at[0] + step[0],
            at[1] + step[1])

def walk_all(at, steps):
    if not steps:
        return at

    return walk_all(step(at, steps[0]), steps[1:])

def distance_from_origin(at):
    return abs(at[0]) + abs(at[1])

def part1(steps):
    end = walk_all(ORIGIN, steps)
    return distance_from_origin(end)

def set_add(set, new):
    set.add(new) # FUNCTION WITH SIDE EFFECTS
    return set

def walk_until_criss_cross(visited, at, steps):
    if at in visited:
        return at

    return walk_until_criss_cross(
        set_add(visited, at),
        step(at, steps[0]),
        steps[1:])

def part2(steps):
    criss_cross = walk_until_criss_cross(set(), ORIGIN, steps)
    return distance_from_origin(criss_cross)

def main():
    with open('01.txt') as input:
        legs = input.read().strip().split(', ')
        steps = all_steps([], legs, N)
        print("part 1:", part1(steps))
        print("part 2:", part2(steps))

if __name__ == '__main__':
    main()
