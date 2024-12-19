package controllers

import (
	"database/sql"

	"fmt"

	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"Doctor-Appointment-Project/models"

	helper "Doctor-Appointment-Project/helper"
)

func Add_patient() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("add patient")

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Patient

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		c.IndentedJSON(http.StatusCreated, data)

		query_data := fmt.Sprintf(`INSERT INTO patient(Name,Age,Gender,Address,City,Phone,Disease,Selected_Specialisation,Patient_history) VALUES('%s',%d,'%s','%s','%s','%s','%s','%s','%s')`, data.Name, data.Age, data.Gender, data.Address, data.City, data.Phone, data.Disease, data.Selected_specialisation, data.Patient_history)

		fmt.Println(query_data)

		//insert data

		insert, err := db.Query(query_data)

		if err != nil {

			log.Panic(err.Error())

		}

		defer insert.Close()

	}
}
func Get_patient_details() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var mob models.Patient
		err = c.BindJSON(&mob)
		if err != nil {
			return
		}
		get_detail := fmt.Sprintf("SELECT * FROM Patient WHERE Phone = '%s'", mob.Phone)
		detail, err := db.Query(get_detail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer detail.Close()

		var output interface{}
		for detail.Next() {

			var ID int
			var Name string
			var Age int
			var Gender string
			var Address string
			var City string
			var Phone string
			var Disease string
			var Selected_specialisation string
			var Patient_history string
			err = detail.Scan(&ID, &Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Selected_specialisation, &Patient_history)

			if err != nil {
				log.Panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'  %d  '%s'  '%s'  '%s'  '%s' '%s' %s' '%s'  ", ID, Name, Age, Gender, Address, City, Phone, Disease, Selected_specialisation, Patient_history)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Patient details": output})

		}

		c.JSON(http.StatusOK, gin.H{"message": "Your details"})

	}
}

// func Getpatient() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
// 		if err != nil {

// 			log.Fatal(err)

// 		}
// 		fmt.Println("Connection Created")
// 		results, err := db.Query("SELECT * FROM Patient")
// 		fmt.Println("Quary exicuted")
//

// 		if err != nil {

// 			log.Panic(err.Error())

// 		}

// 		defer results.Close()

// 		var output interface{}

// 		for results.Next() {

// 			var ID int
// 			var Name string
// 			var Age int
// 			var Gender string
// 			var Address string
// 			var City string
// 			var Phone string
// 			var Disease string
// 			var Selected_specialisation string
// 			var Patient_history string
// 			err = results.Scan(&ID, &Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Selected_specialisation, &Patient_history)

// 			if err != nil {

// 				log.Panic(err.Error())

// 			}

// 			output = fmt.Sprintf("%d  '%s'  %d  '%s'  '%s'  '%s'  '%s' '%s' %s' '%s'  ", ID, Name, Age, Gender, Address, City, Phone, Disease, Selected_specialisation, Patient_history)

// 			fmt.Println(output)

// 			c.JSON(http.StatusOK, gin.H{"Data": output})

// 		}

// 	}
// }

func Delete_patient() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Patient

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		delete_query1 := fmt.Sprintf("DELETE FROM Doctor_feedback WHERE Patient_id = %d", data.ID)
		delete1, err := db.Query(delete_query1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer delete1.Close()

		delete_query2 := fmt.Sprintf("DELETE FROM Doctor_appointment WHERE Patient_id = %d", data.ID)
		delete2, err := db.Query(delete_query2)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer delete2.Close()

		delete_query3 := fmt.Sprintf("DELETE FROM Prescription WHERE Patient_id = %d", data.ID)
		delete3, err := db.Query(delete_query3)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer delete3.Close()

		delete_query4 := fmt.Sprintf("DELETE FROM order_medicines WHERE Patient_id = %d", data.ID)
		delete4, err := db.Query(delete_query4)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer delete4.Close()

		delete_query := fmt.Sprintf("DELETE FROM patient WHERE ID= %d", data.ID)

		delete, err := db.Query(delete_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Patient Deleted successfully"})

	}
}

