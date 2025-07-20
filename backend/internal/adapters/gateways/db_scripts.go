package gateways

const createSpecialitySQL = `
INSERT INTO specialities (speciality_id, speciality_name, speciality_duration)
VALUES ($1, $2, $3)
`
const findAllSpecialitiesSQL = `SELECT speciality_id, speciality_name, speciality_duration FROM specialities`
