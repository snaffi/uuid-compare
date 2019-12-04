package uuid

// Compare uuid's as java.util.UUID  http://hg.openjdk.java.net/jdk8/jdk8/jdk/file/default/src/share/classes/java/util/UUID.java#l435
func Compare(x, y [16]byte) int {
	xMostBits, xLeastBits := parseSigBits(x)
	yMostBits, yLeastBits := parseSigBits(y)
	if xMostBits < yMostBits {
		return -1
	}
	if xMostBits > yMostBits {
		return 1
	}
	if xLeastBits < yLeastBits {
		return -1
	}
	if xLeastBits > yLeastBits {
		return 1
	}
	return 0
}

// http://hg.openjdk.java.net/jdk8/jdk8/jdk/file/default/src/share/classes/java/util/UUID.java#l104
func parseSigBits(uuid [16]byte) (mostBits, leastBits int64) {
	for i := 0; i < 8; i++ {
		mostBits = (mostBits << 8) | int64(uuid[i]&0xff)
	}
	for i := 8; i < 16; i++ {
		leastBits = (leastBits << 8) | int64(uuid[i]&0xff)
	}
	return
}

