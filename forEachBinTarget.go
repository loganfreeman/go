type binaryFunc func(bin binary, OS, arch string)

// forEachBinTarget runs a function for all the target platforms of a binary in
// parallel.
func forEachBinTarget(bin binary, fn binaryFunc) {
	var wg sync.WaitGroup
	for _, t := range bin.targets {
		wg.Add(1)
		go func(bin binary, os, arch string) {
			defer wg.Done()
			fn(bin, os, arch)
		}(bin, t.os, t.arch)
	}
	wg.Wait()
}
