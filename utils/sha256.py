import hashlib
def hash_str(s, sk):
    data = "{}{}".format(s,sk)
    return hashlib.sha256(data.encode('utf-8')).hexdigest()