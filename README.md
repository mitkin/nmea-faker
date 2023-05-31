# NMEA fake position over TCP

Tiny worker sending NMEA 0183 messages over TCP.
This was implemented to assist a demo session with ship positioninig information.

Program will start a coroutine and will broadcast position with true heading changing in real time.

## Build
```
make build

```

## Run
```
./nmea-faker -port 5002
```

