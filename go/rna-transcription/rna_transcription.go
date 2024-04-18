package strand

func ToRNA(dna string) string {

	nucleotideComplement := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	rna := ""

	for _, char := range dna {
		_, ok := nucleotideComplement[char]
		if !ok {
			return ""
		}
		rna += string(nucleotideComplement[char])
	}
	return rna
}
