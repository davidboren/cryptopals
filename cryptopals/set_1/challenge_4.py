import os
from cryptopals.set_1 import get_most_likely_xor_char


def get_most_likely_line():
    with open(os.path.join(os.path.dirname(__file__), "challenge_4.txt"), "rb") as f:
        max_score = None
        best_line_number = None
        best_line = None
        best_line_char = None
        best_line_output = None
        for i, line in enumerate(f):
            char, score, s = get_most_likely_xor_char(line.rstrip())
            if max_score is None or score > max_score:
                max_score = score
                best_line_number = i
                best_line = line.rstrip()
                best_line_char = char
                best_line_output = s
    return best_line_number, best_line_char, max_score, best_line, best_line_output
