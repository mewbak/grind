package p

func f() {
	var i int
	switch x {
	case 1:
		for i = 0; i < 10; i++ {
			use(i)
		}
	case 2:
		for i = 0; i < 10; i++ {
			use(i)
		}
	}

	var j int
	switch x {
	case 1:
		for j = 0; j < 10; j++ {
			_ = &j
			use(j)
		}
		for j = 0; j < 10; j++ {
			_ = &j
			use(j)
		}
	case 2:
		for j = 0; j < 10; j++ {
			use(j)
		}
	}

	var k int
	switch x {
	case 1:
		for k = 0; k < 10; k++ {
			_ = &k
			use(k)
		}
		for k = 0; k < 10; k++ {
			_ = &k
			use(k)
		}
	case 2:
		for k = 0; k < 10; k++ {
			use(k)
		}
	}
	use(k)
}
