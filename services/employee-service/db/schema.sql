CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE employees (
    id             BIGSERIAL PRIMARY KEY,
    first_name     VARCHAR,
    last_name      VARCHAR,
    date_of_birth  DATE,
    gender         VARCHAR,
    email          VARCHAR UNIQUE,
    phone_number   VARCHAR,
    address        VARCHAR,
    username       VARCHAR UNIQUE,
    password       VARCHAR,
    position       VARCHAR,
    department     VARCHAR,
    active         BOOLEAN,
    permissions    TEXT[],
    jmbg           VARCHAR(13) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS idx_employees_first_name ON employees (first_name);
CREATE INDEX IF NOT EXISTS idx_employees_last_name  ON employees (last_name);
CREATE INDEX IF NOT EXISTS idx_employees_position   ON employees (position);

INSERT INTO employees (first_name, last_name, date_of_birth, gender, email, phone_number, address, username, password, position, department, active, permissions, jmbg)
SELECT 'Admin', 'Admin', '1990-01-01', 'M', 'admin@exbanka.com', '', '', 'admin', crypt('admin', gen_salt('bf', 10)), 'Administrator', 'IT', true, ARRAY['ADMIN', 'READ', 'WRITE', 'DELETE'], '0000000000001'
WHERE NOT EXISTS (SELECT 1 FROM employees WHERE username = 'admin');

-- Seed test employees used by Cypress e2e tests (celina1)
INSERT INTO employees (first_name, last_name, date_of_birth, gender, email, phone_number, address, username, password, position, department, active, permissions, jmbg)
SELECT 'Marko', 'Markovic', '1995-05-15', 'M', 'marko@banka.rs', '+381601111111', 'Knez Mihailova 1, Beograd', 'marko', crypt('marko123', gen_salt('bf', 10)), 'Agent', 'Retail', true, ARRAY['READ', 'WRITE'], '0000000000002'
WHERE NOT EXISTS (SELECT 1 FROM employees WHERE username = 'marko');

INSERT INTO employees (first_name, last_name, date_of_birth, gender, email, phone_number, address, username, password, position, department, active, permissions, jmbg)
SELECT 'Luka', 'Lukovic', '1993-08-20', 'M', 'luka@banka.rs', '+381602222222', 'Terazije 5, Beograd', 'luka', crypt('luka123', gen_salt('bf', 10)), 'Agent', 'Retail', true, ARRAY['READ'], '0000000000003'
WHERE NOT EXISTS (SELECT 1 FROM employees WHERE username = 'luka');

-- Inactive employees for activation tests (S8/S9) — explicit IDs so auth-service tokens can reference them
INSERT INTO employees (id, first_name, last_name, date_of_birth, gender, email, phone_number, address, username, password, position, department, active, permissions, jmbg)
VALUES (100, 'Cypress', 'Activate', '1990-01-01', 'M', 'cypress.activate@banka.rs', '', '', 'cypress_activate', NULL, 'Agent', 'IT', false, ARRAY['READ'], '0000000000004')
ON CONFLICT (id) DO NOTHING;

INSERT INTO employees (id, first_name, last_name, date_of_birth, gender, email, phone_number, address, username, password, position, department, active, permissions, jmbg)
VALUES (101, 'Cypress', 'Expired', '1990-01-01', 'M', 'cypress.expired@banka.rs', '', '', 'cypress_expired', NULL, 'Agent', 'IT', false, ARRAY['READ'], '0000000000005')
ON CONFLICT (id) DO NOTHING;

INSERT INTO employees (id, first_name, last_name, date_of_birth, gender, email, phone_number, address, username, password, position, department, active, permissions, jmbg)
VALUES (102, 'Petar', 'Petrovic', '1991-03-10', 'M', 'petar.petrovic@banka.rs', '', '', 'petar', NULL, 'Agent', 'Retail', false, ARRAY[]::TEXT[], '0000000000006')
ON CONFLICT (id) DO NOTHING;

SELECT setval('employees_id_seq', GREATEST((SELECT MAX(id) FROM employees), 102));
