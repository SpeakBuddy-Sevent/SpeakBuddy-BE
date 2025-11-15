package seeder

import (
	"speakbuddy/internal/models"

	"gorm.io/gorm"
)

func SeedExercises(db *gorm.DB) error {
	var count int64
	db.Model(&models.ReadingExerciseTemplate{}).Count(&count)
	if count > 0 {
		// Sudah ada data, skip
		return nil
	}

	// Exercise Level: DASAR
	exerciseBasic := models.ReadingExerciseTemplate{
		Title: "Latihan Membaca - Dasar",
		Level: "dasar",
	}
	if err := db.Create(&exerciseBasic).Error; err != nil {
		return err
	}

	basicItems := []models.ExerciseItem{
		{
			ExerciseID: exerciseBasic.ID,
			ItemNumber: 1,
			TargetText: "Halo, nama saya Budi",
		},
		{
			ExerciseID: exerciseBasic.ID,
			ItemNumber: 2,
			TargetText: "Aku suka bermain bola",
		},
		{
			ExerciseID: exerciseBasic.ID,
			ItemNumber: 3,
			TargetText: "Ini rumah saya",
		},
		{
			ExerciseID: exerciseBasic.ID,
			ItemNumber: 4,
			TargetText: "Hari ini cerah sekali",
		},
		{
			ExerciseID: exerciseBasic.ID,
			ItemNumber: 5,
			TargetText: "Saya makan nasi goreng",
		},
	}
	for _, item := range basicItems {
		if err := db.Create(&item).Error; err != nil {
			return err
		}
	}

	// Exercise Level: MENENGAH
	exerciseIntermediate := models.ReadingExerciseTemplate{
		Title: "Latihan Membaca - Menengah",
		Level: "menengah",
	}
	if err := db.Create(&exerciseIntermediate).Error; err != nil {
		return err
	}

	intermediateItems := []models.ExerciseItem{
		{
			ExerciseID: exerciseIntermediate.ID,
			ItemNumber: 1,
			TargetText: "Kemarin saya pergi ke sekolah dengan teman-teman",
		},
		{
			ExerciseID: exerciseIntermediate.ID,
			ItemNumber: 2,
			TargetText: "Saya memiliki hobi membaca buku cerita yang menarik",
		},
		{
			ExerciseID: exerciseIntermediate.ID,
			ItemNumber: 3,
			TargetText: "Pada musim panas, keluarga saya liburan ke pantai",
		},
		{
			ExerciseID: exerciseIntermediate.ID,
			ItemNumber: 4,
			TargetText: "Ibu saya sedang memasak makanan lezat di dapur",
		},
		{
			ExerciseID: exerciseIntermediate.ID,
			ItemNumber: 5,
			TargetText: "Saya ingin belajar lebih giat untuk mendapatkan nilai bagus",
		},
	}
	for _, item := range intermediateItems {
		if err := db.Create(&item).Error; err != nil {
			return err
		}
	}

	// Exercise Level: LANJUT
	exerciseAdvanced := models.ReadingExerciseTemplate{
		Title: "Latihan Membaca - Lanjut",
		Level: "lanjut",
	}
	if err := db.Create(&exerciseAdvanced).Error; err != nil {
		return err
	}

	advancedItems := []models.ExerciseItem{
		{
			ExerciseID: exerciseAdvanced.ID,
			ItemNumber: 1,
			TargetText: "Pendidikan adalah fondasi penting untuk membangun masa depan yang cerah",
		},
		{
			ExerciseID: exerciseAdvanced.ID,
			ItemNumber: 2,
			TargetText: "Teknologi modern telah mengubah cara kita berkomunikasi dan bekerja",
		},
		{
			ExerciseID: exerciseAdvanced.ID,
			ItemNumber: 3,
			TargetText: "Kita harus menjaga lingkungan untuk generasi mendatang",
		},
		{
			ExerciseID: exerciseAdvanced.ID,
			ItemNumber: 4,
			TargetText: "Kesehatan mental sama pentingnya dengan kesehatan fisik",
		},
		{
			ExerciseID: exerciseAdvanced.ID,
			ItemNumber: 5,
			TargetText: "Kerja sama tim dan kolaborasi adalah kunci kesuksesan dalam proyek besar",
		},
	}
	for _, item := range advancedItems {
		if err := db.Create(&item).Error; err != nil {
			return err
		}
	}

	return nil
}
