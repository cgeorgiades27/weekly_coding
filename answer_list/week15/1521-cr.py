# Gets first char from input_b, splits input at index and concats, checks equality.



def solution(input_a, input_b) -> bool:

    first_char = input_b[0]
    fc_index = input_a.find(first_char)
    shifted = input_a[fc_index:len(input_a)] + input_a[:fc_index]

    if shifted == input_b:
        return True
    else:
        return False


A = "abcde"
B = "cdeab"
