def parity(x):
    p = 0
    while x:
        p ^= x&1
        x >>= 1
    return p


def main():
    print 'uint8_t byte_parity[256] = {'
    for byte in range(256):
        print '{parity},'.format(parity=parity(byte)),
        if (byte + 1) % 16 == 0:
            print
    print '};'


if __name__ == '__main__':
    main()
