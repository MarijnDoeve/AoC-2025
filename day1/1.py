import fileinput


dial = 50
count1 = 0
count2 = 0

for line in fileinput.input():
    direction = line[0]
    num = int(line[1:])

    if direction == "R":
        count2 += (dial + num) // 100
        dial = (dial + num) % 100
    else:
        new_dial = dial - num
        if (new_dial) < 0:
            count2 += abs((100 + new_dial) // 100)
            if dial != 0:
                count2 += 1

        dial = (new_dial) % 100

        if dial == 0:
            count2 += 1


    if dial == 0:
        count1 += 1

print(f"Result 1: {count1}")
print(f"Result 2: {count2}")
