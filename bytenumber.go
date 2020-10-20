package main


type ByteNumber struct {
    bytes []uint8
    base uint8
}


// converInt(19, 2, 8) => [0 0 0 1 0 0 1 1]
func NewByteNumber(num uint64, base uint8, size uint8) *ByteNumber {
    var bytes = make([]uint8, size)

    for i := int(size - 1); i >= 0; i-- {
        bytes[i] = uint8(num % uint64(base))
        num = num / uint64(base)
    }
    return &ByteNumber{bytes, base}
}


func (s *ByteNumber) Bytes() []uint8 {
    return s.bytes
}


func (s *ByteNumber) Base() uint8 {
    return s.base
}


// Increment byte number
func (s *ByteNumber) Inc() {
    var i = len(s.bytes) - 1

    for i >= 0 && s.bytes[i] == s.base - 1 {
        s.bytes[i] = 0
        i--
    } 
    s.bytes[i]++
}


// Convert ByteNumber to uint64
func (s *ByteNumber) ToUInt64() uint64 {
    var result uint64 = 0
    var rank uint64 = 1
    for i := len(s.bytes) - 1; i >= 0; i-- {
        result += uint64(s.bytes[i]) * rank
        rank *= uint64(s.base)
    }

    return result
}


// translateToChars([]uint8{2, 3, 1, 1, 0}, []uint8("olhe")) =>
// => [104 101 108 108 111] ~ "hello" (ASCII characters)
func (s *ByteNumber) Translate(alphabet []uint8) []uint8 {
    var res = make([]uint8, len(s.bytes))
    for i, ord := range s.bytes {
        res[i] = alphabet[ord]
    }
    return res
}