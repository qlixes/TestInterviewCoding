package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func cuti(cuti_bersama int, join_at string, cuti_at string, durasi int) (bool, string) {
	cuti_tahunan := 14
	layout := "2006-01-02"
	limit_days := 180

	// check join was valid date
	join_date, err := time.Parse(layout, join_at)

	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}

	cuti_date, err := time.Parse(layout, cuti_at)

	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}

	is_probation := cuti_date.Sub(join_date)
	diff := is_probation.Hours() / 24

	probation_days := int(math.Floor(diff))

	if probation_days < limit_days {
		return false, "Alasan : Karena belum 180 hari sejak tanggal join karyawan"
	}

	last_date := time.Date(cuti_date.Year(), 12, 31, 23, 59, 59, 0, cuti_date.Location())

	total_join := last_date.Sub(join_date)
	total_days := total_join.Hours() / 24

	cuti_require := int(math.Floor(total_days)) - limit_days

	cuti_pribadi := cuti_tahunan - cuti_bersama

	cuti_aktif := (cuti_require * cuti_pribadi) / 365

	if cuti_aktif < durasi {
		return false, fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti", cuti_aktif)
	}

	return true, ""
}

func main() {
	cuti_bersama := 7
	join_karyawan := "2021-05-01"
	cuti_karyawan := "2021-07-05"
	durasi := 1

	status, message := cuti(cuti_bersama, join_karyawan, cuti_karyawan, durasi)

	if status {
		fmt.Println("true")
	}

	fmt.Printf("%b, %s", status, message)

	cuti_bersama1 := 7
	join_karyawan1 := "2021-05-01"
	cuti_karyawan1 := "2021-11-05"
	durasi1 := 3

	status1, message1 := cuti(cuti_bersama1, join_karyawan1, cuti_karyawan1, durasi1)

	if status1 {
		fmt.Println("true")
	}

	fmt.Printf("%b, %s", status1, message1)
}
