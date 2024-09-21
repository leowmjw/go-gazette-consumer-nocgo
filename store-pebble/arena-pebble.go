package store_pebble

// There is 3rd party lib for Arena called Nuke (see Resources); but since std lib has it; should be considered
// Since Go v1.20; there is experimental support for Arena; but it is not to be used unless u know what u are doing
//	and even then; there may be better way (see below with GOMEMLIMIT) --> https://prog.world/go-1-20-and-the-memory-arena/
// Use memory + address sanitizer to check further: https://uptrace.dev/blog/golang-memory-arena.html#address-sanitizer
// Example of Binary Tree implemented with Arena: https://github.com/thepudds/arena-performance/blob/main/cmd/binarytree-arena/main.go
// Resources:
// - https://dev.to/marknefedov/go-120-memory-arena-4kch
// - https://github.com/ortuman/nuke

// Only implement Arena if necessary; since we will allocate memory heavy; might be wise to leverage on the new
//	GOMEMLIMIT availble since Go v1.19.  The "hack" using Go memory ballast that Twitch used previously is no longer
//	necessary; this "soft" limit; if we peg it to maybe like 90% of limits will be good enough for most OOM

// 	Arg passed in might be something like:
//  -e "GOMEMLIMIT=2750MiB" \
//  -e "GOGC=100" \
//  -e "GODEBUG=gctrace=1" \
// Resources:
// - https://weaviate.io/blog/gomemlimit-a-game-changer-for-high-memory-applications
