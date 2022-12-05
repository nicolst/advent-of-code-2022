print(
    sum(
        sorted(
            (
                sum(int(line) for line in section.split("\n"))
                for section in open("input.txt").read().strip().split("\n\n")
            )
        )[-3:]
    )
)
