# O(n)

def get_edit_distance(word_a: str, word_b: str) -> int:
    """
    Takes two strings and returns the edit distance.
    """
    # base cases
    length_a = len(word_a)
    length_b = len(word_b)
    max_length = max(length_a, length_b)

    if word_a == word_b:
        return 0

    if length_a == 0 or length_b == 0:
        return max_length

    else:
        min_length = min(length_a, length_b)
        edit_count = (max_length - min_length)

        for i in range(min_length):
            if word_a[i] == word_b[i]:
                pass
            else:
                edit_count += 1

        return edit_count


# Tests
a = "kitten"
b = "sitting"
c = "boise"
d = "noisy"
e = ""
f = "kitten"

print(get_edit_distance(a, b))
print(get_edit_distance(c, d))
print(get_edit_distance(e, f))
print(get_edit_distance(a, f))
