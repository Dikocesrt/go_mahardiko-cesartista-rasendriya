package main

import "fmt"

type Customer struct {
    Name    string
    Address string
    Rentals []Rental
}

type Vehicle struct {
    Brand string
    Model string
    Price int
}

type Rental struct {
    Vehicle
    Duration int
}

func main() {
    vehicles := []Vehicle{
        {Brand: "Honda", Model: "Supra", Price: 120000},
        {Brand: "Yamaha", Model: "NMax", Price: 170000},
        {Brand: "Suzuki", Model: "Satria", Price: 80000},
    }

    customers := []Customer{
        {Name: "Diko Cesartista", Address: "Jalan Sorosutan", Rentals: []Rental{}},
        {Name: "Vitri Septia", Address: "Jalan Imogiri", Rentals: []Rental{}},
    }

    // Contoh
    customers[0].rentVehicle(vehicles[0], 5)
    customers[1].rentVehicle(vehicles[1], 3)

    customers[0].returnVehicle(vehicles[0])

    for _, customer := range customers {
        fmt.Printf("Nama Customer: %s\n", customer.Name)
        fmt.Printf("Alamat: %s\n", customer.Address)
        fmt.Println("Kendaraan yang disewa:")
        for _, rental := range customer.Rentals {
            fmt.Printf("- %s %s (Rp %d per hari) selama %d hari\n", rental.Brand, rental.Model, rental.Price, rental.Duration)
        }
        fmt.Println()
    }
}

func (c *Customer) rentVehicle(vehicle Vehicle, duration int) {
    rental := Rental{Vehicle: vehicle, Duration: duration}
    c.Rentals = append(c.Rentals, rental)
}

func (c *Customer) returnVehicle(vehicle Vehicle) {
    var updatedRentals []Rental
    for _, rental := range c.Rentals {
        if rental.Brand != vehicle.Brand || rental.Model != vehicle.Model {
            updatedRentals = append(updatedRentals, rental)
        }
    }
    c.Rentals = updatedRentals
}
