contained = 0
overlap = 0
for line in open("input.txt"):
    ranges = [
        set(range(int(elf.split("-")[0]), int(elf.split("-")[1]) + 1))
        for elf in line.split(",")
    ]

    if len(ranges[0].intersection(ranges[1])) == min(len(s) for s in ranges):
        contained += 1

    if len(ranges[0].intersection(ranges[1])) != 0:
        overlap += 1

print(contained)
print(overlap)
