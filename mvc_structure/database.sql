CREATE TABLE department (
    dept_no integer PRIMARY KEY,
    dept_name varchar(255)
)
INSERT INTO department
VALUES(1, 'it')
INSERT INTO department
VALUES(2, 'accounting')
INSERT INTO department
VALUES(3, 'marketing') CREATE TABLE employee (
        emp_no integer PRIMARY KEY,
        first_name varchar(255),
        last_name varchar(255),
        education varchar(255),
        dept_no integer,
        CONSTRAINT fk_dept_no FOREIGN KEY(dept_no) REFERENCES department(dept_no)
    )