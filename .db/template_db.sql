CREATE TABLE
    IF NOT EXISTS specialities (
        speciality_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        speciality_name TEXT NOT NULL UNIQUE,
        speciality_duration SMALLINT NOT NULL CHECK (speciality_duration > 0)
    );