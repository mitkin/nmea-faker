import socket
import time

def broadcast_nmea_messages(nmea_messages, tcp_ip, tcp_port):
    # Create a TCP socket
    tcp_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp_socket.bind((tcp_ip, tcp_port))
    tcp_socket.listen()
    
    while True:
        # Accept incoming connections
        connection, address = tcp_socket.accept()
        
        # Send NMEA messages continuously
        for message in nmea_messages:
            # Send the message over TCP
            connection.send((message + '\n').encode())
            
            # Wait for a short interval before sending the next message
            time.sleep(0.01)
        
        # Close the connection
        connection.close()

# Example usage
if __name__ == "__main__":
    # Define the TCP IP and port
    tcp_ip = "127.0.0.1"  # Replace with the desired IP address
    tcp_port = 5000  # Replace with the desired port number
    
    # Example NMEA messages
        # Example NMEA messages
    nmea_messages = [
            """
$INZDA,163611.11,10,09,2019,,*76
$INGGA,163611.11,7849.766185,N,00543.603403,W,2,07,1.2,2.95,M,36.08,M,2.0,0123*4D
$INGLL,7849.766185,N,00543.603403,W,163611.11,A,D*65
$INVTG,61.99,T,68.23,M,5.0,N,9.2,K,D*30
$INHDT,66.99,T*25
$PSXN,23,0.01,0.12,66.99,0.07*0D

$INZDA,163612.11,10,09,2019,,*75
$INGGA,163612.11,7849.766811,N,00543.597037,W,2,07,1.2,2.75,M,36.08,M,3.0,0123*48
$INGLL,7849.766811,N,00543.597037,W,163612.11,A,D*6F
$INVTG,64.06,T,70.31,M,5.0,N,9.3,K,D*38
$INHDT,67.83,T*2F
$PSXN,23,-0.10,-0.10,67.83,0.27*07

$INZDA,163613.11,10,09,2019,,*74
$INGGA,163613.11,7849.767422,N,00543.590478,W,2,07,1.2,2.69,M,36.08,M,2.0,0123*40
$INGLL,7849.767422,N,00543.590478,W,163613.11,A,D*6B
$INVTG,64.69,T,70.94,M,5.1,N,9.5,K,D*39
$INHDT,68.64,T*29
$PSXN,23,-0.06,0.01,68.64,0.32*2F
"""
]

    # Broadcast the NMEA messages over TCP
    broadcast_nmea_messages(nmea_messages, tcp_ip, tcp_port)
