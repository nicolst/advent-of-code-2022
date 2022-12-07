contents = open("input.txt").read().strip()

for i in range(len(contents) - 4):
    if len(set(contents[i : i + 4])) == 4:
        print(i + 4)
        break

for i in range(len(contents) - 14):
    if len(set(contents[i : i + 14])) == 14:
        print(i + 14)
        break
