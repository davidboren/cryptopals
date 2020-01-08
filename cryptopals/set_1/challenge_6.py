import os
from statistics import mean
from math import ceil
from cryptopals.set_1 import (
    get_most_likely_xor_char,
    ints_to_hex,
    repeating_key_xor_decrypt,
)


def bytes_to_bits(s):
    return "".join(y for y in bin(int.from_bytes(s, byteorder="big")) if y != "b")


def hamming_distance(s1, s2):
    return sum(x != y for x, y in zip(bytes_to_bits(s1), bytes_to_bits(s2)))


def hamming_by_keysize(s, key_size):
    return hamming_distance(s[:key_size], s[key_size : 2 * key_size]) / key_size


def get_lines():
    with open(os.path.join(os.path.dirname(__file__), "challenge_6.txt"), "rb") as f:
        return [line.rstrip() for line in f]


def get_likely_xor_chars(key_size):
    all_bytes = [b for l in get_lines() for b in l]

    return [
        get_most_likely_xor_char(
            ints_to_hex(
                [
                    all_bytes[key_size * j + i]
                    for j in range(ceil(len(all_bytes) / key_size))
                    if key_size * j + i < len(all_bytes)
                ]
            )
        )[0]
        for i in range(key_size)
    ]


def get_ordered_keysizes():
    distances = {
        key_size: mean(hamming_by_keysize(l, key_size) for l in get_lines())
        for key_size in range(2, 41)
    }
    return [(ks, hd) for ks, hd in sorted(distances.items(), key=lambda item: item[1])]


def decode_rotating_xor():
    key = bytes([int(x) for x in get_likely_xor_chars(get_ordered_keysizes()[0][0])])
    return repeating_key_xor_encrypt(b"".join(get_lines()))
