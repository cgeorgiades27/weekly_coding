# Bottom-up
def get_the_bag(matrx):
    """Returns the maximum amount of coins that can be gained
    by following a path through the matrix"""

    rows = len(matrx)
    cols = len(matrx[rows - 1])
    sm_mat = [[0 for i in range(cols)] for j in range(rows)]

    # Traverses top row of matrx. Sets sm_mat[0][i] to tot.
    i = 0
    tot = 0
    for cell in matrx[0]:
        tot += cell
        sm_mat[0][i] = tot
        i += 1

    # Traverses left col of matrx. Sets sm_mat[i][0] to tot.
    tot = 0
    for i in range(rows):
        tot += matrx[i][0]
        sm_mat[i][0] = tot

    # Traverses remaining rows/cols of matrx. Sets sm_mat[row][col] to max_sum.
    # Two sums are calculated, cell + sm cell above, and cell + sm cell on left
    for row in range(1, rows):
        for col in range(1, cols):
            cell = matrx[row][col]
            lsum = cell + sm_mat[row][col - 1]
            rsum = cell + sm_mat[row - 1][col]
            max_sum = max(lsum, rsum)
            sm_mat[row][col] = max_sum

    # The bottom right cell in sm_mat holds the max possible $$$. The mf bag 🤑🤑
    the_bag = sm_mat[rows - 1][cols - 1]

    return the_bag


# Test
test_mat = [
    [0, 3, 1, 1],
    [2, 0, 0, 4],
    [1, 5, 3, 1]]

print(get_the_bag(test_mat))
