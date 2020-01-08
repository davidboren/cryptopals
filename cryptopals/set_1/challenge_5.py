from cryptopals.set_1 import ints_to_hex, hex_to_ints


def repeating_key_xor_encrypt(s, key):
    return ints_to_hex([a ^ key[i % len(key)] for i, a in enumerate(s)])


def repeating_key_xor_decrypt(s, key):
    return bytes([a ^ key[i % len(key)] for i, a in enumerate(hex_to_ints(s))])
