package crc

import "fmt"

/*
//https://my.oschina.net/iwuyang/blog/198630

CRC4运算

Poly生成项： 1001，长度为4
- 宽度为 w=3（第一位始终为1，异或时始终为0，所以可忽略第一位）

位串BitString ： 11110
- 需借位w  11110 + 000 = 11110000（确保位串每一位都被除一遍）

运算过程：
- 运算中只有位串首位为1才运算，否则跳过

11110000
1001|||| -
-------------
 1100|||
 1001||| -
------------
  1010||
  1001|| -
-----------
   0110|
   0000| -
----------
    1100
    1001 -
---------
     101 --> 5，余数 --> The CRC!
*/
func Crc4(data uint) uint {
	var poly uint = 0x13
	var w uint = 4
	var curBit int = 15
	var register uint = 0x0000

	//借位w=4
	data <<= w
	fmt.Printf("data append %d bit is  %b\n", w, data)

	for ; curBit >= 0; curBit-- {
		if (register>>w)&0x0001 == 0x1 { //register = register ^ POLY;
			fmt.Printf("register %b right shift %d bit is 1, xor with poly(%b):%b\n", register, w, poly, register^poly)
			register ^= poly
		}
		register <<= 1
		//读取目标串下一位
		tmp := data >> uint8(curBit) & 0x0001
		register |= tmp
		fmt.Printf("curBit:%d, register:%b\n", curBit, register)
	}
	if (register>>w)&0x0001 == 0x1 { //register = register ^ POLY;
		fmt.Printf("register %b right shift %d bit is 1, xor with poly(%b):%b\n", register, w, poly, register^poly)
		register ^= poly
		fmt.Printf("curBit:%d, register:%b\n", curBit, register)
	}

	return register
}
