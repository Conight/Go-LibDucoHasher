from ctypes import cdll, Structure, c_char_p, c_longlong, c_double

lib = cdll.LoadLibrary('/path/to/shared/main.so')


class GoString(Structure):
    _fields_ = [('p', c_char_p), ('n', c_longlong)]


class Return(Structure):
    _fields_ = [
        ('r0', c_longlong),
        ('r1', c_double),
    ]


# DUCOS1
lib.DUCOS1.argtypes = [GoString, GoString, c_longlong, c_double]
lib.DUCOS1.restype = Return

# DUCOS1Nonce
lib.DUCOS1Nonce.argtypes = [GoString, GoString, c_longlong, c_double]
lib.DUCOS1Nonce.restype = c_longlong


def DUCOS1Nonce(
    last_h: str, exp_h: str, diff: int, eff: float,
):
    last_h = GoString(last_h.encode(), 40)
    exp_h = GoString(exp_h.encode(), 40)

    r = lib.DUCOS1Nonce(
        last_h,
        exp_h,
        diff,
        eff,
    )

    return r


def DUCOS1(
    last_h: str, exp_h: str, diff: int, eff: float,
):
    last_h = GoString(last_h.encode(), 40)
    exp_h = GoString(exp_h.encode(), 40)

    r = lib.DUCOS1(
        last_h,
        exp_h,
        diff,
        eff,
    )

    return r.r0, r.r1


if __name__ == '__main__':
    last_h: str = 'f316f4cd012371b15da767fa66c6c7478bf9593e'
    exp_h: str = '98ce69810af441b69971eec6ba4d87b766bf0213'
    diff: int = 500000
    eff: float = 0.005
    print(DUCOS1(last_h, exp_h, diff, eff))
    print(DUCOS1Nonce(last_h, exp_h, diff, eff))
