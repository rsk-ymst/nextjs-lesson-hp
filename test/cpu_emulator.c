# include <stdio.h>

#define MOV  0 // 足し算
#define ADD  1 // 足し
#define SUB  2 // 引き
#define	AND  3 // and
#define OR   4 // or
#define SL   5 // shift left
#define SR   6 //
#define SRA  7
#define LDL  8
#define LDH  9
#define CMP  10
#define JE   11
#define JMP  12
#define LD   13
#define ST   14
#define HLT  15

#define REG0 0
#define REG1 1
#define REG2 2
#define REG3 3
#define REG4 4
#define REG5 5
#define REG6 6
#define REG7 7

short reg[8];
short rom[256];
short ram[256];

void assembler(void);

short mov(short, short);
short add(short, short);
short sub(short, short);

short mov(short ra, short rb) {
	printf("%d", (MOV << 11) | (ra << 8) | (rb << 5));
	return ((MOV << 11) | (ra << 8) | (rb << 5));
}

void main(void) {
	short pc;
	short ir;
	short flag_eq;

	assembler();

	pc = 0;

	//switch (op_code(ir)) {

//	}

	short a = mov(1, 2);
}

