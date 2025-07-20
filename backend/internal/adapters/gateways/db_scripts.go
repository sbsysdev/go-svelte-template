package gateways

// SpecialtyRepository
const createSpecialtySQL = `
INSERT INTO specialties (specialty_id, specialty_name, specialty_duration) VALUES ($1, $2, $3)
`
const findSpecialtyByIDSQL = `
SELECT specialty_id, specialty_name, specialty_duration FROM specialties WHERE specialty_id = $1
`
const findSpecialtiesByDoctorIDSQL = `
SELECT s.specialty_id, s.specialty_name, s.specialty_duration
FROM specialties s
JOIN doctor_specialties ds ON s.specialty_id = ds.specialty_id
WHERE ds.doctor_id = $1
`
const findAllSpecialtiesSQL = `
SELECT specialty_id, specialty_name, specialty_duration FROM specialties
`

// DoctorRepository
const createDoctorSQL = `
INSERT INTO doctors (doctor_id, doctor_name) VALUES ($1, $2)
`
const createDoctorSpecialtySQL = `
INSERT INTO doctor_specialties (doctor_id, specialty_id) VALUES ($1, $2)
`
const findDoctorByIDSQL = `
SELECT doctor_id, doctor_name FROM doctors WHERE doctor_id = $1
`
const findDoctorsBySpecialtyIDSQL = `
SELECT d.doctor_id, d.doctor_name
FROM doctors d
JOIN doctor_specialties ds ON d.doctor_id = ds.doctor_id
WHERE ds.specialty_id = $1
`
