package main

import "testing"

type Test struct {
	in  []byte
	err error
	out int
}

var tests = []Test{
	{[]byte("Hello, World!"), nil, 13},
	{[]byte("This is a test."), nil, 15},
	{[]byte("AEIOUaeiou"), nil, 10},
	{[]byte("1234567890"), nil, 10},
	{[]byte(""), nil, 0},
	{[]byte{0xff}, ErrInvalidUTF8, 0},
}

func TestGetUTFLength(t *testing.T) {
	for i, test := range tests {
		result, err := GetUTFLength(test.in)
		if err != nil {
			if err != test.err {
				t.Errorf("#%d: GetUTFLength(%s) returned error %v; want %v", i, test.in, err, test.err)
			}
		} else {
			if result != test.out {
				t.Errorf("#%d: GetUTFLength(%s)=%d; want %d", i, test.in, result, test.out)
			}
		}
	}
}
