#include <Arduino.h>

#include "ACS712.h"
#include <ArduinoJson.h>



// init pin acs
ACS712 sensor1(ACS712_05B, A1);
ACS712 sensor2(ACS712_05B, A2);
ACS712 sensor3(ACS712_05B, A3);
/****** var for reading current sensor*****/
double sensorValue=0;
double sensorValue1=0;
int crosscount=0;
int climbhill=0;
double VmaxD=0;
double VeffD;
double Veff;
// variabel untuk cek jika tegangan tidak terbaca, kirim data 0
// constants won't change. Used here to set a pin number:
const int ledPin =  LED_BUILTIN;// the number of the LED pin

// Variables will change:
int ledState = LOW;             // ledState used to set the LED

// Generally, you should use "unsigned long" for variables that hold time
// The value will quickly become too large for an int to store
unsigned long previousMillis = 0;        // will store last time LED was updated

// constants won't change:
const long interval = 3000;   
float bacaSensorTegangan(){
   //http://sentroino.blogspot.co.id/2015/12/pembacaan-tegangan-ac-menggunakan.html
sensorValue1=sensorValue;
delay(100);
sensorValue = analogRead(A0);
if (sensorValue>sensorValue1 && sensorValue>511){
  climbhill=1;
  VmaxD=sensorValue;
  }
if (sensorValue<sensorValue1 && climbhill==1){
  climbhill=0;
  VmaxD=sensorValue1;
  VeffD=VmaxD/sqrt(2);
  Veff=(((VeffD-420.76)/-90.24)*-210.2)+210.2;
  // Serial.println(Veff);
  VmaxD=0;
  return Veff;
}
return -1;
}
// struct data JSON
struct dataJSONs {
float tegangan;
float arus1;
float arus2;
float arus3;
  
} dataJSON;  
 StaticJsonBuffer<200> jsonBuffer;
   JsonObject& root =jsonBuffer.createObject();
void sendJSON(){
root["tegangan"] = String(dataJSON.tegangan);
root["arus1"] = String(dataJSON.arus1);
root["arus2"] = String(dataJSON.arus2);
root["arus3"] = String(dataJSON.arus3);
  root.printTo(Serial);
  Serial.print("#");
}
void setup() {
  Serial.begin(9600);
  delay(1000);
  // calibrate() method calibrates zero point of sensor
  sensor1.calibrate();
  sensor2.calibrate();
  sensor3.calibrate();
  delay(500);
//   Serial.println("Done!");
}

void loop() {
// baca tegangan sensor , jika sukses nilai != -1
    unsigned long currentMillis = millis();
  float U =bacaSensorTegangan();
  if (U!=-1){
dataJSON.tegangan=U;
// dapatkan 3 nilai sensor
dataJSON.arus1=sensor1.getCurrentAC();
dataJSON.arus2=sensor2.getCurrentAC();
dataJSON.arus3=sensor3.getCurrentAC();
  sendJSON();
  delay(100);
    previousMillis = currentMillis;
  }

  if (currentMillis - previousMillis >= interval) {
    dataJSON.tegangan=0;
// dapatkan 3 nilai sensor
dataJSON.arus1=sensor1.getCurrentAC();
dataJSON.arus2=sensor2.getCurrentAC();
dataJSON.arus3=sensor3.getCurrentAC();
  sendJSON();

    // if the LED is off turn it on and vice-versa:
    if (ledState == LOW) {
      ledState = HIGH;
    } else {
      ledState = LOW;
    }
    // set the LED with the ledState of the variable:
    digitalWrite(ledPin, ledState);
    // save the last time you blinked the LED
    previousMillis = currentMillis;

  }
}

