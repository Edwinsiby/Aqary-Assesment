package repository

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func FindSchool(studentID string) (School, error) {
	var school School
	if err := db.Where("student_id = ?", studentID).First(&school).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return School{}, errors.New("School not found")
		}
		return School{}, err
	}

	return school, nil
}

func FindAdmissionNo(studentID, schoolID string) (int, error) {
	var admissionNo int
	if err := db.Where("student_id = ? AND school_id = ?", studentID, schoolID).Select("admission_no").First(&admissionNo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, errors.New("Admission number not found")
		}
		return 0, err
	}

	return admissionNo, nil
}

func AddAMarkInSchool(markSheet SchoolMarkSheet) error {
	mark := SchoolMarkSheet{
		StudentID:   markSheet.StudentID,
		StudentName: markSheet.StudentName,
		SchoolID:    markSheet.SchoolID,
		TotalMark:   markSheet.TotalMark,
	}
	if err := db.Create(&mark).Error; err != nil {
		return err
	}

	return nil
}

func AddAMarkInState(markSheet StateMarkSheet) error {
	mark := StateMarkSheet{
		StudentID: markSheet.StudentID,
		SchoolID:  markSheet.SchoolID,
		TotalMark: markSheet.TotalMark,
	}
	if err := db.Create(&mark).Error; err != nil {
		return err
	}

	return nil
}

func FindTopScoreBySchool(schoolID string) (int, error) {
	var topScore int
	err = db.Table("schoolmark").Where("school_id = ?", schoolID).Select("MAX(total_mark)").Scan(&topScore).Error
	if err != nil {
		return 0, err
	}

	return topScore, nil
}

func FindTopScoreByState(state string) (int, error) {
	var topScore int
	err = db.Table("schoolmark").Where("state = ?", state).Select("MAX(total_mark)").Scan(&topScore).Error
	if err != nil {
		return 0, err
	}

	return topScore, nil
}

func UpdateSchoolRank(rank SchoolRank) error {
	err = db.Table("schoolrank").
		Where("school_id = ?", rank.SchoolID).
		Updates(map[string]interface{}{"total_mark,": rank.TotalMark, "student_id": rank.StudentID}).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateSateRank(rank StateRank) error {
	err = db.Table("staterank").
		Where("state_id = ? ", rank.StateID).
		Updates(map[string]interface{}{"total_mark": rank.TotalMark, "student_id": rank.StudentID}).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateProhbationList(prohbation Prohbation) error {
	err = db.Create(&prohbation).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateStudentInfo(info StudentInfo) error {
	var existingInfo StudentInfo
	if err := db.Where("student_id = ?", info.StudentID).First(&existingInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = db.Create(&info).Error
		} else {
			return err
		}
	} else {
		err = db.Model(&existingInfo).Updates(info).Error
	}

	if err != nil {
		return err
	}

	return nil
}
