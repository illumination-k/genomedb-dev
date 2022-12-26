package interproscan

/*
https://interproscan-docs.readthedocs.io/en/latest/OutputFormats.html
Protein accession (e.g. P51587)
Sequence MD5 digest (e.g. 14086411a2cdf1c4cba63020e1622579)
Sequence length (e.g. 3418)
Analysis (e.g. Pfam / PRINTS / Gene3D)
Signature accession (e.g. PF09103 / G3DSA:2.40.50.140)
Signature description (e.g. BRCA2 repeat profile)
Start location
Stop location
Score - is the e-value (or score) of the match reported by member database method (e.g. 3.1E-52)
Status - is the status of the match (T: true)
Date - is the date of the run
InterPro annotations - accession (e.g. IPR002093)
InterPro annotations - description (e.g. BRCA2 repeat)
(GO annotations (e.g. GO:0005515) - optional column; only displayed if –goterms option is switched on)
(Pathways annotations (e.g. REACT_71) - optional column; only displayed if –pathways option is switched on)
*/

type InterproscanAnnotation struct {
	Accession   string
	Description string
}

type InterproscanRecord struct {
	Accession               string
	Length                  int
	Start                   int
	Stop                    int
	Analysis                string
	SignatureAccession      string
	SignatureDescription    string
	Score                   float64
	InterproscanAccession   string
	InterproscanDescription string
	GoTermIDs               []string
	PathwayIDs              []string
}
