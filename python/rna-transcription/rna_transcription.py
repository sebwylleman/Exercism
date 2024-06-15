def to_rna(dna_strand: str) -> str:
    rna_strand = []

    nucleotide_complement = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
    }

    for nucleotide in dna_strand:
        if nucleotide in nucleotide_complement:
            rna_strand.append(nucleotide_complement[nucleotide])

    return "".join(rna_strand)
