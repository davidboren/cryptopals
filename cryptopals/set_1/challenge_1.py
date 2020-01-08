import codecs


def hex_to_b64(hex):
    return codecs.encode(codecs.decode(hex, "hex"), "base64").decode().rstrip()