func Get_docter() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("connection not created")
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		fmt.Println("connection is created")
		if err != nil {

			log.Fatal(err)

		}
		fmt.Println("Connection Created")
		results, err := db.Query("SELECT * FROM Doctor")
		fmt.Println("Quary exicuted")

		if err != nil {

			panic(err.Error())

		}

		defer results.Close()

		var output interface{}

		for results.Next() {

			var ID int

			var Name string

			var Gender string

			var Address string

			var City string

			var Phone string

			var Specialisation string

			var Opening_time string

			var Closing_time string

			var Availability_Time string
			var Availability string
			var Available_for_home_visit string
			var Available_for_online_consultancy string
			var Fees float64

			err = results.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Opening_time, &Closing_time, &Availability_Time, &Availability, &Available_for_home_visit, &Available_for_online_consultancy, &Fees)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s' '%s'  '%s' '%s''%s' %f", ID, Name, Gender, Address, City, Phone, Specialisation, Opening_time, Closing_time, Availability_Time, Availability, Available_for_home_visit, Available_for_online_consultancy, Fees)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Data": output})
		}

	}
}

func GetDoctorByLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("connection not created")
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		fmt.Println("connection is created")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection Created")

		var data models.Doctor
		err = c.BindJSON(&data)
		if err != nil {
			return
		}

		query_data := fmt.Sprintf("SELECT * FROM Doctor WHERE City='%s' AND Specialisation='%s'", data.City, data.Specialisation)
		fmt.Println(query_data)

		results, err := db.Query(query_data)
		fmt.Println("Quary exicuted")

		if err != nil {
			panic(err.Error())
		}

		defer results.Close()
		var output interface{}
		for results.Next() {

			var ID int

			var Name string

			var Gender string

			var Address string

			var City string

			var Phone string

			var Specialisation string

			var Opening_time string

			var Closing_time string

			var Availability_Time string

			var Availability string

			var Available_for_home_visit string

			var Available_for_online_consultancy string

			var Fees float64

			err = results.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Opening_time, &Closing_time, &Availability_Time, &Availability, &Available_for_home_visit, &Available_for_online_consultancy, &Fees)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s' '%s'  '%s' '%s''%s' %f", ID, Name, Gender, Address, City, Phone, Specialisation, Opening_time, Closing_time, Availability_Time, Availability, Available_for_home_visit, Available_for_online_consultancy, Fees)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}

	}
}

func Book_doctor_appointment() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var booking_data models.Doctor_appointment

		err = c.BindJSON(&booking_data)
		if err != nil {
			log.Fatal(err)
		}

		get_booking_time := fmt.Sprintf("SELECT Availability_time,Closing_time FROM Doctor WHERE id = %d", booking_data.Doctor_id)
		doctor_result, err := db.Query(get_booking_time)
		// doctor_result,err := db.Exec(get_booking_time)
		if err != nil {
			c.JSON(404, gin.H{"error": "Doctor not found"})
			return
		}

		var people []models.TimeStr

		for doctor_result.Next() {
			var p models.TimeStr
			if err := doctor_result.Scan(&p.Availability_time, &p.Closing_time); err != nil {
				log.Fatal(err)
			}
			people = append(people, p)
		}

		if err := doctor_result.Err(); err != nil {
			log.Fatal(err)
		}

		var booktime string = people[0].Availability_time
		var Closing_time string = people[0].Closing_time

		if Closing_time == booktime {
			c.JSON(http.StatusOK, gin.H{"message": "No slot available visit again thank you"})

		}

		booking_data.Booking_time = booktime

		query_data := fmt.Sprintf(`INSERT INTO Doctor_appointment (Patient_id,Doctor_id,Booking_time) VALUES(%d,%d,'%s')`, booking_data.Patient_id, booking_data.Doctor_id, booking_data.Booking_time)
		_, err = db.Exec(query_data)
		if err != nil {

			log.Panic(err.Error())

		}
		t1 := helper.Add_time(booktime)

		query_data2 := fmt.Sprintf(`UPDATE Doctor SET Availability_time = '%s' WHERE ID = %d`, t1, booking_data.Doctor_id)

		fmt.Println(query_data2)

		_, err = db.Query(query_data2)
		if err != nil {
			log.Panic(err)
		}

		c.IndentedJSON(http.StatusCreated, booking_data)
		c.JSON(http.StatusOK, gin.H{"message": "Your Appointment successfully Booked"})

	}
}

func Cancel_doctor_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Doctor_appointment

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		c.IndentedJSON(http.StatusCreated, data)

		query_data := fmt.Sprintf("DELETE FROM Doctor_appointment WHERE Bookingid =%d", data.Bookingid)

		_, err = db.Exec(query_data)

		if err != nil {

			log.Panic(err.Error())

		}

		c.JSON(http.StatusOK, gin.H{"message": "Doctor Appointment cancelled successfully"})

	}
}

