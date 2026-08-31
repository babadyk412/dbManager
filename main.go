package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Структура для ученика
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

func main() {
	// Открываем базу данных
	db, err := sql.Open("sqlite3", "Studentos.db")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}

	defer db.Close()

	for {
		fmt.Println("=== Школьная база данных ===")
		fmt.Println()

		fmt.Println("1 создать")
		fmt.Println("2 изменить")
		var tt int
		fmt.Scan(&tt)
		if tt == 1 {
			var name string
			fmt.Println("введи имя")
			fmt.Scan(&name)

			var age int
			fmt.Println("введи вораст")
			fmt.Scan(&age)

			var grade string
			fmt.Println("введи класс")
			fmt.Scan(&grade)

			addStudent(db, name, age, grade)
		}
		if tt == 2 {

			var Oldname string
			fmt.Println("введи имя кого изменить")
			fmt.Scan(&Oldname)

			var name string
			fmt.Println("введи имя")
			fmt.Scan(&name)

			var age int
			fmt.Println("введи вораст")
			fmt.Scan(&age)

			var grade string
			fmt.Println("введи класс")
			fmt.Scan(&grade)

			izmen(db, Oldname, name, age, grade)
			fmt.Println("данные изменены")
		}
		if tt == 4 {
			fmt.Println("старое")
			var y string
			fmt.Scan(&y)

			fmt.Println("новое")
			var t string
			fmt.Scan(&t)

			nameizmen(db, t, y)
		}
		if tt == 6 {

			students, err := getStudents(db)
			if err != nil {
				fmt.Println(err, "ашибка")
				return
			}
			for _, i := range students {
				fmt.Println("айди=", i.ID, "ima=", i.Name, "vosrast=", i.Age, "klass=", i.Grade)
			}

		}
		if tt == 7 {
			fmt.Println("id")
			var id int
			fmt.Scan(&id)
			r, err := getStudentAidi(db, id)
			if err != nil {
				fmt.Println(err, "ашибка")
				return
			}
			fmt.Println("айди=", r.ID, "ima=", r.Name, "vosrast=", r.Age, "klass=", r.Grade)

		}
	}
}

func addStudent(db *sql.DB, name string, age int, grade string) {
	query := `INSERT INTO students (name, age, grade) VALUES (?, ?, ?)`

	_, err := db.Exec(query, name, age, grade)
	if err != nil {
		fmt.Println("Ошибка при добавлении:", err)
		return
	}

	fmt.Printf("✓ Добавлен ученик: %s\n", name)
}

func izmen(db *sql.DB, Oldname string, name string, age int, grade string) {
	yer := `UPDATE students  SET name=?, age=?, grade=? WHERE name=?`

	_, err := db.Exec(yer, name, age, grade, Oldname)
	if err != nil {
		fmt.Println("fff", err)
		return
	}
	fmt.Println(yer)

}

func nameizmen(db *sql.DB, name string, Oldname string) {
	yer := `UPDATE students  SET name=? where name=? `

	_, err := db.Exec(yer, name, Oldname)
	if err != nil {
		fmt.Println("fff", err)
		return
	}
}

func getStudents(db *sql.DB) ([]Student, error) {
	query := `SELECT * FROM students`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("fff", err)
		return []Student{}, err
	}
	var students []Student

	for rows.Next() {
		var student Student

		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
			&student.Grade,
		)
		if err != nil {
			return []Student{}, err
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return []Student{}, err
	}
	return students, nil
}

func getStudentAidi(db *sql.DB, id int) (Student, error) {
	query := `SELECT * FROM students WHERE id=?`

	row, err := db.Query(query, id)
	if err != nil {
		fmt.Println("fff", err)
		return Student{}, err
	}

	var student Student
	for row.Next() {

		err := row.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
			&student.Grade,
		)
		if err != nil {
			return Student{}, err
		}

	}

	if err := row.Err(); err != nil {
		return Student{}, err
	}

	return student, nil
}
