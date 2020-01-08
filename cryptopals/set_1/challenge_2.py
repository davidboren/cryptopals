import codecs
import binascii


def hex_to_ints(s1):
    return codecs.decode(s1, "hex")


def int_to_hex(i):
    return format(i, "02x")


def ints_to_hex(arr):
    return "".join(int_to_hex(x) for x in arr)


def hex_to_bytes(s1):
    return bytes(hex_to_ints(s1))


def xor_hex(s1, s2):
    return ints_to_hex([a ^ b for a, b in zip(hex_to_ints(s1), hex_to_ints(s2))])