func Doctor_feedback() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("add feedback")
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var data models.Doctor_feedback
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf(`INSERT INTO Doctor_feedback(Patient_id,Doctor_id,Rating,Feedback_msg) VALUES(%d,%d,%d,'%s')`, data.Patient_id, data.Doctor_id, data.Rating, data.Feedback_msg)
		fmt.Println(query_data)

		//insert data

		insert, err := db.Query(query_data)
		if err != nil {
			log.Panic(err.Error())
		}
		defer insert.Close()
		c.JSON(http.StatusOK, gin.H{"message": "Thanks for giving Feedback"})
	}

}

// method for booking of private nurse

func Book_nurse_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var book_nurse models.Nurse_appointment
		// var nurse_data models.Nurse
		err = c.BindJSON(&book_nurse)
		if err != nil {
			log.Fatal(err)
		}
		get_query := fmt.Sprintf("SELECT Availability FROM Nurse WHERE id = %d", book_nurse.Nurse_id)
		result, err := db.Query(get_query)
		if err != nil {
			c.JSON(404, gin.H{"error": "Nurse not found"})
		}
		defer result.Close()
		var availability string
		for result.Next() {
			if err := result.Scan(&availability); err != nil {
				log.Fatal(err)
			}
		}
		if availability == "NO" {
			c.JSON(404, gin.H{"Messsage": "Nurse is not available for booking"})

		} else {

			query_data := fmt.Sprintf(`INSERT INTO Nurse_appointment (Patient_id,Nurse_id) VALUES(%d,%d)`, book_nurse.Patient_id, book_nurse.Nurse_id)
			_, err = db.Exec(query_data)
			if err != nil {
				log.Panic(err.Error())
			}
			A := "NO"
			query_data2 := fmt.Sprintf(`UPDATE Nurse SET Availability = '%s' WHERE ID = %d`, A, book_nurse.Nurse_id)

			fmt.Println(query_data2)
			_, err = db.Query(query_data2)
			if err != nil {
				log.Panic(err.Error())
			}
			c.IndentedJSON(http.StatusCreated, book_nurse)
			c.JSON(http.StatusOK, gin.H{"message": "Nurse Appointment Booked successfully"})
		}
	}
}

// method to cancel the nurse appointment

func Cancel_nurse_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Nurse_appointment
		err = c.BindJSON(&data)
		if err != nil {
			log.Fatal(err)
		}
		// c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf("DELETE FROM Nurse_appointment WHERE Bookingid =%d", data.Bookingid)

		_, err = db.Exec(query_data)

		if err != nil {
			log.Panic(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"message": "Nurse Appointment cancelled successfully"})

	}
}

// method to give the nurse feedback by patient

func Nurse_feedback() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Nurse_feedback
		err = c.BindJSON(&data)
		if err != nil {
			log.Fatal(err)
		}

		c.IndentedJSON(http.StatusCreated, data)

		query_data := fmt.Sprintf(`INSERT INTO Nurse_feedback(Patient_id,Nurse_id,Rating,Feedback_msg) VALUES(%d,%d,%d,'%s')`, data.Patient_id, data.Nurse_id, data.Rating, data.Feedback_msg)

		fmt.Println(query_data)

		//insert data

		insert, err := db.Query(query_data)

		if err != nil {

			log.Panic(err.Error())

		}

		defer insert.Close()

	}
}

// Prescription - get method

func Get_prescription() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Patient
		err = c.BindJSON(&data)
		if err != nil {
			return
		}

		Get_query := fmt.Sprintf("SELECT Patient.Name,Patient.Age,Patient.Gender,Patient.Address,Patient.City,Patient.Phone,Patient.Disease,Patient.Patient_history,Prescription.Prescription FROM Prescription INNER JOIN Patient ON Prescription.Patient_id = Patient.id where Prescription.Patient_id = %d", data.ID)

		GetResult, err := db.Query(Get_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer GetResult.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Your Appointment"})

		var output interface{}

		for GetResult.Next() {

			var Name string

			var Age int

			var Gender string

			var Address string

			var City string

			var Phone string

			var Disease string

			var Selected_Specialisation string

			var Patient_history string

			var Prescription string

			err = GetResult.Scan(&Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Patient_history, &Prescription)

			if err != nil {

				log.Panic(err.Error())

			}

			output = fmt.Sprintf("'%s' %d  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s','%s", Name, Age, Gender, Address, City, Phone, Disease, Selected_Specialisation, Patient_history, Prescription)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}

	}
}

