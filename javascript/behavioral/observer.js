// Observer interface
class Observer {
    onUpdate(subject) {
        throw new Error('Update method not implemented');
    }
}

// Subject abstract class
class Subject {
    constructor() {
        this.observers = [];
    }
    attatch(observer) {
        this.observers.push(observer);
    }

    // this might not work
    detatch(observer) {
        this.observers.filter( obs => obs !== observer);
    }

    notify() {
        this.observers.forEach( (observer) => {
            observer.onUpdate(this);
        });
    }
}

class ClockTimer extends Subject {
    constructor() {
        super();
        this.seconds = 0;
        setInterval(() => {this.tick()}, 1000)
    }

    tick() {
        this.seconds += 1000;
        this.notify();
    }
}

class DigitalClock extends Observer {
    constructor(clockTimer) {
        super();
        this.clockTimer = clockTimer;
        this.clockTimer.attatch(this);
    }

    onUpdate(subject) {
        console.log('digital clock updated', subject.seconds);
    }
}

class AnalogClock extends Observer {
    constructor(clockTimer) {
        super();
        this.clockTimer = clockTimer;
        this.clockTimer.attatch(this);
    }

    onUpdate(clock) {
        console.log('analog clock updated', clock.seconds / 1000);
    }
}

const clockTimer = new ClockTimer();
const digitalClock = new DigitalClock(clockTimer);
const analogClock = new AnalogClock(clockTimer);