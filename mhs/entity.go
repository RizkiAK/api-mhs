package mhs

type Mahasiswa struct {
	ID     int    `json:"id"`
	Nim    int    `json:"nim" binding:"required"`
	Nama   string `json:"nama" binding:"required"`
	Email  string `json:"email" binding:"required,email"`
	Alamat string `json:"alamat" binding:"required"`
}

type InputMhs struct {
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
}

type InputMhsDetail struct {
	Nim int `uri:"nim"`
}

type MahasiswaFormatter struct {
	Nim  int    `json:"nim"`
	Nama string `json:"nama"`
}

func Formatter(mhs Mahasiswa) MahasiswaFormatter {
	formatter := MahasiswaFormatter{
		Nim:  mhs.Nim,
		Nama: mhs.Nama,
	}

	return formatter
}

func FormatterArray(mhs []Mahasiswa) []MahasiswaFormatter {
	formatter := []MahasiswaFormatter{}

	for _, mahasiswa := range mhs {
		format := Formatter(mahasiswa)
		formatter = append(formatter, format)
	}

	return formatter
}