// Order Medicine - Place Order - Post method

func Order_medicines() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Order_medicines
		var data_prescription models.Prescription
		err = c.BindJSON(&data_prescription)
		if err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusOK, data_prescription)
		query := fmt.Sprintf(`SELECT Patient_id,Doctor_id,Prescription FROM Prescription WHERE Patient_id=%d`, data_prescription.Patient_id)
		result, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer result.Close()

		// var output interface{}

		for result.Next() {

			err = result.Scan(&data.Patient_id, &data.Doctor_id, &data.Prescription)
			if err != nil {
				log.Fatal(err)
			}

		}
		insert_query := fmt.Sprintf(`INSERT INTO Order_medicines (Patient_id,Doctor_id,Prescription) VALUES(%d,%d,'%s')`, data.Patient_id, data.Doctor_id, data.Prescription)

		insert, err := db.Query(insert_query)
		if err != nil {
			log.Fatal(err)
		}
		defer insert.Close()
		c.JSON(http.StatusOK, gin.H{"Message": "Order Placed Successfully! Thank You Visit Again..."})
	}
}

// Order Medicine - Cancel Order - Delete method

func Cancel_ordered_medicines() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Order_medicines
		err = c.BindJSON(&data)
		if err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf("DELETE FROM Order_medocines WHERE id =%d", data.Order_id)

		_, err = db.Exec(query_data)

		if err != nil {
			log.Panic(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"message": " Medicines ordered cancelled successfully"})

	}
}

// get doctor by online consultancy availability

func Get_doctor_by_online_consultancy_availability() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			return
		}
		var data models.Doctor

		err = c.BindJSON(&data)
		if err != nil {
			return
		}

		query_data := fmt.Sprintf("SELECT * FROM Doctor WHERE City='%s' AND Specialisation='%s' AND Available_for_online_consultancy='Yes'", data.City, data.Specialisation)
		fmt.Println(query_data)

		result, err := db.Query(query_data)
		fmt.Println("Quary exicuted")

		if err != nil {
			panic(err.Error())
		}
		defer result.Close()
		fmt.Println(result)

		var output interface{}
		for result.Next() {

			var ID int

			var Name string

			var Gender string

			var Address string

			var City string

			var Phone string

			var Specialisation string

			var Opening_time string

			var Closing_time string

			var Availability_Time string

			var Availability string

			var Available_for_home_visit string

			var Available_for_online_consultancy string

			var Fees float64

			err = result.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Opening_time, &Closing_time, &Availability_Time, &Availability, &Available_for_home_visit, &Available_for_online_consultancy, &Fees)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s' '%s'  '%s' '%s''%s' %f", ID, Name, Gender, Address, City, Phone, Specialisation, Opening_time, Closing_time, Availability_Time, Availability, Available_for_home_visit, Available_for_online_consultancy, Fees)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}
		if output == nil {
			c.JSON(http.StatusOK, gin.H{"Message": "No Doctor Available for online consultancy"})
		}

	}
}

// Online Consultancy - Book Appointment -  Post Method

