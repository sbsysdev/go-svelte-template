CREATE TABLE
    IF NOT EXISTS specialties (
        specialty_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        specialty_name TEXT NOT NULL UNIQUE,
        specialty_duration SMALLINT NOT NULL CHECK (specialty_duration > 0)
    );

CREATE TABLE
    IF NOT EXISTS doctors (
        doctor_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        doctor_name TEXT NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS doctor_specialties (
        doctor_id UUID REFERENCES doctors (doctor_id) ON DELETE CASCADE ON UPDATE CASCADE,
        specialty_id UUID REFERENCES specialties (specialty_id) ON DELETE CASCADE ON UPDATE CASCADE
    );

SELECT
    d.doctor_id,
    d.doctor_name,
    json_agg (
        json_build_object (
            'specialty_id',
            s.specialty_id,
            'specialty_name',
            s.specialty_name,
            'specialty_duration',
            s.specialty_duration
        )
    ) AS doctor_specialties
FROM
    doctors d
    LEFT JOIN doctor_specialties ds ON ds.doctor_id = d.doctor_id
    LEFT JOIN specialties s ON s.specialty_id = ds.specialty_id
WHERE
    d.doctor_id = ''
GROUP BY
    d.doctor_id,
    d.doctor_name;

CREATE TABLE
    IF NOT EXISTS patients (
        patient_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        patient_name TEXT NOT NULL,
        patient_birth DATE NOT NULL
    );

DO $$
BEGIN
IF NOT EXISTS (
    SELECT
    FROM
        pg_type
    WHERE
        typname LIKE 'appointment_state'
) THEN CREATE TYPE appointment_state AS ENUM ('scheduled', 'completed', 'cancelled');
END IF;
END $$;

CREATE TABLE
    IF NOT EXISTS appointments (
        appointment_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        patient_id UUID REFERENCES patients (patient_id) ON DELETE CASCADE ON UPDATE CASCADE,
        doctor_id UUID REFERENCES doctors (doctor_id) ON DELETE CASCADE ON UPDATE CASCADE,
        specialty_id UUID REFERENCES specialties (specialty_id) ON DELETE CASCADE ON UPDATE CASCADE,
        appointment_date TIMESTAMPTZ NOT NULL CHECK (appointment_date > NOW ()),
        appointment_state appointment_state NOT NULL DEFAULT 'scheduled'
    );