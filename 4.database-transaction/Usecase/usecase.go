package usecase

import (
	"db/models"
	"db/repository"
)

var (
	err     error
	avgMark = 250
)

func AddStudentData(userID string, input models.StudentData) error {
	tx, err := repository.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()

	total := input.EnglishMark + input.MathsMark

	schoolData, err := repository.FindSchool(input.StudentID)
	if err != nil {
		tx.Rollback()
		return err
	}

	admissionNo, err := repository.FindAdmissionNo(schoolData.SchoolID, schoolData.StudentID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = repository.AddAMarkInSchool(repository.SchoolMarkSheet{
		StudentID:   input.StudentID,
		StudentName: input.FirstName,
		AdmissionNo: admissionNo,
		SchoolID:    schoolData.SchoolID,
		TotalMark:   total,
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = repository.AddAMarkInState(repository.StateMarkSheet{
		StudentID: input.StudentID,
		SchoolID:  schoolData.SchoolID,
		StateID:   schoolData.StateID,
		TotalMark: total,
	}); err != nil {
		tx.Rollback()
		return err
	}

	schoolTopScore, err := repository.FindTopScoreBySchool(schoolData.SchoolID)
	if err != nil {
		tx.Rollback()
		return err
	}

	stateTopScore, err := repository.FindTopScoreByState(schoolData.StateID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if total > schoolTopScore {
		if err = repository.UpdateSchoolRank(repository.SchoolRank{
			SchoolID:  schoolData.SchoolID,
			StudentID: input.StudentID,
			TotalMark: total,
		}); err != nil {
			tx.Rollback()
			return err
		}
	}

	if total > stateTopScore {
		if err = repository.UpdateSateRank(repository.StateRank{
			SchoolID:  schoolData.SchoolID,
			StudentID: input.StudentID,
			StateID:   schoolData.StateID,
			TotalMark: total,
		}); err != nil {
			tx.Rollback()
			return err
		}
	}

	if total < avgMark {
		if err = repository.UpdateProhbationList(repository.Prohbation{
			StudentID: input.StudentID,
			SchoolID:  schoolData.SchoolID,
		}); err != nil {
			tx.Rollback()
			return err
		}
	}

	studentInfo := repository.StudentInfo{
		StudentID: input.StudentID,
		SchoolID:  schoolData.SchoolID,
	}

	if total < avgMark {
		studentInfo.Prohbation = true
	} else {
		studentInfo.Prohbation = false
	}

	if err = repository.UpdateStudentInfo(studentInfo); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
