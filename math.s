// func add(a,b int) int
TEXT ·add(SB), $0-24
    MOVQ a+0(FP), AX
    ADDQ b+8(FP), AX
    MOVQ AX, r+16(FP)
    RET

// func sub(a,b int) int
TEXT ·sub(SB), $0-24
    MOVQ a+0(FP), AX
    SUBQ b+8(FP), AX
    MOVQ AX, ret+16(FP)
    RET
