from functools import cmp_to_key

packet_pairs = [
    [eval(line.strip()) for line in pair.split("\n")]
    for pair in open("input.txt").read().split("\n\n")[:-1]
]


def compare(list1, list2):
    for i in range(max((len(l) for l in (list1, list2)))):
        if i >= len(list1):
            return True
        elif i >= len(list2):
            return False

        if type(list1[i]) == list or type(list2[i]) == list:
            l1 = [list1[i]] if type(list1[i]) == int else list1[i]
            l2 = [list2[i]] if type(list2[i]) == int else list2[i]
            res = compare(l1, l2)
            if res is not None:
                return res
        else:
            if list1[i] < list2[i]:
                return True
            elif list1[i] > list2[i]:
                return False

    return None


tot = 0
for i in range(len(packet_pairs)):
    p1, p2 = packet_pairs[i]
    tot += i + 1 if compare(p1, p2) else 0

print("Sum of indices is", tot)

all_signals = [l for pair in packet_pairs for l in pair]
all_signals.append([[2]])
all_signals.append([[6]])


def comparator(signal1, signal2):
    res = compare(signal1, signal2)
    if res:
        return 1
    elif res == False:
        return -1
    else:
        return 0


all_signals.sort(key=cmp_to_key(comparator), reverse=True)

decode1 = 0
decode2 = 0
for i in range(len(all_signals)):
    if all_signals[i].__str__() == "[[2]]":
        decode1 = i + 1
    elif all_signals[i].__str__() == "[[6]]":
        decode2 = i + 1

print("Decoder key:", decode1 * decode2)
