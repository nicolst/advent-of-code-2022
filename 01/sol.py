print(
    sum(
        sorted(
            (
                sum(int(line.strip()) for line in section.split("\n"))
                for section in open("input.txt").read()[:-1].split("\n\n")
            )
        )[-3:]
    )
)
