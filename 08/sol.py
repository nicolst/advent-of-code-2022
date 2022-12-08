import numpy as np
from PIL import Image

forest = [list(r.strip()) for r in open("input.txt").readlines()]
numeric_forest = np.array([np.array([int(v) for v in r]) for r in forest])

height, width = numeric_forest.shape

mask = np.full(numeric_forest.shape, False)

view = np.full((height, width, 4), 0)

for r in range(height):
    leftidx = np.full((10), 0)
    rightidx = np.full((10), 0)

    rightmax = -1
    leftmax = -1
    for c in range(width):
        view[r, c, 2] = c - max(leftidx[numeric_forest[r, c] :])
        leftidx[numeric_forest[r, c]] = c

        view[r, width - c - 1, 0] = c - max(
            rightidx[numeric_forest[r, width - c - 1] :]
        )
        rightidx[numeric_forest[r, width - c - 1]] = c

        if numeric_forest[r, c] > rightmax:
            mask[r, c] = True
            rightmax = numeric_forest[r, c]

        if numeric_forest[r, width - c - 1] > leftmax:
            mask[r, width - c - 1] = True
            leftmax = numeric_forest[r, width - c - 1]

for c in range(width):
    upidx = np.full((10), 0j)
    downidx = np.full((10), 0j)

    downmax = -1
    upmax = -1
    for r in range(height):
        view[r, c, 1] = r - max(upidx[numeric_forest[r, c] :])
        upidx[numeric_forest[r, c]] = r

        view[height - r - 1, c, 3] = r - max(
            downidx[numeric_forest[height - r - 1, c] :]
        )
        downidx[numeric_forest[height - r - 1, c]] = r

        if numeric_forest[r, c] > downmax:
            mask[r, c] = True
            downmax = numeric_forest[r, c]

        if numeric_forest[height - r - 1, c] > upmax:
            mask[height - r - 1, c] = True
            upmax = numeric_forest[height - r - 1, c]


print("Trees visible from the outside:", np.count_nonzero(mask))

actual_view = np.apply_along_axis(np.prod, 2, view)

print("The tree with the best view has 'view value'", np.max(actual_view))

image_mapping = (255 // 9) * actual_view
img = Image.fromarray(np.uint8(image_mapping))
img.save("forest_view.png")
