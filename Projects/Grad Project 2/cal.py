from hashlib import sha256
try:
    text = "013802bf0000000000000000"
    cwid = text[:8]
    hash = sha256(bytes.fromhex(text))
    hash_id = hash.hexdigest()[:8]
    while hash_id != cwid:
        text = int(text, 16)
        text += 1
        text = format(text,'x')
        text = text.zfill(24)
        hash = sha256(bytes.fromhex(text))
        hash_id = hash.hexdigest()[:8]
    cwid_id=int(cwid,base=16)
    code_id = int(text[8:24], base=16)
    print('Registered CWID : ', str(cwid_id))
    print('Unlock Code : ',str(code_id))
except KeyboardInterrupt:
    print(text)
