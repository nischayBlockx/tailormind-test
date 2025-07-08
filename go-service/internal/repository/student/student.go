package student

import (
	"database/sql"
	"go-service/internal/models"
	"go-service/pkg/storage"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(store *storage.Store) *StudentRepository {
	return &StudentRepository{db: store.Postgre}
}

func (r *StudentRepository) GetByID(id string) (*models.StudentDetails, error) {
	query := `
		SELECT
			id, name, email, phone, gender, dob, class_name, section_name, roll,
			father_name, father_phone, mother_name, mother_phone,
			guardian_name, guardian_phone, relation_of_guardian,
			current_address, permanent_address, admission_dt, reporter_name
		FROM user_profiles
		WHERE id = $1`

	row := r.db.QueryRow(query, id)
	var s models.StudentDetails
	err := row.Scan(
		&s.ID, &s.Name, &s.Email,
		&s.Phone, &s.Gender, &s.DOB, &s.Class, &s.Section, &s.Roll,
		&s.FatherName, &s.FatherPhone, &s.MotherName, &s.MotherPhone,
		&s.GuardianName, &s.GuardianPhone, &s.RelationOfGuardian,
		&s.CurrentAddress, &s.PermanentAddress, &s.AdmissionDate, &s.ReporterName,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepository) Insert(s *models.StudentDetails) error {
	query := `
		INSERT INTO user_profiles (
			id, name, email, phone, gender, dob, class_name, section_name, roll,
			father_name, father_phone, mother_name, mother_phone,
			guardian_name, guardian_phone, relation_of_guardian,
			current_address, permanent_address, admission_dt, reporter_name
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9,
			$10, $11, $12, $13,
			$14, $15, $16,
			$17, $18, $19, $20
		)`

	_, err := r.db.Exec(query,
		s.ID, s.Name, s.Email,
		s.Phone, s.Gender, s.DOB, s.Class, s.Section, s.Roll,
		s.FatherName, s.FatherPhone, s.MotherName, s.MotherPhone,
		s.GuardianName, s.GuardianPhone, s.RelationOfGuardian,
		s.CurrentAddress, s.PermanentAddress, s.AdmissionDate, s.ReporterName,
	)
	return err
}

func (r *StudentRepository) Update(s *models.StudentDetails) error {
	query := `
		UPDATE user_profiles SET
			name = $2, email = $3, phone = $4, gender = $5, dob = $6,
			class_name = $7, section_name = $8, roll = $9,
			father_name = $10, father_phone = $11, mother_name = $12, mother_phone = $13,
			guardian_name = $14, guardian_phone = $15, relation_of_guardian = $16,
			current_address = $17, permanent_address = $18, admission_dt = $19, reporter_name = $20
		WHERE id = $1`

	_, err := r.db.Exec(query,
		s.ID, s.Name, s.Email,
		s.Phone, s.Gender, s.DOB, s.Class, s.Section, s.Roll,
		s.FatherName, s.FatherPhone, s.MotherName, s.MotherPhone,
		s.GuardianName, s.GuardianPhone, s.RelationOfGuardian,
		s.CurrentAddress, s.PermanentAddress, s.AdmissionDate, s.ReporterName,
	)
	return err
}

func (r *StudentRepository) Delete(id string) error {
	query := `DELETE FROM user_profiles WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
