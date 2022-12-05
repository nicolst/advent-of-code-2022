from collections import deque


def get_init_crates(lines):
    stacks = [deque() for i in range(9)]

    end = 0
    for i in range(len(lines)):
        if lines[i][1] == "1":
            end = i
            break
    for i in range(end - 1, -1, -1):
        for j in range(9):
            if lines[i][1 + 4 * j] != " ":
                stacks[j].append(lines[i][1 + 4 * j])
    return stacks, end + 2


def part_1(stacks, lines, start):
    stacks = stacks[:]
    for line in lines[start:]:
        cmd = line.strip().split(" ")

        times = int(cmd[1])
        fr = int(cmd[3]) - 1
        to = int(cmd[5]) - 1

        for i in range(times):
            stacks[to].append(stacks[fr].pop())
    return "".join((s[-1] for s in stacks))


def part_2(stacks, lines, start):
    stacks = stacks[:]
    for line in lines[start:]:
        print(line)
        cmd = line.strip().split(" ")
        times = int(cmd[1])
        fr = int(cmd[3]) - 1
        to = int(cmd[5]) - 1
        temp = deque()
        for j in range(times):
            temp.append(stacks[fr].pop())
        for j in range(times):
            stacks[to].append(temp.pop())

    return "".join((s[-1] for s in stacks))


lines = open("input.txt").readlines()
stacks, start = get_init_crates(lines)

# first = part_1(stacks, lines, start)
# print(first)
second = part_2(stacks, lines, start)
print(second)
print()
