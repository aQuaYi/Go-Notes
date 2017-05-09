package prime

import "testing"

var ans1000 = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}

func Test_Prime_generate(t *testing.T) {
	k := 100
	c := make(chan int)
	p := &Prime{
		Chan:  c,
		limit: k,
	}
	src := p.generate()
	for i := 2; i <= k; i++ {
		g, isOpen := <-src
		if !isOpen {
			t.Error("generate的src提前关闭了")
		}
		if g != i {
			t.Error("generate发出的数据不对")
		}
	}
	_, isOpen := <-src
	if isOpen {
		t.Error("generate的src，发送完成后，没有关闭")
	}
}
func Test_Prime_generate_2(t *testing.T) {
	k := 2
	c := make(chan int)
	p := &Prime{
		Chan:  c,
		limit: k,
	}
	src := p.generate()
	for i := 2; i <= k; i++ {
		g, isOpen := <-src
		if !isOpen {
			t.Error("generate的src提前关闭了")
		}
		if g != i {
			t.Error("generate发出的数据不对")
		}
	}
	_, isOpen := <-src
	if isOpen {
		t.Error("generate的src，发送完成后，没有关闭")
	}
}

func Test_Prime_generate_1(t *testing.T) {
	k := 1
	c := make(chan int)
	p := &Prime{
		Chan:  c,
		limit: k,
	}
	src := p.generate()
	i, isOpen := <-src
	if isOpen {
		t.Error("generate的src没有关闭,收到了", i)
	}
}
func Test_NewUnder_1000(t *testing.T) {
	primeCh := NewUnder(1000)

	for _, v := range ans1000 {
		p, isOpening := <-primeCh
		if !isOpening {
			t.Error("PrimeCh提前关闭了。")
			break
		}
		if p != v {
			t.Errorf("生成的素数不对, p=%d, v=%d", p, v)
		}
	}

	_, isOpening := <-primeCh
	if isOpening {
		t.Error("发送完素数后，primeCh没有关闭")
	}

	for i := range primeCh {
		t.Error("发送完素数后，primeCh没有关闭，收到", i)
	}
}
func Test_NewUnder_1(t *testing.T) {
	primeCh := NewUnder(1)

	i, isOpening := <-primeCh
	if isOpening {
		t.Error("没有关闭,收到了", i)
	}
}

func Benchmark_NewUnder10K(b *testing.B) {
	for i := 1; i < b.N; i++ {
		primeCh := NewUnder(10 * 1000)
		for {
			_, isOpen := <-primeCh
			if !isOpen {
				break
			}
		}
	}
}
func Benchmark_NewUnder100K(b *testing.B) {
	for i := 1; i < b.N; i++ {
		primeCh := NewUnder(100 * 1000)
		for {
			_, isOpen := <-primeCh
			if !isOpen {
				break
			}
		}
	}
}

func Benchmark_NewUnder1M(b *testing.B) {
	for i := 1; i < b.N; i++ {
		primeCh := NewUnder(1000 * 1000)
		for {
			_, isOpen := <-primeCh
			if !isOpen {
				break
			}
		}
	}
}
