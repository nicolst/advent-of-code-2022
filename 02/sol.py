rounds = [line.split(" ") for line in open("input.txt").read().split("\n")[:-1]]

# Do not judge
scores = {
    "A": {"X": 1 + 3, "Y": 2 + 6, "Z": 3 + 0},
    "B": {"X": 1 + 0, "Y": 2 + 3, "Z": 3 + 6},
    "C": {"X": 1 + 6, "Y": 2 + 0, "Z": 3 + 3},
}

scores2 = {
    "A": {"X": 3 + 0, "Y": 1 + 3, "Z": 2 + 6},
    "B": {"X": 1 + 0, "Y": 2 + 3, "Z": 3 + 6},
    "C": {"X": 2 + 0, "Y": 3 + 3, "Z": 1 + 6},
}


score = 0
score2 = 0
for r in rounds:
    score += scores[r[0]][r[1]]
    score2 += scores2[r[0]][r[1]]

print(score)
print(score2)
