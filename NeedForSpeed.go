package speed

// TODO: define the 'Car' type struct
type Car struct  {
    battery int;
    batteryDrain int;
    speed int;
    distance int;
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    var car Car;

    car.battery = 100;
    car.batteryDrain = batteryDrain;
    car.speed = speed;
    car.distance = 0;
    return (car);
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int;
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	var track Track;

    track.distance = distance;
    return (track);
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery - car.batteryDrain >= 0 {
        car.distance = car.distance + car.speed;
    	car.battery = car.battery - car.batteryDrain;
    }
    return (car);
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    var distanceRatio int;

    distanceRatio = track.distance / car.speed;
    if distanceRatio * car.batteryDrain <= car.battery {
        return (true)
    }else {
    	return (false);
    }
}

