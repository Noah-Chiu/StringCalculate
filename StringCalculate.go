// 字串計算(字串+數字、or字串+字串)
// 輸出會依照stringNum的位元數來輸出，多出來的就不輸出
func StringCalculate(stringNum string, inputNum interface{}) (outputNum string) {
	num := 0

	// 用於字串變數字
	switch temp := inputNum.(type) {
	case string:
		num, _ = strconv.Atoi(temp)
	case int:
		num = temp
	}

	carry := num
	for i := len(stringNum) - 1; i >= 0; i-- {
		// 數字加進位，加完後進位要清空(第一次的carry就是要加的數字)
		digitNum, _ := strconv.Atoi(string(stringNum[i]))
		digitNum += carry
		carry = 0
		// 超過9要進位
		if digitNum > 9 {
			carry = (digitNum) / 10
			digitNum = (digitNum) % 10
		}
		// 小於0要退位(-的num)
		if digitNum < 0 {
			digitNum = 10 + digitNum
			carry = -1
		}
		// 由右到左拼起來
		outputNum = fmt.Sprint(digitNum) + outputNum
	}

	return
}