func Book_online_consultancy_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var booking_data models.Online_consultancy
		var doctor_data models.Doctor
		err = c.BindJSON(&booking_data)
		if err != nil {
			log.Fatal(err)
		}

		get_booking_time := fmt.Sprintf("SELECT Availability_time,Closing_time FROM Doctor WHERE id = %d", doctor_data.ID)
		doctor_result, err := db.Query(get_booking_time)
		// doctor_result,err := db.Exec(get_booking_time)
		if err != nil {
			c.JSON(404, gin.H{"error": "Doctor not found"})
			return
		}

		var people []models.TimeStr

		for doctor_result.Next() {
			var p models.TimeStr
			if err := doctor_result.Scan(&p.Availability_time, &p.Closing_time); err != nil {
				log.Fatal(err)
			}
			people = append(people, p)
		}

		if err := doctor_result.Err(); err != nil {
			log.Fatal(err)
		}

		var booktime string = people[0].Availability_time
		var Closing_time string = people[1].Closing_time

		if Closing_time == booktime {
			c.JSON(http.StatusOK, gin.H{"message": "No slot available visit again thank you"})

		} else {

			booking_data.Booking_time = booktime

			query_data := fmt.Sprintf(`INSERT INTO Online_consultancy (Patient_id,Doctor_id,Booking_time) VALUES(%d,%d,'%s')`, booking_data.Patient_id, booking_data.Doctor_id, booking_data.Booking_time)
			_, err = db.Exec(query_data)
			if err != nil {

				log.Panic(err.Error())

			}
			t1 := helper.Add_time(booktime)

			query_data2 := fmt.Sprintf(`UPDATE Doctor SET Availability_time = '%s' WHERE ID = %d`, t1, booking_data.Doctor_id)

			fmt.Println(query_data2)

			_, err = db.Query(query_data2)
			if err != nil {
				log.Panic(err)
			}

			c.IndentedJSON(http.StatusCreated, booking_data)
			c.JSON(http.StatusOK, gin.H{"message": "Your Appointment successfully Booked"})
		}
	}
}

// Online Consultancy - Cancel Appointment -  Delete Method

func Cancel_online_consultancy_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Online_consultancy
		err = c.BindJSON(&data)
		if err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf("DELETE FROM Home_visit_appointment WHERE id =%d", data.Bookingid)

		_, err = db.Exec(query_data)

		if err != nil {
			log.Panic(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"message": "Doctor's online consultancy appointment cancelled successfully"})

	}
}

// get doctor by home_visit availability

func Get_doctor_by_home_visit_availability() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			return
		}
		var data models.Doctor

		err = c.BindJSON(&data)
		if err != nil {
			return
		}

		query_data := fmt.Sprintf("SELECT * FROM Doctor WHERE City='%s' AND Specialisation='%s' AND Available_for_home_visit='Yes'", data.City, data.Specialisation)
		fmt.Println(query_data)

		result, err := db.Query(query_data)
		fmt.Println("Quary exicuted")

		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		var output interface{}

		for result.Next() {

			var ID int

			var Name string

			var Gender string

			var Address string

			var City string

			var Phone string

			var Specialisation string

			var Opening_time string

			var Closing_time string

			var Availability_Time string

			var Availability string

			var Available_for_home_visit string

			var Available_for_online_consultancy string

			var Fees float64

			err = result.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Opening_time, &Closing_time, &Availability_Time, &Availability, &Available_for_home_visit, &Available_for_online_consultancy, &Fees)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s' '%s'  '%s' '%s''%s' %f", ID, Name, Gender, Address, City, Phone, Specialisation, Opening_time, Closing_time, Availability_Time, Availability, Available_for_home_visit, Available_for_online_consultancy, Fees)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}
		if output == nil {
			c.JSON(http.StatusOK, gin.H{"Message": "No Doctor Available for home visit"})
		}

	}
}

// Home Visit - Book Appointment -  Post Method

func Book_home_visit_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var booking_data models.Home_visit_appointment
		var doctor_data models.Doctor
		err = c.BindJSON(&booking_data)
		if err != nil {
			log.Fatal(err)
		}
		get_booking_time := fmt.Sprintf("SELECT Availability, Available_for_home_visit, FROM Doctor WHERE id=%d", doctor_data.ID)

		result, err := db.Query(get_booking_time)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Doctor who is available for home visit is not found"})
		}
		defer result.Close()
		var people []models.TimeStr

		for result.Next() {
			var p models.TimeStr
			if err := result.Scan(&p.Availability_time, &p.Closing_time); err != nil {
				log.Fatal(err)
			}

			var booking string = people[0].Availability_time
			var available string = people[1].Closing_time

			if available == "No" {
				c.JSON(http.StatusOK, gin.H{"message": "Doctor is not available for home visit"})
			}
			booking_data.Available_for_home_visit = booking
			c.IndentedJSON(http.StatusCreated, booking_data)
		}
	}
}

// Home Visit - Book Appointment -  Delete Method

func Cancel_home_visit_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Home_visit_appointment
		err = c.BindJSON(&data)
		if err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf("DELETE FROM Home_visit_appointment WHERE id =%d", data.Bookingid)

		_, err = db.Exec(query_data)

		if err != nil {
			log.Panic(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"message": "Doctor's home visit appointment cancelled successfully"})

	}
}
