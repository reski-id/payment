package main

import (
	"fmt"
	"math"
	"time"
)

func IsAllowedPersonalLeave(publicHolidays int, joinDate time.Time, plannedLeaveDate time.Time, leaveDuration int) (bool, string) {
	// Menghitung jumlah cuti kantor
	annualLeave := 14

	// Menghitung jumlah cuti bersama
	publicHolidaysInYear := publicHolidays

	// Menghitung jumlah cuti pribadi
	personalLeave := annualLeave - publicHolidaysInYear

	// Menghitung tanggal 180 hari setelah tanggal bergabung
	eligibleDate := joinDate.AddDate(0, 0, 180)
	if plannedLeaveDate.Before(eligibleDate) {
		return false, "Belum 180 hari sejak tanggal join karyawan"
	}

	remainingWorkDays := calculateRemainingWorkDays(plannedLeaveDate, time.Date(plannedLeaveDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC))
	quota := int(math.Floor(float64(remainingWorkDays) / 365 * float64(personalLeave)))

	// Menghitung jumlah cuti pribadi maksimum yang dapat diambil secara berturut-turut
	maxConsecutiveLeave := 3
	if leaveDuration > quota {
		return false, "Jumlah cuti pribadi yang tersedia tidak mencukupi"
	}
	if leaveDuration > maxConsecutiveLeave {
		return false, "Durasi cuti pribadi melebihi batas berturut-turut"
	}

	return true, ""
}

func calculateRemainingWorkDays(startDate time.Time, endDate time.Time) int {
	workDays := 0
	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
			workDays++
		}
	}

	return workDays
}

func main() {
	publicHolidays := 7
	joinDate := time.Date(2021, time.May, 1, 0, 0, 0, 0, time.UTC)
	plannedLeaveDate := time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC)
	leaveDuration := 1

	allowed, reason := IsAllowedPersonalLeave(publicHolidays, joinDate, plannedLeaveDate, leaveDuration)
	fmt.Println("Is Allowed Personal Leave:", allowed)
	fmt.Println("Reason:", reason)
}
