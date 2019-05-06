// +build amd64

// Code autogenerated by go generate; DO NOT EDIT.

#include "fp255_amd64.h"

// func cSelect255(x, y *Element255, b uint)
TEXT ·cSelect255(SB),NOSPLIT,$0
    MOVQ x+0(FP), DI
    MOVQ y+8(FP), SI
    MOVQ b+16(FP), BX
    cselect(0(DI),0(SI),BX)
    RET

// func cSwap255(x, y *Element255, b uint)
TEXT ·cSwap255(SB),NOSPLIT,$0
    MOVQ x+0(FP), DI
    MOVQ y+8(FP), SI
    MOVQ b+16(FP), BX
    cswap(0(DI),0(SI),BX)
    RET

// func addSub255(z, x *Element255)
TEXT ·addSub255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    addSub(0(DI),0(SI))
    RET

// func addLeg255(z,x,y *Element255)
TEXT ·addLeg255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    MOVQ y+16(FP), BX
    addition(0(DI),0(SI),0(BX))
    RET

// func addAdx255(z,x,y *Element255)
TEXT ·addAdx255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    MOVQ y+16(FP), BX
    additionAdx(0(DI),0(SI),0(BX))
    RET

// func sub255(z,x,y *Element255)
TEXT ·sub255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    MOVQ y+16(FP), BX
    subtraction(0(DI),0(SI),0(BX))
    RET

// func mulA24255(z, x *Element255)
TEXT ·mulA24255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    multiplyA24(0(DI),0(SI))
    RET

// func mulA24Adx255(z, x *Element255)
TEXT ·mulA24Adx255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    multiplyA24Adx(0(DI),0(SI))
    RET

// func intMul255(z *bigElement255, x, y *Element255)
TEXT ·intMul255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    MOVQ y+16(FP), BX
    integerMul(0(DI),0(SI),0(BX))
    RET

// func intMulAdx255(z *bigElement255, x, y *Element255)
TEXT ·intMulAdx255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    MOVQ y+16(FP), BX
    integerMulAdx(0(DI),0(SI),0(BX))
    RET

// func intSqr255(z *bigElement255, x *Element255)
TEXT ·intSqr255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    integerSqr(0(DI),0(SI))
    RET

// func intSqrAdx255(z *bigElement255, x *Element255)
TEXT ·intSqrAdx255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    integerSqrAdx(0(DI),0(SI))
    RET

// func reduce255(z *Element255, x *bigElement255)
TEXT ·reduce255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    reduceFromDouble(0(DI),0(SI))
    RET

// func reduceAdx255(z *Element255, x *bigElement255)
TEXT ·reduceAdx255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ x+8(FP), SI
    reduceFromDoubleAdx(0(DI),0(SI))
    RET

// func sqrn255(z *Element255, buffer *bigElement255, times uint)
TEXT ·sqrn255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ buffer+8(FP), SI
    MOVQ times+16(FP),BX
    L0:
        CMPQ BX, $0
        JZ L1
        integerSqr(0(SI),0(DI))
        reduceFromDouble(0(DI),0(SI))
        DECQ BX
        JMP L0
    L1:
    RET

// func sqrnAdx255(z *Element255, buffer *bigElement255, times uint)
TEXT ·sqrnAdx255(SB),NOSPLIT,$0
    MOVQ z+0(FP), DI
    MOVQ buffer+8(FP), SI
    MOVQ times+16(FP),BX
    L0:
        CMPQ BX, $0
        JZ L1
        integerSqrAdx(0(SI),0(DI))
        reduceFromDoubleAdx(0(DI),0(SI))
        DECQ BX
        JMP L0
    L1:
    RET

// Code autogenerated by go generate; DO NOT EDIT.
