from cryptopals.set_1 import (
    hex_to_b64,
    hex_to_bytes,
    xor_hex,
    get_most_likely_xor_char,
    get_most_likely_line,
    repeating_key_xor_encrypt,
    repeating_key_xor_decrypt,
    hamming_distance,
    get_ordered_keysizes,
    get_likely_xor_chars,
)


def test_challenge_1():
    assert (
        hex_to_b64(
            b"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
        )
        == "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
    )


def test_challenge_2():
    hex_res = xor_hex(
        "1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"
    )

    assert hex_res == "746865206b696420646f6e277420706c6179"
    assert hex_to_bytes(hex_res) == b"the kid don't play"


def test_challenge_3():
    assert get_most_likely_xor_char(
        b"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
    ) == ("58", -137.0327050351889, b"Cooking MC's like a pound of bacon")


def test_challenge_4():
    assert get_most_likely_line() == (
        170,
        "15",
        -118.40716405612973,
        b"7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f",
        b"nOW\x00THAT\x00THE\x00PARTY\x00IS\x00JUMPING*",
    )


def test_challenge_5():
    stanza = b"""
Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal
"""
    assert repeating_key_xor_encrypt(stanza, b"ICE") == (
        "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a262263242"
        "72765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
    )


def test_challenge_5_invertibility():
    stanza = b"""
Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal
"""
    assert (
        repeating_key_xor_decrypt(repeating_key_xor_encrypt(stanza, b"ICE"), b"ICE")
        == stanza
    )


def test_challenge_6_hamming():
    assert hamming_distance(b"this is a test", b"wokka wokka!!!") == 37


def test_challenge_6():
    assert get_ordered_keysizes()[0][0] == 40


def test_challenge_6_xor_key():
    bytes([int(x) for x in get_likely_xor_chars(get_ordered_keysizes()[0][0])])
