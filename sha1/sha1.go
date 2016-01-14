package sha1
import "fmt"

//const H [5]uint32 = [5]uint32{0x67452301, 0xEFCDAB89, 0x98BACDFE, 0x10325476, 0xC3D2E1F0}

func OutputBit(input []byte) {
	for _, v := range input {
		fmt.Printf("%b ", v)
	}
	return
}

func getLengthArray(len int64) [8]byte{
	var tmp [8]byte
	tmp[0] = byte(len >> 56)
	tmp[1] = byte(len >> 48)
	tmp[2] = byte(len >> 40)
	tmp[3] = byte(len >> 32)
	tmp[4] = byte(len >> 24)
	tmp[5] = byte(len >> 16)
	tmp[6] = byte(len >> 8)
	tmp[7] = byte(len)
	return tmp
}

func Padding(message []byte) []byte{
	var original_bit_len int64 = int64(len(message) * 8)
	message = append(message, 0x80)
	append_one_bit_length := len(message) * 8

	for append_one_bit_length % 512 != 448 {
		append_one_bit_length += 8
	}

	for i := len(message); i < append_one_bit_length / 8; i++ {
		message = append(message, 0x00)
	}

	lengthArray := getLengthArray(original_bit_len)

	for _, v := range lengthArray {
		message = append(message, v)
	}

	return message
}
