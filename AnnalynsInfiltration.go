package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
    return (!knightIsAwake);
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var canSpy bool;

    canSpy = knightIsAwake || archerIsAwake || prisonerIsAwake;
    return (canSpy);
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	var canSignal bool;

    canSignal = !archerIsAwake && prisonerIsAwake;
    return (canSignal);
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var canFreePrisoner bool;
    var strategyA bool;
    var strategyB bool;
    
	strategyA = petDogIsPresent && !archerIsAwake;
    strategyB = !archerIsAwake && !knightIsAwake && prisonerIsAwake;
    canFreePrisoner = strategyA || strategyB;
    return (canFreePrisoner);
}

