# NMEA fake position over TCP

Tiny worker sending NMEA 0183 messages over TCP.
Adjust port and host to your liking and start:

```
python3 main.py
```
This was implemented to assist a demo session with ship positioninig information.

The go version will start a coroutine and will broadcast position with true heading changing in real time.
To run it execute:

```
go run main.go

```
