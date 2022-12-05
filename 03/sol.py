a = ord("a")
z = ord("z")

A = ord("A")
Z = ord("Z")


def pri(c):
    v = ord(c)
    return v - a + 1 if a <= v <= z else v - A + 27


tot = 0
with open("input.txt") as f:
    while line := list(f.readline().strip()):
        if len(line) == 0:
            break
        dup = (
            set(line[: len(line) // 2]).intersection(set(line[len(line) // 2 :])).pop()
        )
        tot += pri(dup)

print(tot)

tot = 0
with open("input.txt") as f:
    while True:
        group = [set(f.readline().strip()) for i in range(3)]
        if len(group[0]) == 0:
            break
        tot += pri(group[0].intersection(group[1], group[2]).pop())

print(tot)
