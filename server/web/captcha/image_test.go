// Copyright 2014 beego Author. All Rights Reserved.
//

package captcha

import (
	"testing"

	"github.com/W3-Engineers-Ltd/Radiant/core/utils"
)

type byteCounter struct {
	n int64
}

func (bc *byteCounter) Write(b []byte) (int, error) {
	bc.n += int64(len(b))
	return len(b), nil
}

func BenchmarkNewImage(b *testing.B) {
	b.StopTimer()
	d := utils.RandomCreateBytes(challengeNums, defaultChars...)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewImage(d, stdWidth, stdHeight)
	}
}

func BenchmarkImageWriteTo(b *testing.B) {
	b.StopTimer()
	d := utils.RandomCreateBytes(challengeNums, defaultChars...)
	b.StartTimer()
	counter := &byteCounter{}
	for i := 0; i < b.N; i++ {
		img := NewImage(d, stdWidth, stdHeight)
		img.WriteTo(counter)
		b.SetBytes(counter.n)
		counter.n = 0
	}
}
