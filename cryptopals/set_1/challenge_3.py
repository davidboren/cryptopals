from functools import reduce
from math import log

from cryptopals.set_1 import xor_hex, int_to_hex, hex_to_ints, hex_to_bytes

ENGLISH_OCC = {
    "E": 12.0,
    "T": 9.10,
    "A": 8.12,
    "O": 7.68,
    "I": 7.31,
    "N": 6.95,
    "S": 6.28,
    "R": 6.02,
    "H": 5.92,
    "D": 4.32,
    "L": 3.98,
    "U": 2.88,
    "C": 2.71,
    "M": 2.61,
    "F": 2.30,
    "Y": 2.11,
    "W": 2.09,
    "G": 2.03,
    "P": 1.82,
    "B": 1.49,
    "V": 1.11,
    "K": 0.69,
    "X": 0.17,
    "Q": 0.11,
    "J": 0.10,
    "Z": 0.07,
}

ENGLISH_FREQS = {c: v / sum(ENGLISH_OCC.values()) for c, v in ENGLISH_OCC.items()}
MIN_ENGLISH_FREQ = min(ENGLISH_FREQS.values())


def xor_1char(s, c):
    return "".join(xor_hex(int_to_hex(i), c) for i in hex_to_ints(s))


def get_freqency_likelihood(s):
    s = [chr(c).upper() for c in s]
    charset = sorted(set(s))
    counts = {c: s.count(c) for c in charset}
    return reduce(
        lambda a, b: a + log(ENGLISH_FREQS.get(b, MIN_ENGLISH_FREQ) ** counts[b]),
        charset,
        0,
    )


def get_likely_xor_chars(s):
    chars = [int_to_hex(i) for i in range(127)]
    bytestrings = {c: hex_to_bytes(xor_1char(s, c)) for c in chars}
    likelihoods = {c: get_freqency_likelihood(bs) for c, bs in bytestrings.items()}
    return [
        (c, lh, bytestrings[c])
        for c, lh in sorted(likelihoods.items(), key=lambda item: -item[1])
    ]


def get_most_likely_xor_char(s):
    return get_likely_xor_chars(s)[0]
