CREATE DATABASE dicdoc_service;

\connect "dicdoc_service";

CREATE TABLE IF NOT EXISTS dicdoc_service.public.doctors
(
    id             SERIAL PRIMARY KEY,
    fio            TEXT NOT NULL,
    phoneNumber    TEXT NOT NULL,
    email          TEXT NOT NULL,
    specialization TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS dicdoc_service.public.patients
(
    id          SERIAL PRIMARY KEY,
    fio         TEXT NOT NULL,
    phoneNumber TEXT NOT NULL,
    email       TEXT NOT NULL,
    insurance   TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS dicdoc_service.public.branches
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    address     TEXT NOT NULL,
    phoneNumber TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS dicdoc_service.public.offices
(
    id       SERIAL PRIMARY KEY,
    number   INT NOT NULL,
    floor    INT NOT NULL,
    branchID INT NOT NULL,
    FOREIGN KEY (branchID) REFERENCES branches (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dicdoc_service.public.timetable
(
    id       SERIAL,
    doctorID INT NOT NULL,
    officeID INT NOT NULL,
    workDays INT NOT NULL,
    FOREIGN KEY (doctorID) REFERENCES doctors (id) ON DELETE CASCADE,
    FOREIGN KEY (officeID) REFERENCES offices (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dicdoc_service.public.appointments
(
    id        SERIAL PRIMARY KEY,
    doctorID  INT       NOT NULL,
    patientID INT       NOT NULL,
    datetime  TIMESTAMP NOT NULL,
    FOREIGN KEY (doctorID) REFERENCES doctors (id) ON DELETE CASCADE,
    FOREIGN KEY (patientID) REFERENCES patients (id) ON DELETE CASCADE
);

CREATE INDEX doc_pat_idx ON appointments (doctorID, patientID);

CREATE TABLE IF NOT EXISTS dicdoc_service.public.users
(
    id        SERIAL PRIMARY KEY,
    login     TEXT NOT NULL,
    password  TEXT NOT NULL,
    role      INT  NOT NULL,
    patientID INT,
    doctorID  INT,
    FOREIGN KEY (patientID) REFERENCES patients (id) ON DELETE CASCADE,
    FOREIGN KEY (doctorID) REFERENCES doctors (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dicdoc_service.public.medical_histories
(
    id               SERIAL PRIMARY KEY,
    chronic_diseases text,
    allergies        text,
    blood_type       text,
    vaccination      text,
    patientID        INT NOT NULL UNIQUE,
    FOREIGN KEY (patientID) REFERENCES patients (id) ON DELETE CASCADE
);


CREATE OR REPLACE PROCEDURE delete_appointment(IN appointment_id INT)
    LANGUAGE plpgsql
AS
$$
BEGIN
    DELETE FROM appointments WHERE id = appointment_id;
END;
$$;

CREATE OR REPLACE FUNCTION add_appointment(
    p_doctorID INT,
    p_patientID INT,
    p_datetime TIMESTAMP
)
    RETURNS INT
    LANGUAGE plpgsql
AS
$$
DECLARE
    v_doctorExists  BOOLEAN;
    v_patientExists BOOLEAN;
    v_appointmentID INT;
BEGIN
    SELECT EXISTS(SELECT 1 FROM doctors WHERE id = p_doctorID) INTO v_doctorExists;
    IF NOT v_doctorExists THEN
        RAISE EXCEPTION 'Doctor with ID % does not exist', p_doctorID;
    END IF;

    SELECT EXISTS(SELECT 1 FROM patients WHERE id = p_patientID) INTO v_patientExists;
    IF NOT v_patientExists THEN
        RAISE EXCEPTION 'Patient with ID % does not exist', p_patientID;
    END IF;

    INSERT INTO appointments (doctorID, patientID, datetime)
    VALUES (p_doctorID, p_patientID, p_datetime)
    RETURNING id INTO v_appointmentID;

    RETURN v_appointmentID;
END;
$$;

-- Roles

-- админ
CREATE ROLE admin WITH
    SUPERUSER
    NOINHERIT
    NOREPLICATION
    NOBYPASSRLS
    LOGIN
    PASSWORD 'admin';

GRANT ALL PRIVILEGES
    ON ALL TABLES IN SCHEMA public
    TO admin;

-- врач
CREATE ROLE doctor WITH
    NOSUPERUSER
    NOINHERIT
    NOCREATEROLE
    NOCREATEDB
    NOBYPASSRLS
    NOREPLICATION
    LOGIN
    PASSWORD 'doc';

GRANT SELECT
    ON TABLE public.appointments TO doctor;
GRANT SELECT, UPDATE, INSERT
    ON TABLE public.medical_histories TO doctor;
GRANT SELECT
    ON TABLE public.patients TO doctor;

-- пациент
CREATE ROLE patient WITH
    NOSUPERUSER
    NOINHERIT
    NOCREATEROLE
    NOCREATEDB
    NOBYPASSRLS
    NOREPLICATION
    LOGIN
    PASSWORD 'patient';

GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLE public.appointments TO patient;
GRANT SELECT
    ON TABLE public.doctors TO patient;
GRANT SELECT
    ON TABLE public.timetable TO patient;



INSERT INTO dicdoc_service.public.doctors
VALUES (100, 'fio1', 'phone1', 'email1', 'Терапевт'),
       (200, 'fio2', 'phone2', 'email2', 'Хирург');

INSERT INTO dicdoc_service.public.patients
VALUES (100, 'fio1', 'phone1', 'email1', 'ins1'),
       (200, 'fio2', 'phone2', 'email2', 'ins2');

INSERT INTO dicdoc_service.public.branches
VALUES (100, 'name1', 'address1', 'phone1');

INSERT INTO dicdoc_service.public.offices
VALUES (100, 250, 2, 100);

INSERT INTO dicdoc_service.public.timetable
VALUES (1, 100, 100, 1),
       (2, 100, 100, 2),
       (3, 100, 100, 3),

       (4, 200, 100, 4),
       (5, 200, 100, 5),
       (6, 200, 100, 6);

INSERT INTO dicdoc_service.public.appointments
VALUES (100, 100, 200, '2024-12-03 12:00'),
       (200, 200, 200, '2024-12-06 13:00'),
       (300, 200, 200, '2024-12-05 14:00'),
       (400, 100, 200, '2024-12-04 15:00');

INSERT INTO dicdoc_service.public.users
VALUES (100, 'log1', 'pass1', 0, null, null),
       (200, 'log2', 'pass2', 1, 100, null),
       (300, 'log3', 'pass3', 2, null, 100),
       (400, 'log3', 'pass3', 2, null, 200),
       (500, 'log4', 'pass4', 1, 200, null);

INSERT INTO dicdoc_service.public.medical_histories
VALUES (100, 'Нет', 'Нет', '2+', 'Нет', 100),
       (200, 'Нет', 'Цветение, Пыль', '1-', 'COVID', 200);

