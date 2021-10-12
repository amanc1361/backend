package modelsout

type Taraz struct {
	Rcode       int64
	Rid         int64
	Rname       string
	Rdebtor     int64
	Rcreditor   int64
	Rnature     byte
	Lcode       int64
	Lname       string
	Lid         int64
	Ldebtor     int64
	Lcreditor   int64
	Lnature     byte
	Fcode       int64
	Fname       string
	Fid         int64
	Fdebtor     int64
	Fcreditor   int64
	Fnature     byte
	Rparentcode int
	Rparentname string
	Fparentcode int
	Fparentname string
	Lparentcode int
	Lparentname string
}
