package		main

import "os/exec"

const USHRT_MAX = 65535

type s_tram struct {

	guid    string
	hash	uint16
}

func getHash(str string) uint16 {
	var ret uint16

	ret = 0
	for _, value := range str {
		ret += uint16(value)
		ret += (ret << 10)
		ret ^= (ret >> 6)
	}

	ret += (ret << 3)
	ret ^= (ret >> 11)
	ret += (ret << 15)
	ret = (ret % USHRT_MAX)
	return ret
}

func		getGuid() string {

	cmd := exec.Command("sh", "-c", "/usr/bin/base64 /dev/random | /usr/bin/head -c 64")
	out, _ := cmd.Output()
	return string(out)
}

func		getTram() s_tram {

	guid := getGuid()
	hash := getHash(string(guid))
	tram := s_tram{guid, hash}
	return tram
}
