package Entity

type Cooperativa struct {
	ID                int    `json:"ID"`
	RazonSocial       string `json:"RazonSocial"`
	MatriculaNacional int8   `json:"MatriculaNacional"`
}


type Cooperativas [] Cooperativa

const TableCooperativa = `
	CREATE TABLE IF NOT EXISTS cooperativa (
	    ID INT AUTO_INCREMENT PRIMARY KEY,
	    RazonSocial VARCHAR(255) NOT NULL,
	    MatriculaNacional VARCHAR(255) NOT NULL
	)  ENGINE=INNODB;
`

