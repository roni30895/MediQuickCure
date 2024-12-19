package main

import (
	routes "Doctor-Appointment-Project/routes"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}
	Connect()
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.DoctorRoutes(router)
	routes.PatientRoutes(router)
	routes.LabtRoutes(router)
	routes.NurseRoutes(router)
	routes.MedicalRoutes(router)

	router.Run(":" + port)
}

func Connect() {

	// Create Database
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/")
	if err != nil {

		panic(err.Error())

	}
	defer db.Close()

	//  Create Database

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS das")

	if err != nil {

		panic(err.Error())

	}

	// Make Database Connection
	db, err = sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Connected to MySQL database!")

	// Create Patient table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Patient(ID INT NOT NULL AUTO_INCREMENT, Name VARCHAR(30),Age INT,Gender VARCHAR(10),Address VARCHAR(50), City VARCHAR(20),Phone VARCHAR(15),Disease VARCHAR(25),Selected_Specialisation VARCHAR(20),Patient_history VARCHAR(250), PRIMARY KEY (ID) );")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Patient Table Created")

	// Create Docter table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Doctor(ID INT NOT NULL AUTO_INCREMENT,Name VARCHAR(30),Gender VARCHAR(10),Address VARCHAR(50), City VARCHAR(20),Phone VARCHAR(15),Specialisation VARCHAR(20),Opening_time VARCHAR(10),Closing_time VARCHAR(10),Availability_time VARCHAR(30),Availability VARCHAR(10),Available_for_home_visit VARCHAR(4),Available_for_online_consultancy VARCHAR(4),Fees INT ,PRIMARY KEY (ID) );")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Docter Table Created")

	// Create Doctor Appointment table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Doctor_appointment(Bookingid INT NOT NULL AUTO_INCREMENT,Patient_id INT,Doctor_id INT,Booking_time VARCHAR(10),PRIMARY KEY (Bookingid),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Doctor_id) REFERENCES Doctor(ID));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Doctor Appointment Table Created")

	// Create Doctor Feedback table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Doctor_feedback(ID INT NOT NULL AUTO_INCREMENT,Patient_id INT,Doctor_id INT,Rating INT,Feedback_msg VARCHAR(255),PRIMARY KEY (ID),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Doctor_id) REFERENCES Doctor(ID));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Doctor_feedback table created")

	// Create Nursing table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Nurse(ID INT NOT NULL AUTO_INCREMENT,Name VARCHAR(30),Gender VARCHAR(10),Address VARCHAR(50),City VARCHAR(20),Phone VARCHAR(15),Specialisation VARCHAR(20),Start_time VARCHAR(10),End_time VARCHAR(10),Charge_per_day INT,Availability VARCHAR(30),PRIMARY KEY (ID));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Nursing Table Created")

	// Create Nurse Appointment table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Nurse_appointment(Bookingid INT NOT NULL AUTO_INCREMENT,Patient_id INT,Nurse_id INT,PRIMARY KEY (Bookingid),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Nurse_id) REFERENCES Nurse(ID));")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Nurse Appointment Table Created")

	// Nurse feedback table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Nurse_feedback(ID INT NOT NULL AUTO_INCREMENT,Patient_id INT,Nurse_id INT,Rating INT,Feedback_msg VARCHAR(255),PRIMARY KEY (ID),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Nurse_id) REFERENCES Nurse(ID));")

	if err != nil {

		panic(err.Error())

	}

	fmt.Println("Nurse_feedback table created")

	// Prescription table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Prescription(prescriptionid INT NOT NULL AUTO_INCREMENT,Patient_id INT,Doctor_id INT,Prescription VARCHAR(255),PRIMARY KEY (prescriptionid),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Doctor_id) REFERENCES Doctor(ID));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Prescription Table Created")

	// Order table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Order_medicines(Orderid INT NOT NULL AUTO_INCREMENT,Patient_id INT,Doctor_id INT,Prescription VARCHAR(255),Prescriptionid INT,PRIMARY KEY (Orderid),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Doctor_id) REFERENCES Doctor(ID),FOREIGN KEY (prescriptionid) REFERENCES Prescription(prescriptionid));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("OrderTable Created")

	// Create Online Consultancy Appointment table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Online_consultancy(OnlineAppointmentBookingid INT NOT NULL AUTO_INCREMENT,Patient_id INT,Doctor_id INT,Booking_time VARCHAR(10),PRIMARY KEY (OnlineAppointmentBookingid),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Doctor_id) REFERENCES Doctor(ID));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println(" Online Appointment Table Created")

	// Create Home Visit Appointment table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Home_visit_appointment(HomeAppointmentBookingid INT NOT NULL AUTO_INCREMENT,Patient_id INT,Doctor_id INT,Booking_time VARCHAR(10),PRIMARY KEY (HomeAppointmentBookingid),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Doctor_id) REFERENCES Doctor(ID));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println(" Home Appointment Table Created")

	// Create test Appointment table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS TestAppointment(TestAppointmentBookingid INT NOT NULL AUTO_INCREMENT,Patient_id INT,Doctor_id INT,Labid INT,Test_Name VARCHAR(255),Booking_time VARCHAR(10),PRIMARY KEY (TestAppointmentBookingid),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Doctor_id) REFERENCES Doctor(ID));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println(" test Appointment Table Created")

	// Create lab table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS lab(Labid INT NOT NULL AUTO_INCREMENT,Lab_Name VARCHAR(50),Lab_Operator_Name VARCHAR(50),Phone VARCHAR(15),Address VARCHAR(255),City VARCHAR(50),Pin_Code VARCHAR(7),Available_test_name VARCHAR(255),Opening_time VARCHAR(10),Closing_time VARCHAR(10),Availability  VARCHAR(10),Availability_time_for_test VARCHAR(30),PRIMARY KEY (Labid));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println(" Lab Detail Table Created")

	// Create Lab_Feedback table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Lab_feedback(ID INT NOT NULL AUTO_INCREMENT,Patient_id INT,Labid INT,Rating INT,feedback_msg VARCHAR(255),PRIMARY KEY (ID),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Labid) REFERENCES Doctor(ID));")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println(" Lab_Feedback Table Created")

}
