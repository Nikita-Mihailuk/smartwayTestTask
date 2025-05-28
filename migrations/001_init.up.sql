CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS passports (
    id SERIAL PRIMARY KEY,
    type VARCHAR(4) NOT NULL,
    number VARCHAR(6) NOT NULL
);

CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    company_id INT REFERENCES companies(id)
);

CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    company_id INT REFERENCES companies(id),
    passport_id INT UNIQUE REFERENCES passports(id),
    department_id INT REFERENCES departments(id)
);

CREATE INDEX ON employees(company_id);
CREATE INDEX ON employees(department_id);
CREATE INDEX ON passports(number,  type);

INSERT INTO companies (name) VALUES
('TestCompany'),
('TestCompany2');

INSERT INTO passports (type, number) VALUES
('1111', '123456'),
('2222', '654321'),
('3333', '111111'),
('4444', '222222');

INSERT INTO departments (name, phone, company_id) VALUES
('IT', '+7 900 123 33 11', 1),
('HR', '8-999-888-77-12', 2),
('Support', '89001003243', 1);

INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) VALUES
('Liam',     'Anderson', '+7-202-555-0140', 1, 1, 1),
('Sophia',   'Martinez', '+7-202-555-0172', 2, 2, 2),
('Noah',     'Clark',    '+7-202-555-0114', 1, 3, 3),
('Isabella', 'Lewis',    '+7-202-555-0183', 2, 4, 2);

