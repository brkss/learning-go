package elon

import "strconv"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
    if car.battery - car.batteryDrain >= 0 {
        car.distance = car.distance + car.speed;
    	car.battery = car.battery - car.batteryDrain;
    }
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
    var out string;

    out = "Driven " + strconv.Itoa(car.distance) + " meters";
    return (out);
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
    var out string;

    out = "Battery at " + strconv.Itoa(car.battery) + "%" 
	return (out);
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
    var nDrive int;

    nDrive = trackDistance / car.speed;
    return (nDrive * car.batteryDrain) <= car.battery;
}
